package model

import (
	"github.com/google/uuid"
	"github.com/jsnctl/hypervector/pkg/data"
)

// Ensemble corresponds to a single test fixture instance of
// a Definition. A given Ensemble returns the same data every time
type Ensemble struct {
	ID           uuid.UUID `json:"id"`
	definition   *Definition
	DefinitionID uuid.UUID `json:"definitionId"`
	N            int       `json:"N"`
}

func NewEnsemble(definition *Definition, N int) (*Ensemble, error) {
	ensemble := Ensemble{
		definition:   definition,
		DefinitionID: definition.ID,
		N:            N,
	}
	ensemble.ID = uuid.New()

	return &ensemble, nil
}

func (e *Ensemble) Generate() *Vector {
	fv := make(Vector, e.N)
	for f, feature := range e.definition.Features {
		fn := data.DistributionLookup[feature.Distribution.Type]
		distribution := fn(e.N, feature.Distribution.Parameters)
		for i, value := range distribution.Values {
			if f == 0 {
				fv[i] = make([]interface{}, len(e.definition.Features))
			}
			if feature.Type == IntegerFeature {
				fv[i][f] = value.(int)
			} else if feature.Type == FloatFeature {
				fv[i][f] = value.(float64)
			} else if feature.Type == StringFeature {
				fv[i][f] = value.(string)
			}
		}
	}
	return &fv
}
