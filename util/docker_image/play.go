package docker_image

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

func PullImageBlo(ctx context.Context, imageName string) (err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return errors.Wrap(err, "client.NewClientWithOpts: ")
	}

	if imageExists(cli, imageName) {
		//log.Println("imageExists ok")
		return nil
	}

	log.Println("ImagePull starting")
	resp, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return errors.Wrap(err, "cli.ImagePull: ")
	}
	defer resp.Close()

	_, err = io.Copy(os.Stdout, resp)
	if err != nil {
		return errors.Wrap(err, "io.Copy: ")
	}

	return nil
}

func imageExists(cli *client.Client, imageName string) bool {
	_, _, err := cli.ImageInspectWithRaw(context.Background(), imageName)
	return err == nil
}

func RemoveImage(ctx context.Context, imageName string) (err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {

	}

	imageInspect, _, err := cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		if client.IsErrNotFound(err) {
			fmt.Printf("Image '%s' not found.\n", imageName)
		} else {
			fmt.Printf("Failed to inspect image: %s\n", err)
		}
		os.Exit(1)
	}

	// Delete the image
	_, err = cli.ImageRemove(context.Background(), imageInspect.ID, types.ImageRemoveOptions{})
	if err != nil {
		fmt.Printf("Failed to remove image: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Image '%s' deleted successfully.\n", imageName)

	return nil
}
