package docker_vol

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"io/ioutil"
	"os"
)

// client.NewClientWithOpts( client.WithVersion("1.42"))

// client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

func CreateVolume(ctx context.Context, volumeName string) (id string, e error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)

	}

	err = cli.VolumeRemove(context.Background(), volumeName, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Volume %s deleted\n", volumeName)

	volume, err := cli.VolumeCreate(ctx, volume.CreateOptions{
		ClusterVolumeSpec: nil,
		Driver:            "",
		DriverOpts:        nil,
		Labels:            nil,
		Name:              volumeName,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created volume: %s\n", volume.Name)

	// 获取卷的挂载路径
	volumeMountpoint := volume.Mountpoint

	// 定义要写入卷中的文件内容
	fileContent := "Hello, Docker Volume!"

	// 将文件内容写入卷中
	filePath := volumeMountpoint + "/my_file.txt"
	err = ioutil.WriteFile(filePath, []byte(fileContent), os.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File %s created in volume\n", filePath)

	volumeInspectResp, err := cli.VolumeInspect(context.Background(), volume.Name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Volume Info: %+v\n", volumeInspectResp)

	return "", nil
}
