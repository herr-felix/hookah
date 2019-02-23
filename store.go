package main

import "./model"

// BuildHistoryStore ...
type BuildHistoryStore interface {
	SaveBuild(data *model.BuildHistory) error
	GetLatestBuilds() (model.BuildHistories, error)
	GetAllBuilds(projectName string) (model.BuildHistories, error)
}
