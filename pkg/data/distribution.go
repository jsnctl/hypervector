package data

type DistributionType string

type Distribution struct {
	Type       DistributionType `json:"type"`
	Parameters DistributionOpts `json:"parameters"`
}

const (
	GaussianType         DistributionType = "GAUSSIAN"
	DiscreteGaussianType DistributionType = "DISCRETE_GAUSSIAN"
	CategoricalType      DistributionType = "CATEGORICAL"
)

type DistributionOpts struct {
	Seed       int64   `json:"seed"`
	Mu         float64 `json:"mu"`
	Sigma      float64 `json:"sigma"`
	Categories []Category
}

type Result struct {
	Values []any
}

func (r *Result) GetFloat64() *[]float64 {
	floats := make([]float64, len(r.Values))
	for i, value := range r.Values {
		floats[i] = value.(float64)
	}
	return &floats
}

var DistributionLookup = map[DistributionType]func(int, DistributionOpts) *Result{
	GaussianType:         Gaussian,
	DiscreteGaussianType: DiscreteGaussian,
	CategoricalType:      CategoryChoice,
}
