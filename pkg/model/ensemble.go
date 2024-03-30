package model

import (
	"github.com/google/uuid"
)

// Ensemble corresponds to a single test fixture instance of
// a Definition. A given Ensemble returns the same data every time
type Ensemble struct {
	ID           uuid.UUID `json:"id"`
	DefinitionID uuid.UUID `json:"definitionId"`
	N            int       `json:"N"`
}

func NewEnsemble(definitionId uuid.UUID, N int) (*Ensemble, error) {
	ensemble := Ensemble{
		DefinitionID: definitionId,
		N:            N,
	}
	ensemble.ID = uuid.New()

	return &ensemble, nil
}
