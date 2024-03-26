package test

import (
	"github.com/jsnctl/hypervector/pkg/data"
	"github.com/jsnctl/hypervector/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefinition(t *testing.T) {
	definition := model.NewDefinition("test")
	definition.N = 100
	featureA := model.Feature{
		Type: model.FloatFeature,
		Distribution: data.Distribution{
			Type: data.GaussianType,
			Parameters: data.DistributionOpts{
				Sigma: 10.0,
				Mu:    1.0,
			},
		},
	}
	featureB := model.Feature{
		Type: model.IntegerFeature,
		Distribution: data.Distribution{
			Type: data.GaussianType,
			Parameters: data.DistributionOpts{
				Sigma: 1.0,
				Mu:    10.0,
			},
		},
	}
	definition.Features = []*model.Feature{&featureA, &featureB}

	results := definition.Generate()

	assert.IsType(t, &model.Definition{}, definition)
	assert.NotNil(t, results)

	x, y := results.Shape()
	assert.Equal(t, 100, x)
	assert.Equal(t, 2, y)

	assert.IsType(t, float64(0), (*results)[0][0])
	assert.IsType(t, int(0.0), (*results)[0][1])
}
