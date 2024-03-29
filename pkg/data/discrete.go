package data

import "math/rand"

type Category struct {
	Value       string
	Probability float64
}

func CategoryChoice(N int, opts DistributionOpts) *Result {
	rng := rand.New(rand.NewSource(opts.Seed))
	var result Result

	for i := 0; i < N; i++ {
		choice := opts.Categories[rng.Intn(len(opts.Categories))]
		result.Values = append(result.Values, choice.Value)
	}

	return &result
}

func DiscreteGaussian(N int, opts DistributionOpts) *Result {
	continuous := Gaussian(N, opts)
	for i, value := range continuous.Values {
		continuous.Values[i] = int(value.(float64))
	}
	return continuous
}
