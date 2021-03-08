package explorestatisticsclient

import (
	"time"

	"github.com/sirupsen/logrus"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/ldsec/medco/connector/restapi/client/explore_statistics"
	utilclient "github.com/ldsec/medco/connector/util/client"
)

func (exploreStatistics *ExploreStatistics) submitToNode(nodeIdx int) (results *explore_statistics.ExploreStatisticsOKBody, err error) {

	params := explore_statistics.NewExploreStatisticsParamsWithTimeout(time.Duration(utilclient.ExploreStatisticsTimeoutSeconds) * time.Second)

	body := explore_statistics.ExploreStatisticsBody{
		ID:              exploreStatistics.id,
		UserPublicKey:   exploreStatistics.userPublicKey,
		CohortName:      exploreStatistics.cohortName,
		Concept:         exploreStatistics.conceptPath,
		Modifier:        exploreStatistics.modifier,
		NumberOfBuckets: exploreStatistics.nbBuckets,
	}

	params.SetBody(body)
	response, err := exploreStatistics.httpMedCoClients[nodeIdx].ExploreStatistics.ExploreStatistics(params, httptransport.BearerToken(exploreStatistics.authToken))

	if err != nil {
		logrus.Error("Explore statistics error: ", err)
		return
	}
	results = response.GetPayload()

	return
}
