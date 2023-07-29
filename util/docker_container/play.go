package docker_container

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"os"
)

func StartContainerBlo(ctx context.Context, imageName string, cmdStringArr []string, memoryLimit int64, containerName string) (err error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// start container
	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: imageName,
			Cmd:   cmdStringArr,
		},
		&container.HostConfig{
			Resources: container.Resources{
				//NanoCPUs:  int64(20 * 100 * 1000 * 1000), // 20 * 0.1核心
				Memory: memoryLimit, // 1GB
				//CPUPeriod:  int64(10 * 1000),
				//CPUQuota:   int64(10),
				CpusetCpus: "5-6",
			},
		},
		nil,
		nil,
		containerName,
	)
	if err != nil {
		panic(err)
	}

	// 启动容器
	err = cli.ContainerStart(
		ctx,
		resp.ID,
		types.ContainerStartOptions{},
	)
	if err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, out)
	if err != nil {
		panic(err)
	}

	return nil
}
