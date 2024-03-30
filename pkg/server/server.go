package server

import (
	"encoding/json"
	"fmt"
	"github.com/jsnctl/hypervector/pkg/data"
	"github.com/jsnctl/hypervector/pkg/model"
	"log"
	"net/http"
)

type Server struct {
	Repository *Repository
}

func NewServer(repository Repository) *Server {
	if repository == nil {
		repository = NewInMemoryRepository()
	}
	return &Server{
		Repository: &repository,
	}
}

func (s *Server) bootstrapData() {
	definition := model.NewDefinition("test")
	definition.N = 1000
	definition.Features = []*model.Feature{
		model.NewFeature(model.FloatFeature, data.IdentityGaussianDistribution),
		model.NewFeature(model.FloatFeature, data.GaussianFactory(0.0, 10.0, data.GaussianType)),
		model.NewFeature(model.IntegerFeature, data.GaussianFactory(10.0, 50.0, data.DiscreteGaussianType)),
		model.NewFeature(model.FloatFeature, data.IdentityGaussianDistribution),
		model.NewFeature(model.FloatFeature, data.GaussianFactory(0.01, 0.5, data.GaussianType)),
	}
	(*s.Repository).AddDefinition(definition)
}

func (s *Server) RunServer() {
	s.bootstrapData()
	http.Handle("/definitions", allDefinitionsHandler(s.Repository))
	http.Handle("/definition", definitionHandler(s.Repository))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func allDefinitionsHandler(repo *Repository) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		switch r.Method {
		case http.MethodGet:
			definitions := (*repo).GetAllDefinitions()
			js, err := json.Marshal(definitions)
			if err != nil {
				println(err.Error())
			}
			w.Write(js)
		}
	}
	return http.HandlerFunc(fn)
}

func definitionHandler(repo *Repository) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			definition, err := (*repo).GetDefinition(r.URL.Query().Get("id"))
			if err != nil {
				println(err.Error())
				fmt.Fprintf(w, err.Error())
				return
			}
			toReturn := model.VectorResult{
				Definition: definition.ID.String(),
				Vector:     definition.Generate(),
			}
			js, err := json.Marshal(toReturn)
			if err != nil {
				println(err.Error())
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(js)
		}
	}
	return http.HandlerFunc(fn)
}
