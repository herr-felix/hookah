package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
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

type views struct {
	templates map[string]*template.Template
	root      []string
}

func (v *views) Register(name string) {
	v.templates[name] = template.Must(template.New("").ParseFiles(
		filepath.Join(append(v.root, "layout.html")...),
		filepath.Join(append(v.root, name+".html")...),
	))
}

func (v *views) Render(w io.Writer, name string, data interface{}) {
	if tmpl, exists := v.templates[name]; exists {
		tmpl.ExecuteTemplate(w, "base", data)
		return
	}
	fmt.Fprintf(w, "Could not find template '%s'", name)
}

func loadViews(root string) *views {
	return &views{
		templates: make(map[string]*template.Template),
		root:      strings.Split(root, "/"),
	}
}

func wantsJSON(r *http.Request) bool {
	return strings.HasPrefix(r.Header.Get("Accept"), "application/json")
}

func (s *Server) setupRoutes() chi.Router {

	r := chi.NewRouter()

	views := loadViews("./ui/views/")
	views.Register("all_builds")
	views.Register("project_builds")

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
		if wantsJSON(r) {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(builds)
			return
		}
		w.Header().Add("Content-Type", "text/html")

		views.Render(w, "all_builds", builds)
	})

	// GET : All Builds by project
	r.Get("/builds/{project}", func(w http.ResponseWriter, r *http.Request) {
		builds, err := s.Store.GetAllBuilds(chi.URLParam(r, "project"))
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "No builds with the projectName '%s'", chi.URLParam(r, "project"))
			return
		}
		if wantsJSON(r) {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(builds)
			return
		}
		w.Header().Add("Content-Type", "text/html")
		views.Render(w, "project_builds", builds)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/dist/"))).ServeHTTP(w, r)
	})

	return r
}

// Listen ...
func (s *Server) Listen() {

	r := s.setupRoutes()

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
