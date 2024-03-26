package data

import (
	"math"
	"math/rand"
)

type DistributionType string

const (
	GaussianType DistributionType = "GAUSSIAN"
)

type Distribution struct {
	Type       DistributionType `json:"type"`
	Parameters DistributionOpts `json:"parameters"`
}

type DistributionOpts struct {
	Seed  int     `json:"seed"`
	Mu    float64 `json:"mu"`
	Sigma float64 `json:"sigma"`
}

var DistributionLookup = map[DistributionType]func(int, DistributionOpts) *[]float64{
	GaussianType: Gaussian,
}

func Gaussian(N int, opts DistributionOpts) *[]float64 {
	// box-muller transform
	var values []float64
	for i := 0; i < int(N/2); i++ {
		left := rand.Float64()
		right := rand.Float64()
		x1 := math.Sqrt(-2.0*math.Log(left)) * math.Cos(2.0*math.Pi*right)
		x2 := math.Sqrt(-2.0*math.Log(left)) * math.Sin(2.0*math.Pi*right)

		values = append(values, x1*opts.Sigma+opts.Mu)
		values = append(values, x2*opts.Sigma+opts.Mu)
	}
	return &values
}

var IdentityGaussianDistribution = Distribution{
	Type: GaussianType,
	Parameters: DistributionOpts{
		Sigma: 1.0,
		Mu:    0.0,
	},
}
