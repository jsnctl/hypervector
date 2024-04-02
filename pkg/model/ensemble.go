package model

import (
	"github.com/google/uuid"
	"github.com/jsnctl/hypervector/pkg/data"
	"hash/fnv"
	"strconv"
)

// Ensemble corresponds to a single test fixture instance of
// a Definition. A given Ensemble returns the same data every time
type Ensemble struct {
	ID           int `json:"id"`
	definition   *Definition
	DefinitionID uuid.UUID `json:"-"`
	N            int       `json:"N"`
}

func NewEnsemble(definition *Definition, N int) *Ensemble {
	ensemble := Ensemble{
		definition:   definition,
		DefinitionID: definition.ID,
		N:            N,
	}
	definition.Ensembles = append(definition.Ensembles, &ensemble)
	ensemble.ID = len(definition.Ensembles) + 1
	return &ensemble
}

func (e *Ensemble) HydrateEnsemble(parent *Definition, index int) {
	e.definition = parent
	e.DefinitionID = parent.ID
	e.ID = index
}

func (e *Ensemble) Generate() *Vector {
	fv := make(Vector, e.N)
	for f, feature := range e.definition.Features {
		fn := data.DistributionLookup[feature.Distribution.Type]
		seed := ensembleToRNGSeed(e)
		distribution := fn(e.N, seed, feature.Distribution.Parameters)
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

func ensembleToRNGSeed(ensemble *Ensemble) int64 {
	// hashes increment ID and N as deterministic seed
	stringToHash := strconv.Itoa(ensemble.ID) + "-" + strconv.Itoa(ensemble.N)
	hash := fnv.New64a()
	hash.Write([]byte(stringToHash))
	return int64(hash.Sum64())
}
