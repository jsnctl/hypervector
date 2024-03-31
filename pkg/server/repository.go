package server

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jsnctl/hypervector/pkg/model"
)

type Repository interface {
	AddDefinition(definition *model.Definition)
	GetDefinition(id string) (*model.Definition, error)
	GetAllDefinitions() []model.Definition

	AddEnsemble(ensemble *model.Ensemble)
	GetEnsemble(id string) (*model.Ensemble, error)

	Overwrite(*[]model.Definition)
}

type InMemoryRepository struct {
	R           Repository
	Definitions map[uuid.UUID]*model.Definition
	Ensembles   map[uuid.UUID]*model.Ensemble
}

func NewInMemoryRepository() *InMemoryRepository {
	inMemoryRepository := InMemoryRepository{}
	inMemoryRepository.Definitions = make(map[uuid.UUID]*model.Definition)
	inMemoryRepository.Ensembles = make(map[uuid.UUID]*model.Ensemble)
	return &inMemoryRepository
}

func (r *InMemoryRepository) Overwrite(newDefinitions *[]model.Definition) {
	for k := range r.Definitions {
		delete(r.Definitions, k)
	}
	for _, definition := range *newDefinitions {
		r.Definitions[definition.ID] = &definition
		ensemble, _ := model.NewEnsemble(&definition, 1000)
		r.AddEnsemble(ensemble)
	}
}

func (r *InMemoryRepository) AddDefinition(definition *model.Definition) {
	r.Definitions[definition.ID] = definition
}

func (r *InMemoryRepository) GetDefinition(id string) (*model.Definition, error) {
	uuidFromString, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id was not a valid UUID")
	}
	definition, ok := r.Definitions[uuidFromString]
	if ok {
		return definition, nil
	}
	return nil, errors.New("definition not found")
}

func (r *InMemoryRepository) GetAllDefinitions() []model.Definition {
	var definitions []model.Definition
	for _, v := range r.Definitions {
		definitions = append(definitions, *v)
	}
	return definitions
}

func (r *InMemoryRepository) AddEnsemble(ensemble *model.Ensemble) {
	r.Ensembles[ensemble.ID] = ensemble
	parent, _ := r.GetDefinition(ensemble.DefinitionID.String())
	parent.Ensembles = append(parent.Ensembles, ensemble.ID)
}

func (r *InMemoryRepository) GetEnsemble(id string) (*model.Ensemble, error) {
	uuidFromString, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("id was not a valid UUID")
	}
	ensemble, ok := r.Ensembles[uuidFromString]
	if ok {
		return ensemble, nil
	}
	return nil, errors.New("definition not found")
}
