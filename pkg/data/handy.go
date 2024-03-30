package data

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

var EqualWeightBoolean = Distribution{
	Type: CategoricalType,
	Parameters: DistributionOpts{
		Categories: []Category{
			{"True", 0.5},
			{"False", 0.5},
		},
	},
}
