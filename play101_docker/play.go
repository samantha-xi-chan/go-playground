package play101_docker

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-playground/util/docker_container"
	"go-playground/util/docker_image"
	"go-playground/util/docker_vol"
	"log"
	"time"
)

func Play2() {
	IMAGE_NAME := "joedval/stress"
	CMD_STRING_ARR := []string{"--cpu", "200"}
	VOL_NAME := "volName"

	//IMAGE_NAME := "busybox"
	//CMD_STRING_ARR := []string{"/usr/bin/curl", "http://192.168.32.13:8080/download/date.sh", "|", "sh"}

	docker_image.PullImageBlo(context.Background(), IMAGE_NAME)

	docker_vol.CreateVolume(context.Background(), VOL_NAME)

	err, containerId := docker_container.StartContainerNBlo(context.Background(),
		IMAGE_NAME, CMD_STRING_ARR,
		int64(1*1024*1024*1024),
		500,
		"0-3",
		fmt.Sprintf("cn_%d", time.Now().UnixMilli()),
		VOL_NAME,
		"/path/in/container",
	)
	if err != nil {
		log.Println("docker_image.PullImageB, err = ", err)
		return
	}

	log.Println("started containerId: ", containerId)

	err, isRunning := docker_container.IsContainerRunning(containerId)
	if err != nil {
		logrus.Error(err)
		return
	}
	if isRunning {
		logrus.Debug("Running")

		time.Sleep(time.Second * 8888)
		docker_container.StopContainer(containerId)
		return
	}

}
