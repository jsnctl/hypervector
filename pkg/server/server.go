package server

import (
	"fmt"
	"github.com/jsnctl/hypervector/pkg/data"
	"github.com/jsnctl/hypervector/pkg/model"
	"log"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", dummyHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	// taken from test case for now
	definition := model.NewDefinition("test")
	definition.N = 100
	featureA := model.Feature{
		Type: model.FloatFeature,
		Distribution: model.Distribution{
			Type: data.Gaussian,
			Parameters: map[string]any{
				"sigma": 10.0,
				"mu":    1.0,
			},
		},
	}
	featureB := model.Feature{
		Type: model.IntegerFeature,
		Distribution: model.Distribution{
			Type: data.Gaussian,
			Parameters: map[string]any{
				"sigma": 1.0,
				"mu":    10.0,
			},
		},
	}
	definition.Features = []*model.Feature{&featureA, &featureB}

	results := definition.Generate()
	fmt.Fprintf(w, fmt.Sprint(results))
}
