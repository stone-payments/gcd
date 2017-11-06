package main

import "time"

type App struct {
	dockerService DockerServiceInterface
	imageRecord map[string]time.Time
}

func (this App) updateImageRecord() {
	imagesUsedByContainers := this.dockerService.ImagesUsedByContainers()
	allImages := this.dockerService.AllImages()

	var imagesNotUsedByContainers []string

	used := false
	for _, image := range allImages {
		for _, imageUsed := range imagesUsedByContainers {
			if image == imageUsed {
				used = true
				break
			}
		}
		if !used {
			imagesNotUsedByContainers = append(imagesNotUsedByContainers, image)
		}
	}

	for _, image := range imagesNotUsedByContainers {
		this.imageRecord[image] = time.Unix(0, 0)
	}

}