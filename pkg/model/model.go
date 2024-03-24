package model

import (
	"github.com/google/uuid"
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

type Definition struct {
	ID       uuid.UUID `json:"id"`
	Name     string
	Added    time.Time
	Features []*Feature `json:"features"`
	N        int        `json:"N"`
}

func NewDefinition(name string) *Definition {
	definition := Definition{Name: name}
	definition.Added = time.Now()
	definition.ID = uuid.New()
	return &definition
}

type Feature struct {
	ID           uuid.UUID    `json:"id"`
	Type         FeatureType  `json:"type"`
	Distribution Distribution `json:"distribution"`
}

func NewFeature(featureType FeatureType, distribution Distribution) *Feature {
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

type Distribution struct {
	Type       FeatureType            `json:"type"`
	Parameters map[string]interface{} `json:"parameters"`
}
