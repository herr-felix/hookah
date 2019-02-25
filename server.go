package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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
func wantsJSON(r *http.Request) bool {
	return strings.HasPrefix(r.Header.Get("Accept"), "application/json")
}

func (s *Server) setupRoutes() chi.Router {

	r := chi.NewRouter()

	views := loadViews("./ui/views/")
	views.Register("root")
	views.Register("project")

	// POST : Build Request
	r.Post("/build", func(w http.ResponseWriter, r *http.Request) {
		// Decode the body to a model.BuildRequest
		var buildRequest *model.BuildRequest

		if err := json.NewDecoder(r.Body).Decode(&buildRequest); err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Could not bind the body to a build request:\n%s", err)
			return
		}
		defer r.Body.Close()

		todo := s.Queuer.Queue(buildRequest)
		log.Println("Build requested")
		fmt.Fprint(w, todo)
	})

	// GET : Latest Builds
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		builds, err := s.Store.GetLatestBuilds()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error while retrieving the latest builds:\n%s", err)
		}
		if wantsJSON(r) {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(builds)
			return
		}
		w.Header().Add("Content-Type", "text/html")
		views.Render(w, "root", builds)
	})

	// GET : All Builds by project
	r.Get("/project/{project}", func(w http.ResponseWriter, r *http.Request) {
		projectName := chi.URLParam(r, "project")
		builds, err := s.Store.GetAllBuilds(projectName)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "No builds with the projectName '%s'\n%s", projectName, err)
			return
		}
		if wantsJSON(r) {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(builds)
			return
		}
		if len(builds) == 0 {
			w.WriteHeader(500)
			fmt.Fprint(w, "Project not found")
			return
		}
		w.Header().Add("Content-Type", "text/html")
		views.Render(w, "project", builds)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/dist/"))).ServeHTTP(w, r)
	})

	return r
}

// Listen ...
func (s *Server) Listen() {

	r := s.setupRoutes()

	go s.Queuer.Start(func(history *model.BuildHistoryItem) {
		err := s.Store.SaveBuild(history)
		if err != nil {
			log.Println(err)
			return
		}
		// Do any kind of notification here
	})

	log.Println("Running on :8080")
	panic(http.ListenAndServe(":8080", r))
}
