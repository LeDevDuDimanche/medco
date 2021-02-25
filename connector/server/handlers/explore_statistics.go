package handlers

import (
	"fmt"
	"strconv"

	explore_statistics "github.com/ldsec/medco/connector/restapi/server/operations/explore_statistics"

	querytoolsserver "github.com/ldsec/medco/connector/server/querytools"

	"github.com/go-openapi/runtime/middleware"
	"github.com/ldsec/medco/connector/restapi/models"
	referenceintervalserver "github.com/ldsec/medco/connector/server/reference_intervals"
	"github.com/sirupsen/logrus"
)

// ExploreStatisticsHandler handles /survival-analysis API endpoint
func ExploreStatisticsHandler(param explore_statistics.ExploreStatisticsParams, principal *models.User) middleware.Responder {

	query, err := referenceintervalserver.NewQuery(
		principal.ID,
		param.Body.ID,
		param.Body.UserPublicKey,
		param.Body.CohortName,
		param.Body.Concept,
		param.Body.Modifier,
		param.Body.NumberOfBuckets)

	if err != nil {
		logrus.Error(err)
		explore_statistics.NewExploreStatisticsBadRequest().WithPayload(
			&explore_statistics.ExploreStatisticsBadRequestBody{
				Message: "Explore statistics query creation error:" + err.Error()})

	}

	logrus.Debug("survivalAnalysis: ", query)
	err = query.Validate()
	if err != nil {
		logrus.Error(err)
		explore_statistics.NewExploreStatisticsBadRequest().WithPayload(
			&explore_statistics.ExploreStatisticsBadRequestBody{
				Message: "Explore statistics query validation error:" + err.Error()})

	}

	found := false
	logrus.Info("checking cohort's existence")
	found, err = querytoolsserver.DoesCohortExist(principal.ID, param.Body.CohortName)
	if err != nil {
		logrus.Error(err)
		explore_statistics.NewExploreStatisticsDefault(500).WithPayload(
			&explore_statistics.ExploreStatisticsDefaultBody{
				Message: "Explore statistics query execution error:" + err.Error()})
	}

	if !found {
		logrus.Info("cohort not found")
		return explore_statistics.NewExploreStatisticsNotFound().WithPayload(
			&explore_statistics.ExploreStatisticsNotFoundBody{
				Message: fmt.Sprintf("Cohort %s not found", param.Body.CohortName)})
	}
	logrus.Info("cohort found")

	err = query.Execute()

	if err != nil {
		logrus.Error(err)
		explore_statistics.NewExploreStatisticsDefault(500).WithPayload(
			&explore_statistics.ExploreStatisticsDefaultBody{
				Message: "Explore statistics query execution error:" + err.Error()})
	}

	results := query.Result

	//parse timers
	modelsTimers := results.Timers.TimersToAPIModel()

	apiIntervals := make([]*models.IntervalBucket, len(results.Intervals))

	for _, bucket := range results.Intervals {
		higherBound := strconv.FormatFloat(bucket.HigherBound, 'f', 5, 64)
		lowerBound := strconv.FormatFloat(bucket.LowerBound, 'f', 5, 64)

		apiBucket := models.IntervalBucket{
			EncCount:    &bucket.EncCount,
			HigherBound: &higherBound,
			LowerBound:  &lowerBound,
		}

		apiIntervals = append(apiIntervals, &apiBucket)
	}

	requestResult := explore_statistics.ExploreStatisticsOKBody{
		Intervals: apiIntervals,
		Unit:      results.Unit,
		Timers:    modelsTimers,
	}

	return explore_statistics.NewExploreStatisticsOK().WithPayload(&requestResult)

}
