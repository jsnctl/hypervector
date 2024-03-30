package model

import (
	"github.com/google/uuid"
	"time"
)

type Vector [][]any

func (r *Vector) Shape() (int, int) {
	return len(*r), len((*r)[0])
}

type VectorResult struct {
	EnsembleId   string  `json:"ensembleId""`
	DefinitionId string  `json:"definitionId"`
	Vector       *Vector `json:"vector"`
}

// Definition contains all metadata and statistical information to
// generate a test data fixture
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
