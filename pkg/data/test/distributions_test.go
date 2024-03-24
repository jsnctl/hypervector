package test

import (
	"github.com/jsnctl/hypervector/pkg/data"
	"github.com/jsnctl/hypervector/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGaussian(t *testing.T) {
	// repeating tests against central limit theorem
	numberOfTests := 25
	testMeans := make([]float64, numberOfTests)

	for i := 0; i < numberOfTests; i++ {
		opts := data.DistributionOpts{
			0, 5000, 0.0, 1.0,
		}
		distribution := data.Gaussian(opts)
		mean := helpers.Mean(distribution)

		testMeans = append(testMeans, mean)
	}

	meanOfMeans := helpers.Mean(&testMeans)

	assert.True(t, helpers.IsApproxEqual(0, meanOfMeans, 1e-2))
}
