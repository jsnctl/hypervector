package model

import (
	"github.com/google/uuid"
	"github.com/jsnctl/hypervector/pkg/data"
	"time"
)

type Vector [][]any

func (r *Vector) Shape() (int, int) {
	return len(*r), len((*r)[0])
}

type VectorResult struct {
	Definition string  `json:"definitionId"`
	Vector     *Vector `json:"vector"`
}

type Definition struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Added     time.Time   `json:"added"`
	N         int         `json:"N"`
	Features  []*Feature  `json:"features"`
	Ensembles []uuid.UUID `json:"ensembles"`
}

func NewDefinition(name string) *Definition {
	definition := Definition{Name: name}
	definition.Added = time.Now()
	definition.ID = uuid.New()
	return &definition
}

func (d *Definition) Generate() *Vector {
	fv := make(Vector, d.N)
	for f, feature := range d.Features {
		fn := data.DistributionLookup[feature.Distribution.Type]
		distribution := fn(d.N, feature.Distribution.Parameters)
		for i, value := range distribution.Values {
			if f == 0 {
				fv[i] = make([]interface{}, len(d.Features))
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
