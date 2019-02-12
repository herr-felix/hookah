package buildingspaces

import (
	"log"
	"testing"

	"../../model"
)

func TestDockerBuild(t *testing.T) {
	bs := NewDockerBuildingSpace()

	req := model.BuildRequest{
		BuildPath:   ".",
		ProjectName: "demobuild",
	}

	h, err := bs.Make(req)

	if err != nil {
		t.Error(err)
		return
	}

	log.Println(h.Output)
}
