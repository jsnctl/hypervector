package server

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jsnctl/hypervector/pkg/data"
	"github.com/jsnctl/hypervector/pkg/model"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
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
		model.NewFeature(model.FloatFeature, data.IdentityGaussianDistribution),
		model.NewFeature(model.FloatFeature, data.GaussianFactory(0.0, 10.0, data.GaussianType)),
		model.NewFeature(model.IntegerFeature, data.GaussianFactory(10.0, 50.0, data.DiscreteGaussianType)),
		model.NewFeature(model.FloatFeature, data.IdentityGaussianDistribution),
		model.NewFeature(model.FloatFeature, data.GaussianFactory(0.01, 0.5, data.GaussianType)),
		model.NewFeature(model.StringFeature, data.EqualWeightBoolean),
	}
	(*s.Repository).AddDefinition(definition)
	model.NewEnsemble(definition, 1000)
}

func readFromYaml() *YAMLConfig {
	var definitions YAMLConfig
	log.Println("reading from YAML")
	file, err := os.ReadFile("example.yaml")
	if err != nil {
		return nil
	}
	yaml.Unmarshal(file, &definitions)
	definitions.hydrate()
	return &definitions
}

func (y *YAMLConfig) hydrate() {
	for _, definition := range y.Definitions {
		definition.ID = uuid.New()
		definition.Added = time.Now()
		for i, ensemble := range definition.Ensembles {
			ensemble.HydrateEnsemble(definition, i)
		}
	}
}

type YAMLConfig struct {
	Definitions []*model.Definition `yaml:"definitions"`
}

func (s *Server) RunServer() {
	log.Println("hypervector is up!")
	yamlData := readFromYaml()
	if yamlData == nil {
		log.Println("reading from YAML failed - using bootstrapped data")
		s.bootstrapData()
	} else {
		(*s.Repository).Overwrite(yamlData.Definitions)
	}
	http.Handle("/definitions", allDefinitionsHandler(s.Repository))
	http.Handle("/definition", definitionHandler(s.Repository))
	http.Handle("/ensemble/*", ensembleHandler(s.Repository))
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
			js, err := json.Marshal(definition)
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

func ensembleHandler(repo *Repository) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			log.Println(r.URL.Path)
			ensemble, err := (*repo).GetEnsemble(r.URL.Query().Get("id"), 0)
			if err != nil {
				println(err.Error())
				fmt.Fprintf(w, err.Error())
				return
			}
			toReturn := model.VectorResult{
				EnsembleId:   ensemble.ID,
				DefinitionId: ensemble.DefinitionID.String(),
				N:            ensemble.N,
				Vector:       ensemble.Generate(),
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
