package main

import (
	"testing"
)

func TestApp_WhenNoImagesRecorded_AddAllImagesNotUsed(t *testing.T) {
	mockDockerService := &MockDockerService{}
	mockDockerService.ImagesUsedByContainersFunc = func() []string {
		return []string { "image1" }
	}
	mockDockerService.AllImagesFunc = func() []string {
		return []string { "image1", "image2" }
	}
	app := App{dockerService: mockDockerService}

	app.updateImageRecord()

	assert(t, len(app.imageRecord) == 1, "imageRecord should have one key")

	_, exists := app.imageRecord["image2"]
	assert(t, exists, "imageRecord should have 'image2'")
}

//func TestApp_WhenNoImagesRecorded_AddAllImagesWithCorrectTimestamp(t *testing.T) {
//	//
//}
//
//func TestApp_WhenImagesRecorded_AddOnlyNewImagesThatAreNotUsed(t *testing.T) {
//	//
//}
//
//func TestApp_WhenImagesRecorded_NotUpdateImagesThatAreAlreadyKnow(t *testing.T) {
//	//
//}
//
//func TestApp_WhenDeleteImages_OnlyDeleteOlderThanConfiguredAge(t *testing.T) {
//	//
//}

func assert(t *testing.T, expression bool, explanation string) {
	if !expression {
		t.Fatal(explanation)
	}
}