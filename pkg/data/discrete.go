package data

import "math/rand"

type Category struct {
	Value       string
	Probability float64
}

func CategoryChoice(N int, seed int64, opts DistributionOpts) *Result {
	rng := rand.New(rand.NewSource(seed))
	var result Result

	choiceSlice := generateChoiceSlice(opts.Categories)
	for i := 0; i < N; i++ {
		choice := choiceSlice[rng.Intn(len(choiceSlice))]
		result.Values = append(result.Values, choice.Value)
	}

	return &result
}

func generateChoiceSlice(choices []Category) []Category {
	var choiceSlice []Category
	for _, choice := range choices {
		nSlice := int(choice.Probability * 100)
		for j := 0; j < nSlice; j++ {
			choiceSlice = append(choiceSlice, choice)
		}
	}
	return choiceSlice
}

func DiscreteGaussian(N int, seed int64, opts DistributionOpts) *Result {
	continuous := Gaussian(N, seed, opts)
	for i, value := range continuous.Values {
		continuous.Values[i] = int(value.(float64))
	}
	return continuous
}
