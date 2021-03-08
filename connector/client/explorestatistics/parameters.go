package explorestatisticsclient

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Parameters holds parameters to build an explore statistics query
type Parameters struct {
	CohortName  string    `yaml:"cohort_name"`
	ConceptPath string    `yaml:"start_concept_path"`
	Modifier    *modifier `yaml:"start_modifier,omitempty"`
	nbBuckets   int64     `yaml:"number_buckets"`
}

type modifier struct {
	ModifierKey string `yaml:"modifier_key"`
	AppliedPath string `yaml:"applied_path"`
}

// NewParametersFromFile builds a Parameters instance from YAML file
func NewParametersFromFile(fileURL string) (*Parameters, error) {
	logrus.Trace(fileURL)
	logrus.Trace(os.Getenv("PWD"))
	file, err := os.Open(fileURL)
	if err != nil {
		return nil, fmt.Errorf("while opening parameter file %s: %s", fileURL, err.Error())
	}
	decoder := yaml.NewDecoder(file)
	params := &Parameters{}
	err = decoder.Decode(params)
	if err != nil {
		return nil, fmt.Errorf("while decoding parameter file: %s", err.Error())
	}

	err = file.Close()
	if err != nil {
		return nil, fmt.Errorf("while closing parameter file: %s", err.Error())
	}
	return params, nil
}

// String inplements the Stringer interface
func (p *Parameters) String() string {

	modifierString := ""
	if startMod := p.Modifier; startMod != nil {
		modifierString = fmt.Sprintf("{ModifierKey:%s AppliedPath:%s}", startMod.ModifierKey, startMod.AppliedPath)
	}

	return fmt.Sprintf("{CohortName:%s ConceptPath:%s Modifier:%s}",
		p.CohortName,
		p.ConceptPath,
		modifierString)
}
