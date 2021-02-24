package handlers

import (
	"fmt"

	explore_statistics "github.com/ldsec/medco/connector/restapi/server/operations/explore_statistics"

	querytoolsserver "github.com/ldsec/medco/connector/server/querytools"

	"github.com/go-openapi/runtime/middleware"
	"github.com/ldsec/medco/connector/restapi/models"
	"github.com/ldsec/medco/connector/restapi/server/operations/survival_analysis"
	referenceintervalserver "github.com/ldsec/medco/connector/server/reference_intervals"
	"github.com/sirupsen/logrus"
)

// ExploreStatisticsHandler handles /survival-analysis API endpoint
func ExploreStatisticsHandler(param explore_statistics.ExploreStatisticsParams, principal *models.User) middleware.Responder {

	survivalAnalysisQuery := referenceintervalserver.NewQuery(
		principal.ID,
		param.Body.ID,
		param.Body.UserPublicKey,
		param.Body.CohortName,
		param.Body.Concept,
		param.Body.Modifier,
		param.Body.NumberOfBuckets)

	logrus.Debug("survivalAnalysis: ", survivalAnalysisQuery)
	err := survivalAnalysisQuery.Validate()
	if err != nil {
		logrus.Error(err)
		return survival_analysis.NewSurvivalAnalysisBadRequest().WithPayload(
			&survival_analysis.SurvivalAnalysisBadRequestBody{
				Message: "Survival query validation error:" + err.Error()})
	}
	found := false
	logrus.Info("checking cohort's existence")
	found, err = querytoolsserver.DoesCohortExist(principal.ID, *param.Body.CohortName)
	if err != nil {
		logrus.Error(err)
		return survival_analysis.NewSurvivalAnalysisDefault(500).WithPayload(
			&survival_analysis.SurvivalAnalysisDefaultBody{
				Message: "Survival query execution error:" + err.Error()})
	}

	if !found {
		logrus.Info("cohort not found")
		return survival_analysis.NewSurvivalAnalysisNotFound().WithPayload(
			&survival_analysis.SurvivalAnalysisNotFoundBody{
				Message: fmt.Sprintf("Cohort %s not found", *param.Body.CohortName),
			},
		)
	}
	logrus.Info("cohort found")

	err = survivalAnalysisQuery.Execute()

	if err != nil {
		logrus.Error(err)
		return survival_analysis.NewSurvivalAnalysisDefault(500).WithPayload(
			&survival_analysis.SurvivalAnalysisDefaultBody{
				Message: "Survival query execution error:" + err.Error()})

	}
	results := survivalAnalysisQuery.Result

	resultList := make([]*survival_analysis.SurvivalAnalysisOKBodyResultsItems0, 0)
	for _, group := range survivalAnalysisQuery.Result.EncEvents {

		timePoints := make([]*survival_analysis.SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0, 0)
		for _, timePoint := range group.TimePointResults {
			timePoints = append(timePoints, &survival_analysis.SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0{Timepoint: int64(timePoint.TimePoint),
				Events: &survival_analysis.SurvivalAnalysisOKBodyResultsItems0GroupResultsItems0Events{
					Eventofinterest: timePoint.Result.EventValueAgg,
					Censoringevent:  timePoint.Result.CensoringValueAgg,
				}})
		}
		resultList = append(resultList, &survival_analysis.SurvivalAnalysisOKBodyResultsItems0{
			GroupID:      group.GroupID,
			InitialCount: group.EncInitialCount,
			GroupResults: timePoints,
		})
	}

	//parse timers
	modelsTimers := results.Timers.TimersToAPIModel()

	requestResult := &survival_analysis.SurvivalAnalysisOKBody{Results: resultList, Timers: modelsTimers}

	return survival_analysis.NewSurvivalAnalysisOK().WithPayload(requestResult)

}
