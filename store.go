package main

import "./model"

// BuildHistoryStore ...
type BuildHistoryStore interface {
	SaveBuild(data *model.BuildHistoryItem) error
	GetLatestBuilds() (model.BuildHistory, error)
	GetAllBuilds(projectName string) (model.BuildHistory, error)
}
