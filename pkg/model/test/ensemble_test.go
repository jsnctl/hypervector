package test

import (
	"github.com/jsnctl/hypervector/pkg/model"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestEnsemble(t *testing.T) {
	N := rand.Intn(5000)
	definition := DefinitionFixture()
	ensemble, _ := model.NewEnsemble(definition, N)

	results := ensemble.Generate()

	assert.IsType(t, &model.Definition{}, definition)
	assert.NotNil(t, results)

	x, y := results.Shape()
	assert.Equal(t, N, x)
	assert.Equal(t, len(definition.Features), y)

	assert.IsType(t, float64(0), (*results)[0][0])
	assert.IsType(t, int(0.0), (*results)[0][1])
}
