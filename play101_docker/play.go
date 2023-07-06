package play101_docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func Play() {
	// 创建Docker客户端
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// 配置容器的资源限制
	resources := &container.Resources{
		NanoCPUs: int64(500000000),              // 0.5核心
		Memory:   int64(2 * 1024 * 1024 * 1024), // 2GB
	}

	// 创建容器的配置
	config := &container.Config{
		Image: "joedval/stress", // 替换为你的Docker镜像名称
		Cmd:   []string{"--cpu", "2"},
	}

	// 创建容器的主机配置
	hostConfig := &container.HostConfig{
		Resources: *resources,
	}

	// 创建容器
	resp, err := cli.ContainerCreate(
		context.Background(),
		config,
		hostConfig,
		nil,
		nil,
		fmt.Sprintf("containerName_%d", time.Now().UnixMilli()),
	)
	if err != nil {
		panic(err)
	}

	// 启动容器
	err = cli.ContainerStart(
		context.Background(),
		resp.ID,
		types.ContainerStartOptions{},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("容器已启动:", resp.ID)

	// 等待容器退出
	statusCh, errCh := cli.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	// 获取容器日志
	out, err := cli.ContainerLogs(context.Background(), resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		panic(err)
	}

	// 打印容器日志
	_, err = io.Copy(os.Stdout, out)
	if err != nil {
		panic(err)
	}
}
