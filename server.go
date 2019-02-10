package main

import (
	"log"

	"github.com/labstack/echo"
)

// Server ...
type Server struct {
	BuildSpace model.BuildingSpace
	Store      BuildHistoryStore
}

// NewServer ...
func NewServer(buildSpace model.BuildingSpace) *Server {
	return &Server{
		BuildSpace: buildSpace,
	}
}

// Listen ...
func (s *Server) Listen() {
	e := echo.New()

	log.Println("Running on :6666")
	e.Logger.Fatal(e.Start(":6666"))
}
