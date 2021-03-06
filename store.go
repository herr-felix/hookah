package main

import "./model"

// BuildHistoryStore ...
type BuildHistoryStore interface {
	SaveBuild(data *model.BuildHistoryItem) error
	GetLatestBuilds() (model.BuildHistory, error)
	GetBuilds(projectName string) (model.BuildHistory, error)
	InvalidateBuild(projectName, buildID string) error
}
