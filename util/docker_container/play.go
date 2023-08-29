package docker_container

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

func StartContainerBlo(ctx context.Context, imageName string, cmdStringArr []string, memLim int64, cpuPercent int, cpuSetCpus string, containerName string) (err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
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
				Memory:     memLim * 1000 * 1000, // MB
				CPUPeriod:  int64(100 * 1000),
				CPUQuota:   int64(cpuPercent * 1000),
				CpusetCpus: cpuSetCpus,
			},
		},
		nil,
		nil,
		containerName,
	)
	if err != nil {
		panic(err)
	}

	log.Print("ContainerCreate resp: ", resp)

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

func StartContainerNBlo(ctx context.Context, imageName string, cmdStringArr []string, memLim int64, cpuPercent int, cpuSetCpus string, containerName string, volumeName string, pathInCont string) (err error, containerId string) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	var binds = []string{}
	if volumeName != "" {
		binds = []string{fmt.Sprintf("%s:%s", volumeName, pathInCont)}
	}

	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: imageName,
			Cmd:   cmdStringArr,
		},
		&container.HostConfig{
			Resources: container.Resources{
				Memory:     memLim * 1000 * 1000, // MB
				CPUPeriod:  int64(100 * 1000),
				CPUQuota:   int64(cpuPercent * 1000),
				CpusetCpus: cpuSetCpus,
			},
			LogConfig: container.LogConfig{
				Type:   "json-file",
				Config: map[string]string{"max-file": "2", "max-size": "20m"},
			},

			Binds: binds,
		},
		nil,
		nil,
		containerName,
	)
	if err != nil {
		panic(err)
	}

	log.Print("ContainerCreate resp: ", resp)

	err = cli.ContainerStart(
		ctx,
		resp.ID,
		types.ContainerStartOptions{},
	)
	if err != nil {
		panic(err)
	}

	return nil, resp.ID
}

func IsContainerRunning(containerNameOrID string) (e error, bRunning bool) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return errors.Wrap(err, "client.NewClientWithOpts: "), false
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return errors.Wrap(err, "cli.ContainerList()"), false
	}

	for _, container := range containers {
		if container.Names[0] == "/"+containerNameOrID || container.ID == containerNameOrID {
			log.Println("container running ")
			return nil, true
		}
	}

	return nil, false
}

func StopContainer(containerId string) (e error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return errors.Wrap(err, "client.NewEnvClient(): ")
	}

	log.Println("trying to invoke ContainerStop ")
	err = cli.ContainerStop(context.Background(), containerId, container.StopOptions{
		Signal:  "",
		Timeout: nil,
	})
	if err != nil {
		return errors.Wrap(err, "cli.ContainerStop(): ")
	}

	log.Println("Container Stopped")
	return nil
}

/*
func AAA() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// 容器 ID 或名称
	containerID := "your_container_id_or_name"

	// 目标导出文件路径
	exportFilePath := "exported_container.tar"

	// 打开目标文件以写入导出数据
	exportFile, err := os.Create(exportFilePath)
	if err != nil {
		panic(err)
	}
	defer exportFile.Close()

	// 调用 Docker API 导出容器
	resp, err := cli.ContainerExport(ctx, containerID, types.ContainerExportOptions{})
	if err != nil {
		panic(err)
	}
	defer resp.Close()

	// 将导出数据写入目标文件
	_, err = io.Copy(exportFile, resp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("容器")
}

*/
