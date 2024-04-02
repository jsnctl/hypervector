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
	GetEnsemble(parentDefinition string, ensembleId int) (*model.Ensemble, error)

	Overwrite([]*model.Definition)
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

func (r *InMemoryRepository) Overwrite(newDefinitions []*model.Definition) {
	for k := range r.Definitions {
		delete(r.Definitions, k)
	}
	for _, definition := range newDefinitions {
		for _, feature := range definition.Features {
			if feature.ID == uuid.Nil {
				feature.ID = uuid.New()
			}
		}
		r.AddDefinition(definition)
	}
}

func (r *InMemoryRepository) AddDefinition(definition *model.Definition) {
	if definition.ID == uuid.Nil {
		definition.ID = uuid.New()
	}
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
	parent, _ := r.GetDefinition(ensemble.DefinitionID.String())
	parent.Ensembles = append(parent.Ensembles, ensemble)
}

func (r *InMemoryRepository) GetEnsemble(parentDefinition string, ensembleId int) (*model.Ensemble, error) {
	definition, err := r.GetDefinition(parentDefinition)
	if err != nil {
		return nil, errors.New("definition of ensemble not found")
	}
	ensemble := definition.Ensembles[ensembleId]
	return ensemble, nil
}
