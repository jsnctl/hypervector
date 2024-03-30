package data

import (
	"math"
	"math/rand"
)

func Gaussian(N int, opts DistributionOpts) *Result {
	// box-muller transform
	rng := rand.New(rand.NewSource(opts.Seed))
	var result Result
	for i := 0; i < int(N/2); i++ {
		left := rng.Float64()
		right := rng.Float64()
		x1 := math.Sqrt(-2.0*math.Log(left)) * math.Cos(2.0*math.Pi*right)
		x2 := math.Sqrt(-2.0*math.Log(left)) * math.Sin(2.0*math.Pi*right)

		result.Values = append(result.Values, x1*opts.Sigma+opts.Mu)
		result.Values = append(result.Values, x2*opts.Sigma+opts.Mu)
	}
	return &result
}

func GaussianFactory(mu float64, sigma float64, gaussianType DistributionType) Distribution {
	return Distribution{
		Type: gaussianType,
		Parameters: DistributionOpts{
			Mu:    mu,
			Sigma: sigma,
		},
	}
}

var IdentityGaussianDistribution = GaussianFactory(0.0, 1.0, GaussianType)
