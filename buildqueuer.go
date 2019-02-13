package main

import (
	"log"
	"time"

	"./model"
)

// BuildQueuer ...
type BuildQueuer struct {
	buildSpace model.BuildingSpace
	requests   []*model.BuildRequest
}

// NewBuildQueuer ...
func NewBuildQueuer(buildspace model.BuildingSpace) *BuildQueuer {
	return &BuildQueuer{
		buildSpace: buildspace,
		requests:   []*model.BuildRequest{},
	}
}

// Start ...
func (q *BuildQueuer) Start(onBuildEnd func(*model.BuildHistory)) {
	buildLogger := make(chan *model.BuildHistory, 1)

	go func(logger chan *model.BuildHistory) {
		for history := range logger {
			log.Println("Build logged")
			log.Println(history)
			onBuildEnd(history)
		}
	}(buildLogger)
	// Any initializations to builder should come here
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		if len(q.requests) == 0 {
			continue
		}
		var req *model.BuildRequest
		req, q.requests = q.requests[0], q.requests[1:]
		go func(r model.BuildRequest) {
			log.Println("Build started")
			history, err := q.buildSpace.Make(r, "./buildingspaces/handlers") // Handler path more dynamic
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Build done")
			buildLogger <- history
			log.Println("Build done")
		}(*req)
	}
}

// Queue schedule a BuildRequest
func (q *BuildQueuer) Queue(request *model.BuildRequest) int {
	// Any verification?
	q.requests = append(q.requests, request)
	return len(q.requests)
}
