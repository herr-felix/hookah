package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./model"
	"github.com/go-chi/chi"
)

// Server ...
type Server struct {
	Queuer *BuildQueuer
	Store  BuildHistoryStore
}

// NewServer ...
func NewServer(queuer *BuildQueuer, store BuildHistoryStore) *Server {
	return &Server{
		Queuer: queuer,
		Store:  store,
	}
}

// Listen ...
func (s *Server) Listen() {
	r := chi.NewRouter()

	// POST : Build Request
	r.Post("/build", func(w http.ResponseWriter, r *http.Request) {
		// Decode the body to a model.BuildRequest
		var buildRequest *model.BuildRequest

		if err := json.NewDecoder(r.Body).Decode(buildRequest); err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Could not bind the body to a build request."))
			return
		}
		defer r.Body.Close()
		todo := s.Queuer.Queue(buildRequest)
		log.Println("Build requested")
		w.Write([]byte(string(todo)))
	})

	// GET : Latest Builds
	r.Get("/builds", func(w http.ResponseWriter, r *http.Request) {
		builds, err := s.Store.GetLatestBuilds()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("Error while retrieving the latest builds: %e", err)))
		}
		json.NewEncoder(w).Encode(builds)
	})

	// GET : All Builds by project
	r.Get("/builds/{project}", func(w http.ResponseWriter, r *http.Request) {
		builds, err := s.Store.GetAllBuilds(chi.URLParam(r, "project"))
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, "No builds with the projectName '%s'", chi.URLParam(r, "project"))
		}
		json.NewEncoder(w).Encode(builds)
	})

	// GET : Build's output
	r.Get("/build/output/{buildID}", func(w http.ResponseWriter, r *http.Request) {
		output, err := s.Store.GetBuildOutput(chi.URLParam(r, "buildID"))
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprint(w, "No builds found")
		}
		fmt.Fprint(w, output)
	})

	r.Handle("/static", http.FileServer(http.Dir("/ui/dist")))

	go s.Queuer.Start(func(history *model.BuildHistory) {
		err := s.Store.SaveBuild(history)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(history)
		// Do any kind of notification here
	})

	log.Println("Running on :8080")
	panic(http.ListenAndServe(":8080", r))
}
