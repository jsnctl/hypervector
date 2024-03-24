package helpers

import (
	"math"
	"math/rand"
)

func Mean(values *[]float64) float64 {
	total := 0.0
	for _, value := range *values {
		total += value
	}
	return total / float64(len(*values))
}

func IsApproxEqual(expected float64, actual float64, precision float64) bool {
	return math.Abs(expected-actual) <= precision
}

func SeedFixture(N int) []int64 {
	// convenience func to generate list of seeds
	// for rand tests in ensemble (mainly CLT stuff)
	var seeds []int64
	for i := 0; i < N; i++ {
		seeds = append(seeds, int64(rand.Int()))
	}
	return seeds
}
