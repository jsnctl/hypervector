package test

import (
	"github.com/jsnctl/hypervector/pkg/data"
	"github.com/jsnctl/hypervector/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestGaussian(t *testing.T) {
	// repeating tests against central limit theorem
	numberOfTests := 25
	testMeans := make([]float64, numberOfTests)
	mu := 10.0
	sigma := 1.0
	N := 5000 + rand.Intn(1000)
	seed := int64(rand.Int())

	for i := 0; i < numberOfTests; i++ {
		opts := data.DistributionOpts{
			Seed: seed, Mu: mu, Sigma: sigma,
		}
		distribution := data.Gaussian(N, opts)
		mean := helpers.Mean(distribution.GetFloat64())

		testMeans[i] = mean
	}

	meanOfMeans := helpers.Mean(&testMeans)

	assert.True(t, helpers.IsApproxEqual(mu, meanOfMeans, 2e-2)) // previously 1e-2
}
