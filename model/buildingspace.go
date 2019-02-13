package model

// BuildingSpace is driver that can execute a build request
type BuildingSpace interface {
	Make(req BuildRequest, handlersPath string) (*BuildHistory, error)
}
