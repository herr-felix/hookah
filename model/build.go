package model

// BuildStatus is either 'successful' or 'failed'
type BuildStatus string

// The 2 possible build status
const (
	SuccessfulBuild BuildStatus = "successful"
	FailedBuild     BuildStatus = "failed"
)

// BuildHistory ..
type BuildHistory struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	ProjectName string      `json:"projectName,omitempty"`
	Start       int64       `json:"start"`
	Duration    int64       `json:"duration"`
	Status      BuildStatus `json:"status"`
	Output      string      `json:"output,omitempty"`
}

// BuildRequest ...
type BuildRequest struct {
	BuildPath   string // Where the building must happen, relative to the project's root. "." Most of the time
	ProjectName string
}
