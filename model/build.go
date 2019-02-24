package model

import "sort"

// BuildStatus is either 'successful' or 'failed'
type BuildStatus string

// The 2 possible build status
const (
	SuccessfulBuild BuildStatus = "success"
	FailedBuild     BuildStatus = "failure"
)

// BuildHistoryItem ..
type BuildHistoryItem struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	ProjectName string      `json:"projectName,omitempty"`
	Start       int64       `json:"start"`
	Duration    int64       `json:"duration"`
	Status      BuildStatus `json:"status"`
	Output      string      `json:"output,omitempty"`
}

// BuildHistory ...
type BuildHistory []*BuildHistoryItem

// BuildRequest ...
type BuildRequest struct {
	BuildPath   string `json:"buildPath"` // Where the building must happen, relative to the project's root. "." Most of the time
	ProjectName string `json:"projectName"`
	PullHandler string `json:"pullHandler"`
	PullParams  string `json:"pullParams"`
	PushHandler string `json:"pushHandler"`
	PushParams  string `json:"pushParams"`
}

// OrderByStart Orders all the build history by start. Decreasing.
func (builds BuildHistory) OrderByStart() {
	sorter := &buildHistorySorter{
		builds: builds,
		by: func(a, b *BuildHistoryItem) bool {
			return a.Start > b.Start
		},
	}
	sort.Sort(sorter)
}

type buildHistorySorter struct {
	builds BuildHistory
	by     func(a, b *BuildHistoryItem) bool
}

func (s *buildHistorySorter) Len() int {
	return len(s.builds)
}

func (s *buildHistorySorter) Swap(i, j int) {
	s.builds[i], s.builds[j] = s.builds[j], s.builds[i]
}

func (s *buildHistorySorter) Less(i, j int) bool {
	return s.by(s.builds[i], s.builds[j])
}
