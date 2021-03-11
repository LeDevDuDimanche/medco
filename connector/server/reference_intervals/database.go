package referenceintervalserver

import (
	"fmt"
	"strconv"

	utilserver "github.com/ldsec/medco/connector/util/server"
	"github.com/sirupsen/logrus"
)

//QueryResult contains the information about a row that comes out from the query executed in RetrievePatients.
type QueryResult struct {
	NumericValue  float64
	Units         string
	PatientNumber int64
}

//RetrieveObservations returns the numerical values that correspond to the concept or modifier passed as argument for the specified cohort.
func RetrieveObservations(conceptCode string, modifierCode string, patients []int64) (queryResults []QueryResult, err error) {
	strPatientList := utilserver.ConvertIntListToString(patients)

	logrus.Info("About to execute explore stats SQL query ", sql, " ", conceptCode, " ", modifierCode, " ", strPatientList)

	rows, err := utilserver.I2B2DBConnection.Query(sql, conceptCode, modifierCode, strPatientList)
	if err != nil {
		err = fmt.Errorf("while execution SQL query: %s", err.Error())
		return
	}

	queryResults = make([]QueryResult, 0)

	for rows.Next() {
		numericValue := new(string)
		patientNb := new(string)
		units := new(string)
		scanErr := rows.Scan(numericValue, patientNb, units)
		if scanErr != nil {
			err = scanErr
			err = fmt.Errorf("while scanning SQL record: %s", err.Error())
			return
		}

		//TODO remove this
		logrus.Info("Scanned numerical values explore stats ", numericValue, " ", patientNb, " ", units)

		var queryResult QueryResult

		queryResult.Units = *units

		queryResult.NumericValue, err = strconv.ParseFloat(*numericValue, 64)
		if err != nil {
			err = fmt.Errorf("Error while converting numerical value %s for the (concept, modifier) with code (%s, %s)", *numericValue, conceptCode, modifierCode)
			return
		}

		queryResult.PatientNumber, err = strconv.ParseInt(*patientNb, 10, 64)
		if err != nil {
			err = fmt.Errorf("Error while parsing the patient identifier %s for the (concept, modifier) with code (%s, %s)", *patientNb, conceptCode, modifierCode)
			return
		}

		queryResults = append(queryResults, queryResult)
	}

	return

}

/*
	* This query will return the numerical values from all observations where
	* the patient_num is contained within the list passed as argument (the list is in principle a list of patient from a specific cohort).

	TODO In the same way I gathered the schema and table in which the ontology is contained, gather the schema in which observations are contained.
	For the moment I hardcode the table and schema.

	We only keep rows where nval_num is exactly equal to a specific values hence the required value of TVAL_CHAR.
	We could keep values which are GE or LE or L or G the problem is that we would need open brackets for intervals.
	VALTYPE_CD = 'N' because we only care about numerical values.
*/

const sql string = `
SELECT nval_num, patient_num, units_cd FROM i2b2demodata_i2b2.observation_fact
	WHERE
		concept_cd = $1 AND modifier_cd = $2 AND
		valtype_cd = 'N' AND tval_char = 'E' AND nval_num is not null AND units_cd is not null AND units_cd != '@'
		AND patient_num = ANY($3::integer[]) 
`
