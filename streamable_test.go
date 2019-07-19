package streamable

import (
	"os"
	"testing"
)

var (
	email = os.Getenv("STREAMABLE_EMAIL")
	pass  = os.Getenv("STREAMABLE_PASSWORD")
)

func TestGetVideoFromShortCode(t *testing.T) {
	streamable := New(email, pass)

	video, err := streamable.GetVideoFromShortcode("ts9vt")
	if err != nil {
		t.Error(err)
	}

	actualTitle := video.Title
	expectedTitle := "Test Import"
	if actualTitle != expectedTitle {
		t.Errorf("actual %q, expected %q", actualTitle, expectedTitle)
	}

	actualURL := video.URL
	expectedURL := "streamable.com/ts9vt"
	if actualURL != expectedURL {
		t.Errorf("actual %q, expected %q", actualURL, expectedURL)
	}

	actualHeight := video.Files.Mp4.Height
	expectedHeight := 720
	if actualHeight != expectedHeight {
		t.Errorf("actual %q, expected %q", actualHeight, expectedHeight)
	}

	actualMp4MobileHeight := video.Files.Mp4Mobile.Height
	expectedMp4MobileHeight := 360
	if actualMp4MobileHeight != expectedMp4MobileHeight {
		t.Errorf("actual %q, expected %q", actualMp4MobileHeight, expectedMp4MobileHeight)
	}
}

func TestGetVideoFromURL(t *testing.T) {
	streamable := New(email, pass)

	videoEmbed, err := streamable.GetVideoFromURL("https://streamable.com/ts9vt")
	if err != nil {
		t.Error(err)
	}

	actualTitle := videoEmbed.Title
	expectedTitle := "Test Import"
	if actualTitle != expectedTitle {
		t.Errorf("actual %q, expected %q", actualTitle, expectedTitle)
	}

}

func TestImport(t *testing.T) {
	streamable := New(email, pass)

	result, err := streamable.Import("http://www.sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4")
	if err != nil {
		t.Error(err)
	}

	actualStatus := result.Status
	expectedStatus := 1
	if actualStatus != expectedStatus {
		t.Errorf("actual %q, expected %q", actualStatus, expectedStatus)
	}
}

func TestImportError(t *testing.T) {
	streamable := New(email, pass)

	result, err := streamable.Import("bad url")
	if err != nil {
		t.Error(err)
	}

	actualStatus := result.Status
	expectedStatus := 0

	if actualStatus != expectedStatus {
		t.Errorf("actual %q, expected %q", actualStatus, expectedStatus)
	}
}

func TestUpload(t *testing.T) {
	streamable := New(email, pass)

	iResult, err := streamable.Upload("SampleVideo_1280x720_1mb.mp4")
	if err != nil {
		t.Error(err)
	}

	actualStatus := iResult.Status
	expectedStatus := 1
	if actualStatus != expectedStatus {
		t.Errorf("actual %q, expected %q", actualStatus, expectedStatus)
	}
}
