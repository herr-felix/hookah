package main

import "./model"

// BuildHistoryStore ...
type BuildHistoryStore interface {
	GetSummaries() ([]model.ProjectSummary, error)
	GetBuild(ID string) (*model.BuildHistory, error)
	SaveBuild(data *model.BuildHistory) error
}
