package buildingspaces

import (
	"bytes"
	"context"
	"time"

	"../../model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// Docker ...
type Docker struct{}

// NewDockerBuildingSpace ...
func NewDockerBuildingSpace() *Docker {
	return &Docker{}
}

// Make execute the build request
func (dbs *Docker) Make(req model.BuildRequest) (*model.BuildHistory, error) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	cli.NegotiateAPIVersion(ctx)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "test",
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/var/run/docker.sock",
				Target: "/var/run/docker.sock",
			},
		},
	}, nil, "")
	if err != nil {
		return nil, err
	}
	defer cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true})

	startTime := time.Now()

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}

	buildStatus := model.SuccessfulBuild // Default to success

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return nil, err
		}
	case status := <-statusCh:
		if status.StatusCode != 0 {
			buildStatus = model.FailedBuild
		}
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer([]byte{})
	stdcopy.StdCopy(buffer, buffer, out)

	return &model.BuildHistory{
		ID:          resp.ID, // sha256 of this?
		Name:        "MAKE BUILD NAME DYNAMIC!",
		ProjectName: req.ProjectName,
		Start:       startTime.Unix(),
		Duration:    int64(time.Now().Sub(startTime).Seconds()),
		Status:      buildStatus,
		Output:      string(buffer.Bytes()),
	}, nil
}