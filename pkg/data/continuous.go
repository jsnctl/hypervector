package data

import (
	"math"
	"math/rand"
)

func Gaussian(N int, seed int64, opts DistributionOpts) *Result {
	// box-muller transform
	rng := rand.New(rand.NewSource(seed))
	var result Result
	for i := 0; i < int(N/2); i++ {
		left := rng.Float64()
		right := rng.Float64()
		x1 := math.Sqrt(-2.0*math.Log(left)) * math.Cos(2.0*math.Pi*right)
		x2 := math.Sqrt(-2.0*math.Log(left)) * math.Sin(2.0*math.Pi*right)

		result.Values = append(result.Values, x1*opts.Sigma+opts.Mu)
		result.Values = append(result.Values, x2*opts.Sigma+opts.Mu)
	}

	// occasionally above N/2 shortens result.Values
	if len(result.Values) != N {
		randomResample := result.Values[rand.Intn(N)]
		result.Values = append(result.Values, randomResample)
	}

	return &result
}
