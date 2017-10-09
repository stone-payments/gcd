package main

import (
	"context"
	"testing"

	docker "github.com/fsouza/go-dockerclient"
)

const dockerHost = "unix:///var/run/docker.sock"

func TestRemoveHealthyContainersExited(t *testing.T) {
	dc, err := docker.NewClient(dockerHost)
	if err != nil {
		t.Error(err)
	}

	defer dc.RemoveImage("hello-world") // nolint: errcheck
	err = dc.PullImage(docker.PullImageOptions{
		Repository: "hello-world",
		Tag:        "latest",
	}, docker.AuthConfiguration{})
	if err != nil {
		t.Error(err)
	}

	// Run a sample docker
	container, err := dc.CreateContainer(docker.CreateContainerOptions{
		Name:    "hello-test",
		Config:  &docker.Config{Image: "hello-world"},
		Context: context.Background(),
	})
	if err != nil {
		t.Error(err)
	}

	run(
		dc,
		options{
			dockerHost:                    dockerHost,
			removeImages:                  false,
			removeHealthyContainersExited: true,
		},
	)

	// Check if the sample docker removed
	assertContainerNotExists(t, dc, container)
}

func TestRemoveImages(t *testing.T) {
	dc, err := docker.NewClient(dockerHost)
	if err != nil {
		t.Error(err)
	}

	err = dc.PullImage(docker.PullImageOptions{
		Repository: "hello-world",
		Tag:        "latest",
	}, docker.AuthConfiguration{})
	if err != nil {
		t.Error(err)
	}

	run(
		dc,
		options{
			dockerHost:                    "unix:///var/run/docker.sock",
			removeImages:                  true,
			removeHealthyContainersExited: false,
		},
	)

	// Check if the sample docker removed
	assertImageNotExists(t, dc, "hello-world")
}

func assertContainerNotExists(t *testing.T, dc *docker.Client, container *docker.Container) {
	containers, err := dc.ListContainers(docker.ListContainersOptions{
		All: true,
	})
	if err != nil {
		t.Error(err)
	}
	for _, c := range containers {
		if c.ID == container.ID {
			t.Errorf("Container %s exists which shouldn't.", container.Name)
			return
		}
	}
}

func assertImageNotExists(t *testing.T, dc *docker.Client, tag string) {
	images, err := dc.ListImages(docker.ListImagesOptions{
		All: true,
	})
	if err != nil {
		t.Error(err)
	}
	for _, i := range images {
		for _, tag := range i.RepoTags {
			if tag == "hello-world:latest" {
				t.Errorf("Image %s exists which shouldn't.", tag)
			}
		}
	}
}
