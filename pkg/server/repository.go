package server

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jsnctl/hypervector/pkg/model"
)

type Repository interface {
	AddDefinition(definition *model.Definition)
	GetDefinition(id uuid.UUID) (*model.Definition, error)
	GetAllDefinitions() []model.Definition
}

type InMemoryRepository struct {
	R           Repository
	Definitions map[uuid.UUID]*model.Definition
}

func NewInMemoryRepository() *InMemoryRepository {
	inMemoryRepository := InMemoryRepository{}
	inMemoryRepository.Definitions = make(map[uuid.UUID]*model.Definition)
	return &inMemoryRepository
}

func (r InMemoryRepository) AddDefinition(definition *model.Definition) {
	r.Definitions[definition.ID] = definition
}

func (r InMemoryRepository) GetDefinition(id uuid.UUID) (*model.Definition, error) {
	definition, ok := r.Definitions[id]
	if ok {
		return definition, nil
	}
	return nil, errors.New("definition not found")
}

func (r InMemoryRepository) GetAllDefinitions() []model.Definition {
	var definitions []model.Definition
	for _, v := range r.Definitions {
		definitions = append(definitions, *v)
	}
	return definitions
}
