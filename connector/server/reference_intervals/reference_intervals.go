package referenceintervalserver

import (
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
	"time"

	medcomodels "github.com/ldsec/medco/connector/models"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	querytoolsserver "github.com/ldsec/medco/connector/server/querytools"
	utilserver "github.com/ldsec/medco/connector/util/server"
	"github.com/ldsec/medco/connector/wrappers/i2b2"
	"github.com/ldsec/medco/connector/wrappers/unlynx"

	"github.com/ldsec/medco/connector/restapi/server/operations/explore_statistics"
)

// Interval is a structure containing the lower bound and higher bound of an interval.
// The lower bound is inclusive and the higher bound is exclusive: [lower bound, higher bound[
type Interval struct {
	LowerBound  float64
	HigherBound float64
	EncCount    string // contains the count of subject in this interval
}

// Query holds the ID of the survival analysis, its parameters and a pointer to its results
type Query struct {
	UserID        string
	UserPublicKey string
	QueryName     string
	CohortName    string
	Concept       string
	Modifier      *explore_statistics.ExploreStatisticsParamsBodyModifier //TODO export this class out of the survival package make it a common thing
	Result        *struct {
		Timers    medcomodels.Timers
		Intervals []Interval
		Unit      string
	}
}

// NewQuery query constructor
func NewQuery(
	UserID string,
	QueryName string,
	UserPublicKey string,
	CohortName string,
	Concept string,
	modifier *explore_statistics.ExploreStatisticsParamsBodyModifier,
	nbBounds int64,
) (q *Query, err error) {

	if nbBounds < 0 {
		err := fmt.Errorf("no intervals specified in the parameters of the query")
		return nil, err
	}

	res := &Query{
		UserID:        UserID,
		UserPublicKey: UserPublicKey,
		CohortName:    CohortName,
		QueryName:     QueryName,
		Concept:       Concept,
		Modifier:      modifier,
		Result: &struct {
			Timers    medcomodels.Timers
			Intervals []Interval
			Unit      string
		}{}}

	res.Result.Intervals = make([]Interval, nbBounds)
	res.Result.Timers = make(map[string]time.Duration)

	return res, nil
}

// Execute runs the survival analysis query
func (q *Query) Execute() error {
	//TODO verify the user has the right to execute such a query.

	encCounts := make([]string, 0)
	timer := time.Now()

	conceptCode, modifierCode, cohort, timers, err := prepareArguments(q.UserID, q.CohortName, q.Concept, q.Modifier)
	if err != nil {
		err = fmt.Errorf("while retrieving concept codes and patient indices: %s", err.Error())
		return err
	}
	q.Result.Timers.AddTimers("", timer, timers)

	queryResults, err := RetrieveObservations(conceptCode, modifierCode, cohort)

	//TODO build intervals from modifier or concept

	// TODO il faut que tu t'assures d'utiliser tout le temps la même unité
	// pour cela il faudrait que l'utilisateur ait le choix de l'unité qui sera utilisée dans l'affichage distribution.
	// voir https://community.i2b2.org/wiki/display/DevForum/Metadata+XML+for+Medication+Modifiers
	//TODO Une fois que tu connais l'unité il faut remplir le champs Unit de la requête de réponse.

	if err != nil {
		return err
	}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(q.Result.Intervals))
	channels := make([]chan struct {
		encCount *string
		medcomodels.Timers
	}, len(q.Result.Intervals))
	errChan := make(chan error, len(q.Result.Intervals))
	signal := make(chan struct{})

	//TODO dans le futur il faudra que tu puisses reproduire les intervalles de références. Il faudra que tu sauvegardes des informations dans la db medco pour ça.

	for i, interval := range q.Result.Intervals {
		if interval.LowerBound >= interval.HigherBound {
			err := fmt.Errorf("the lower bound of the interval is greater than the higher bound: %f >= %f", interval.LowerBound, interval.HigherBound)
			errChan <- err
			break
		}

		channels[i] = make(chan struct {
			encCount *string
			medcomodels.Timers
		}, 1)

		go func(i int, interval *Interval) {
			defer waitGroup.Done()
			timers := medcomodels.NewTimers()

			keptResults := make([]QueryResult, 0)

			for _, queryResult := range queryResults {
				if queryResult.NumericValue >= interval.LowerBound && queryResult.NumericValue < interval.HigherBound {
					keptResults = append(keptResults, queryResult)
				}
			}

			timer = time.Now()

			encCount, err := unlynx.EncryptWithCothorityKey(int64(len(keptResults)))
			timers.AddTimers(fmt.Sprintf("medco-connector-encrypt-interval-count-group%d", i), timer, nil)
			if err != nil {
				err = fmt.Errorf("while encrypting the count of an interval of the future reference interval: %s", err.Error())
				errChan <- err
				return
			}

			channels[i] <- struct {
				encCount *string
				medcomodels.Timers
			}{&encCount, timers}
		}(i, &interval)

	}
	go func() {
		waitGroup.Wait()
		signal <- struct{}{}
	}()

	select {
	case err := <-errChan:
		return err
	case <-signal:
		break
	}

	for _, channel := range channels {
		chanResult := <-channel

		encCounts = append(encCounts, *chanResult.encCount)
		q.Result.Timers.AddTimers("", timer, chanResult.Timers)
	}

	// aggregate and key switch locally encrypted results
	timer = time.Now()
	var aggregationTimers medcomodels.Timers
	var aggValues []string
	aggValues, aggregationTimers, err = unlynx.AggregateAndKeySwitchValues(q.QueryName+"_AGG_AND_KEYSWITCH", encCounts, q.UserPublicKey)

	//assign the encrypted count to the matching interval
	for i, interval := range q.Result.Intervals {
		interval.EncCount = aggValues[i]
	}

	q.Result.Timers.AddTimers("medco-connector-aggregate-and-key-switch", timer, aggregationTimers)
	if err != nil {
		err = fmt.Errorf("during aggregation and keyswitch: %s", err.Error())
	}
	return err
}

// Validate checks members of a Query instance for early error detection.
// Heading and trailing spaces are silently trimmed. Granularity string is silently written in lower case.
// If any other wrong member can be defaulted, a warning message is printed, otherwise an error is returned.
func (q *Query) Validate() error {

	q.QueryName = strings.TrimSpace(q.QueryName)
	if q.QueryName == "" {
		return fmt.Errorf("empty query name")
	}

	q.Concept = strings.TrimSpace(q.Concept)
	if q.Concept == "" {
		return fmt.Errorf("emtpy start concept path, queryID: %s", q.QueryName)
	}
	if q.Modifier != nil {
		*q.Modifier.ModifierKey = strings.TrimSpace(*q.Modifier.ModifierKey)
		if *q.Modifier.ModifierKey == "" {
			return fmt.Errorf("empty start modifier key, queryID: %s, start concept: %s", q.QueryName, q.Concept)
		}
		*q.Modifier.AppliedPath = strings.TrimSpace(*q.Modifier.AppliedPath)
		if *q.Modifier.AppliedPath == "" {
			return fmt.Errorf(
				"empty start modifier applied path, queryID: %s, start concept: %s, start modifier key: %s",
				q.QueryName, q.Concept,
				*q.Modifier.ModifierKey,
			)
		}
	}

	q.UserID = strings.TrimSpace(q.UserID)
	if q.UserID == "" {
		return fmt.Errorf("empty user name, queryID: %s", q.QueryName)
	}

	q.UserPublicKey = strings.TrimSpace(q.UserPublicKey)
	if q.UserPublicKey == "" {
		return fmt.Errorf("empty user public keyqueryID: %s", q.QueryName)
	}
	_, err := base64.URLEncoding.DecodeString(q.UserPublicKey)
	if err != nil {
		return fmt.Errorf("user public key is not valid against the alternate RFC4648 base64 for URL: %s; queryID: %s", err.Error(), q.QueryName)
	}
	return nil

}

// prepareArguments retrieves concept codes and patients that will be used as the arguments of direct SQL call
func prepareArguments(
	userID,
	cohortName,
	concept string,
	modifier *explore_statistics.ExploreStatisticsParamsBodyModifier,
) (
	conceptCode,
	modifierCode string,
	cohort []int64, timers medcomodels.Timers,
	err error,
) {
	timers = make(map[string]time.Duration)
	// --- cohort patient list
	timer := time.Now()
	logrus.Info("get patients")
	cohort, err = querytoolsserver.GetPatientList(userID, cohortName)

	if err != nil {
		logrus.Error("error while getting patient list")
		return
	}

	timers.AddTimers("medco-connector-get-patient-list", timer, nil)
	logrus.Info("got patients")

	// --- get concept and modifier codes from the ontology
	logrus.Info("get concept and modifier codes")
	err = utilserver.I2B2DBConnection.Ping()
	if err != nil {
		err = fmt.Errorf("while connecting to clear project database: %s", err.Error())
		return
	}
	conceptCode, err = getCode(concept)
	if err != nil {
		err = fmt.Errorf("while retrieving start concept code: %s", err.Error())
		return
	}
	if modifier == nil {
		modifierCode = "@"
	} else {
		modifierCode, err = getModifierCode(*modifier.ModifierKey, *modifier.AppliedPath)
	}
	if err != nil {
		err = fmt.Errorf("while retrieving start modifier code: %s", err.Error())
		return
	}

	if err != nil {
		err = fmt.Errorf("while retrieving end modifier code: %s", err.Error())
		return
	}
	logrus.Info("got concept and modifier codes")
	return
}

// getCode takes the full path of a I2B2 concept and returns its code
func getCode(path string) (string, error) {
	logrus.Debugf("get code concept path %s", path)
	res, err := i2b2.GetOntologyConceptInfo(path)
	if err != nil {
		return "", err
	}
	if len(res) != 1 {
		return "", errors.Errorf("Result length of GetOntologyConceptInfo is expected to be 1. Got: %d", len(res))
	}

	if res[0].Code == "" {
		return "", errors.New("Code is empty")
	}
	logrus.Debugf("got concept code %s", res[0].Code)

	return res[0].Code, nil

}

// getModifierPath takes the full path of a I2B2 modifier and its applied paht and returns its code
func getModifierCode(path string, appliedPath string) (string, error) {
	logrus.Debugf("get modifier code modifier path %s applied path %s", path, appliedPath)
	res, err := i2b2.GetOntologyModifierInfo(path, appliedPath)
	if err != nil {
		return "", err
	}

	if len(res) != 1 {
		return "", errors.Errorf("Result length of GetOntologyTermInfo is expected to be 1. Got: %d. "+
			"Is applied path %s available for modifier key %s ?", len(res), appliedPath, path)
	}
	if res[0].Code == "" {
		return "", errors.New("Code is empty")
	}
	if res[0].AppliedPath != appliedPath {
		return "", fmt.Errorf("applied paths don't match. Is applied path %s available for modifier key %s ?", appliedPath, path)
	}
	logrus.Debugf("got modifier code %s", res[0].Code)

	return res[0].Code, nil
}
