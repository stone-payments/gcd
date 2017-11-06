package main

type DockerServiceInterface interface {
	ImagesUsedByContainers() []string
	AllImages() []string
}
