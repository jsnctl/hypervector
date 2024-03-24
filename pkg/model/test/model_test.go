package test

import (
	"github.com/jsnctl/hypervector/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModel(t *testing.T) {
	project := model.NewProject("test")
	assert.IsType(t, &model.Project{}, project)

	definition := model.NewDefinition("test")
	assert.IsType(t, &model.Definition{}, definition)
}
