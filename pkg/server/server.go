package server

import (
	"encoding/json"
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
	definition.Features = []*model.Feature{
		model.NewFeature(model.IntegerFeature, data.IdentityGaussianDistribution),
		model.NewFeature(model.FloatFeature, data.IdentityGaussianDistribution),
		model.NewFeature(model.IntegerFeature, data.IdentityGaussianDistribution),
	}
	(*s.Repository).AddDefinition(definition)
}

func (s *Server) RunServer() {
	s.bootstrapData()
	http.Handle("/definitions", definitionHandler(s.Repository))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func definitionHandler(repo *Repository) http.Handler {
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
