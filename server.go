package main

import (
	"fmt"
	"log"

	"./model"
	"github.com/labstack/echo"
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
	e := echo.New()

	// POST : Build Request
	e.POST("/build", func(c echo.Context) error {
		// Decode the body to a model.BuildRequest
		buildRequest := new(model.BuildRequest)
		if err := c.Bind(buildRequest); err != nil {
			return c.String(400, "Could not bind the body to a build request.")
		}
		todo := s.Queuer.Queue(buildRequest)
		log.Println("Build requested")
		return c.String(200, string(todo))
	})

	// GET : Latest Builds
	e.GET("/builds", func(c echo.Context) error {
		builds, err := s.Store.GetLatestBuilds()
		if err != nil {
			return c.String(500, fmt.Sprintf("Error while retrieving the latest builds: %e", err))
		}
		return c.JSON(200, builds)
	})

	// GET : All Builds by project
	e.GET("/builds/:project", func(c echo.Context) error {
		builds, err := s.Store.GetAllBuilds(c.Param("project"))
		if err != nil {
			return c.String(404, fmt.Sprintf("No builds with the projectName '%s'", c.Param("project")))
		}
		return c.JSON(200, builds)
	})

	// GET : Build's output
	e.GET("/build/output/:buildID", func(c echo.Context) error {
		output, err := s.Store.GetBuildOutput(c.Param("buildID"))
		if err != nil {
			return c.String(404, "No builds found")
		}
		return c.String(200, output)
	})

	go s.Queuer.Start(func(history *model.BuildHistory) {
		err := s.Store.SaveBuild(history)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(history)
		// Do any kind of notification here
	})

	log.Println("Running on :6666")
	e.Logger.Fatal(e.Start(":6666"))
}
