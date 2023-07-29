package play101_docker

import (
	"context"
	"fmt"
	"go-playground/util/docker_container"
	"go-playground/util/docker_image"
	"log"
	"time"
)

func Play() {
	IMAGE_NAME := "joedval/stress"
	CMD_STRING_ARR := []string{"--cpu", "100"}

	err := docker_image.PullImageBlo(context.Background(), IMAGE_NAME)
	if err != nil {
		//logrus.Error("docker_image.PullImageB, err = ", err)
		return
	}

	err = docker_image.RemoveImage(context.Background(), IMAGE_NAME)
	if err != nil {
		//logrus.Error("docker_image.PullImageB, err = ", err)
		return
	}
	return

	err = docker_container.StartContainerBlo(context.Background(),
		IMAGE_NAME, CMD_STRING_ARR,
		int64(1*1024*1024*1024),
		fmt.Sprintf("cn_%d", time.Now().UnixMilli()))
	if err != nil {
		log.Println("docker_image.PullImageB, err = ", err)
		return
	}
}
