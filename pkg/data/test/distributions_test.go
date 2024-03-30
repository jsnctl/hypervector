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

func TestCategoryChoice(t *testing.T) {
	categories := []data.Category{
		{Value: "True", Probability: 0.2},
		{Value: "False", Probability: 0.8},
	}
	N := 5000 + rand.Intn(1000)
	seed := int64(rand.Int())

	distribution := data.CategoryChoice(N, data.DistributionOpts{
		Seed:       seed,
		Categories: categories,
	})
	vector := distribution.GetString()

	assert.NotNil(t, vector)

	var falseCount int
	for _, value := range *vector {
		if value == "False" {
			falseCount += 1
		}
	}
	propFalse := float64(falseCount) / float64(N)

	assert.True(t, helpers.IsApproxEqual(0.2, 1-propFalse, 0.05))
	assert.True(t, helpers.IsApproxEqual(0.8, propFalse, 0.05))
}
