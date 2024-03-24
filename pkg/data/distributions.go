package data

import (
	"math"
	"math/rand"
)

type DistributionOpts struct {
	Seed  int64
	N     int64
	Mu    float64
	Sigma float64
}

func Gaussian(opts DistributionOpts) *[]float64 {
	// box-muller transform

	var values []float64
	for i := 0; i < int(opts.N/2); i++ {
		left := rand.Float64()
		right := rand.Float64()
		x1 := math.Sqrt(-2.0*math.Log(left)) * math.Cos(2.0*math.Pi*right)
		x2 := math.Sqrt(-2.0*math.Log(left)) * math.Sin(2.0*math.Pi*right)

		values = append(values, x1*opts.Sigma+opts.Mu)
		values = append(values, x2*opts.Sigma+opts.Mu)
	}
	return &values
}
