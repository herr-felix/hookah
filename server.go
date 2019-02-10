package main

import (
	"log"

	"./model"
	"github.com/labstack/echo"
)

// Server ...
type Server struct {
	BuildSpace model.BuildingSpace
	Store      BuildHistoryStore
}

// NewServer ...
func NewServer(buildSpace model.BuildingSpace, store BuildHistoryStore) *Server {
	return &Server{
		BuildSpace: buildSpace,
		Store:      store,
	}
}

// Listen ...
func (s *Server) Listen() {
	e := echo.New()

	// POST : Build Request

	// GET : Latest Builds

	// GET : All Builds by project

	// GET : Build's output

	log.Println("Running on :6666")
	e.Logger.Fatal(e.Start(":6666"))
}
