package test

import (
	"github.com/jsnctl/hypervector/pkg/data"
	"github.com/jsnctl/hypervector/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func DefinitionFixture() *model.Definition {
	definition := model.NewDefinition("test")
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
			Type: data.DiscreteGaussianType,
			Parameters: data.DistributionOpts{
				Sigma: 1.0,
				Mu:    10.0,
			},
		},
	}
	definition.Features = []*model.Feature{&featureA, &featureB}
	return definition
}

func TestDefinition(t *testing.T) {
	definition := DefinitionFixture()
	assert.Equal(t, len(definition.Features), 2)
}
