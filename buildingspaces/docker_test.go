package buildingspaces

import (
	"log"
	"testing"

	"../internal"
)

func TestDockerBuild(t *testing.T) {
	bs := NewDockerBuildingSpace()

	req := internal.BuildRequest{
		BuildPath:   ".",
		ProjectName: "DemoBuild",
	}

	h, err := bs.Make(req)

	if err != nil {
		t.Error(err)
		return
	}

	log.Println(h.Output)
}
