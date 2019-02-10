package model

// ProjectSummary ...
type ProjectSummary struct {
	Name   string         `json:"name"`
	Builds []BuildHistory `json:"builds"`
}
