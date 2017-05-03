package handler

import (
	"log"

	"github.com/fsouza/go-dockerclient"
)

func showContainers() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		panic(err)
	}
	for _, img := range imgs {
		log.Println("ID: ", img.ID)
		log.Println("RepoTags: ", img.RepoTags)
		log.Println("Created: ", img.Created)
		log.Println("Size: ", img.Size)
		log.Println("VirtualSize: ", img.VirtualSize)
		log.Println("ParentId: ", img.ParentID)
	}

}
