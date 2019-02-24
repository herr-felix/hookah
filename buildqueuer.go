package main

import (
	"log"
	"strings"
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
func (q *BuildQueuer) Start(onBuildEnd func(*model.BuildHistoryItem)) {
	buildLogger := make(chan *model.BuildHistoryItem, 1)

	go func(logger chan *model.BuildHistoryItem) {
		for history := range logger {
			log.Println("Build", history.ID, "completed")
			onBuildEnd(history)
		}
	}(buildLogger)

	// Any initializations to builder should come here

	ticker := time.NewTicker(time.Second)

	for range ticker.C {

		// No requests in queue ?
		if len(q.requests) == 0 {
			continue
		}

		var req *model.BuildRequest
		// head, tail
		req, q.requests = q.requests[0], q.requests[1:]

		go func(r model.BuildRequest) {

			history, err := q.buildSpace.Make(r, "./handlers")

			if err != nil {
				log.Println(err)
				return
			}

			buildLogger <- history
		}(*req)
	}
}

// Queue schedule a BuildRequest
func (q *BuildQueuer) Queue(request *model.BuildRequest) int {
	// Any verification?
	if request.BuildName == "" {
		request.BuildName = strings.Replace(request.ProjectName, " ", "_", -1) +
			" @ " +
			time.Now().Format("2006-1-2 15:04:05 MST")
	}
	q.requests = append(q.requests, request)
	return len(q.requests)
}
