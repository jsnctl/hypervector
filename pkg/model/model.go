package model

import (
	"github.com/google/uuid"
	"github.com/jsnctl/hypervector/pkg/data"
	"time"
)

type Project struct {
	ID          uuid.UUID `json:"id"`
	Name        string
	Added       time.Time
	Definitions []*Definition `json:"definitions"`
}

func NewProject(name string) *Project {
	project := Project{Name: name}
	project.Added = time.Now()
	project.ID = uuid.New()
	return &project
}

type Feature struct {
	ID           uuid.UUID         `json:"id"`
	Type         FeatureType       `json:"type"`
	Distribution data.Distribution `json:"distribution"`
}

func NewFeature(featureType FeatureType, distribution data.Distribution) *Feature {
	feature := Feature{Type: featureType, Distribution: distribution}
	feature.ID = uuid.New()
	return &feature
}

type FeatureType string

const (
	IntegerFeature FeatureType = "INTEGER"
	FloatFeature               = "FLOAT"
	StringFeature              = "STRING"
)
