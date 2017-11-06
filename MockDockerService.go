package main

type MockDockerService struct {
	ImagesUsedByContainersFunc func() []string
	AllImagesFunc func() []string
}

func (this *MockDockerService) ImagesUsedByContainers() []string {
	if this.ImagesUsedByContainersFunc != nil {
		return this.ImagesUsedByContainersFunc()
	}

	return []string{}
}

func (this *MockDockerService) AllImages() []string {
	if this.AllImagesFunc != nil {
		return this.AllImagesFunc()
	}

	return []string{}
}
