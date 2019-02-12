package buildingspaces

import (
	"log"
	"testing"

	"../../model"
)

func TestDockerBuildSuccess(t *testing.T) {
	bs := NewDockerBuildingSpace()

	req := model.BuildRequest{
		BuildPath:   "./pass",
		ProjectName: "demobuild",
	}

	h, err := bs.Make(req, "../handlers/")

	if err != nil {
		t.Error(err)
		return
	}

	log.Println(h.Output)

	if h.Status != model.SuccessfulBuild {
		t.Errorf("The build should have passed")
	}
}

func TestDockerBuildFail(t *testing.T) {
	bs := NewDockerBuildingSpace()

	req := model.BuildRequest{
		BuildPath:   "./fail",
		ProjectName: "demobuild",
	}

	h, err := bs.Make(req, "../handlers/")

	if err != nil {
		t.Error(err)
		return
	}

	log.Println(h.Output)

	if h.Status != model.FailedBuild {
		t.Errorf("The build should have failed")
	}

}
