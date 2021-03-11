package explorestatisticsclient

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	medcomodels "github.com/ldsec/medco/connector/models"

	medcoclient "github.com/ldsec/medco/connector/client"
	utilclient "github.com/ldsec/medco/connector/util/client"

	"github.com/ldsec/medco/connector/restapi/client/explore_statistics"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/ldsec/medco/connector/restapi/client"
	"github.com/ldsec/medco/connector/wrappers/unlynx"
	"github.com/sirupsen/logrus"
)

//ExploreStatistics represents an explore query request
type ExploreStatistics struct {
	// httpMedCoClient is the HTTP client for the MedCo connector
	httpMedCoClients []*client.MedcoCli
	// authToken is the OIDC authentication JWT
	authToken string

	id string

	cohortName string

	nbBuckets int64

	conceptPath string
	modifier    *explore_statistics.ExploreStatisticsParamsBodyModifier

	userPublicKey  string
	userPrivateKey string

	formats strfmt.Registry

	timers medcomodels.Timers
}

// NewExploreStatisticsQuery constructor for explore statistics request
func NewExploreStatisticsQuery(
	token string,
	cohortName string,
	nbBuckets int64,
	conceptPath string,
	modifier *explore_statistics.ExploreStatisticsParamsBodyModifier,
	disableTLSCheck bool,
) (q *ExploreStatistics, err error) {
	q = &ExploreStatistics{
		authToken:   token,
		id:          "MedCo_Explore_Statistics" + time.Now().Format(time.RFC3339),
		cohortName:  cohortName,
		conceptPath: conceptPath,
		nbBuckets:   nbBuckets,
		modifier:    modifier,
		formats:     strfmt.Default,
		timers:      make(map[string]time.Duration),
	}

	getMetadataResp, err := medcoclient.MetaData(token, disableTLSCheck)
	if err != nil {
		logrus.Error(err)
		return
	}

	q.httpMedCoClients = make([]*client.MedcoCli, len(getMetadataResp.Payload.Nodes))
	for _, node := range getMetadataResp.Payload.Nodes {
		if q.httpMedCoClients[*node.Index] != nil {
			err = errors.New("duplicated node index in network metadata")
			logrus.Error(err)
			return
		}

		nodeURL, err := url.Parse(node.URL)
		if err != nil {
			logrus.Error("cannot parse MedCo connector URL: ", err)
			return nil, err
		}

		nodeTransport := httptransport.New(nodeURL.Host, nodeURL.Path, []string{nodeURL.Scheme})
		nodeTransport.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: disableTLSCheck}
		q.httpMedCoClients[*node.Index] = client.New(nodeTransport, nil)
	}

	q.userPublicKey, q.userPrivateKey, err = unlynx.GenerateKeyPair()
	if err != nil {
		return
	}

	return

}

//Decrypt deciphers a value that is expected to be encrypted under user public key
func (exploreStatistics *ExploreStatistics) Decrypt(value string) (int64, error) {
	return unlynx.Decrypt(value, exploreStatistics.userPrivateKey)
}

//NodeResult contains the raw cyphered answer from the server with index NodeIndex.
type NodeResult struct {
	Body      *explore_statistics.ExploreStatisticsOKBody
	NodeIndex int
}

//Execute makes a call to the API that handle explore statistics requests,
func (exploreStatistics *ExploreStatistics) Execute() (results []*EncryptedResults, err error) {

	var nbOfNodes = len(exploreStatistics.httpMedCoClients)
	errChan := make(chan error)
	resultChan := make(chan NodeResult, nbOfNodes)
	results = make([]*EncryptedResults, nbOfNodes)
	logrus.Infof("There are %d nodes", nbOfNodes)

	for idx := 0; idx < nbOfNodes; idx++ {

		go func(idx int) {
			logrus.Infof("Submitting to node %d", idx)
			res, Error := exploreStatistics.submitToNode(idx)
			if Error != nil {
				logrus.Errorf("Explore statistics execution error : %s", Error)
				errChan <- Error
				return
			}

			resultChan <- NodeResult{Body: res, NodeIndex: idx}
		}(idx)
	}
	timeout := time.After(time.Duration(utilclient.ExploreStatisticsTimeoutSeconds) * time.Second)
	for idx := 0; idx < nbOfNodes; idx++ {
		select {
		case err = <-errChan:
			return
		case <-timeout:
			err = fmt.Errorf("Timeout %d seconds elapsed", utilclient.ExploreStatisticsTimeoutSeconds)
			logrus.Error(err)
			return
		case nodeRes := <-resultChan:
			encryptedResult := new(EncryptedResults)
			encryptedResult.Intervals = nodeRes.Body.Intervals
			encryptedResult.Unit = nodeRes.Body.Unit
			encryptedResult.Timers = medcomodels.NewTimersFromAPIModel(nodeRes.Body.Timers)
			results[idx] = encryptedResult
			logrus.Infof("Node %d successfully fetched explore statistics data", idx)
		}
	}

	return
}

func (exploreStatistics *ExploreStatistics) addTimer(label string, since time.Time) (err error) {
	if _, exists := exploreStatistics.timers[label]; exists {
		err = errors.New("Timer label " + label + " already exists")
		return
	}
	exploreStatistics.timers[label] = time.Since(since)
	return

}

func (exploreStatistics *ExploreStatistics) addTimers(timers map[string]time.Duration) (err error) {
	for label, duration := range timers {
		if _, exists := exploreStatistics.timers[label]; exists {
			err = errors.New("Timer label " + label + " already exists")
			return
		}
		exploreStatistics.timers[label] = duration
	}
	return

}
