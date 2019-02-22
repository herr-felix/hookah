package main

import "./model"

// BuildHistoryStore ...
type BuildHistoryStore interface {
	SaveBuild(data *model.BuildHistory) error
	GetLatestBuilds() ([]*model.BuildHistory, error)
	GetAllBuilds(projectName string) ([]*model.BuildHistory, error)
	GetBuild(ID string) (*model.BuildHistory, error)
}
