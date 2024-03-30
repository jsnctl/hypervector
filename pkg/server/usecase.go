package server

import "github.com/jsnctl/hypervector/pkg/model"

func AddEnsembleToDefinition(repo *Repository, definitionId string, N int) {
	definition, _ := (*repo).GetDefinition(definitionId)
	ensemble, _ := model.NewEnsemble(definition.ID, N)
	(*repo).AddEnsemble(ensemble)
}
