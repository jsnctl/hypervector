package server

import (
	"encoding/json"
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
	(*s.Repository).AddDefinition(model.NewDefinition("test"))
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
			json.NewEncoder(w).Encode(definitions)
		}
	}
	return http.HandlerFunc(fn)
}
