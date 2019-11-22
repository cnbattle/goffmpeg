package goffmpeg

import (
	"os"
	"testing"
)

// TestGetDuration
func TestAddBackgroundMusic(t *testing.T) {
	testFileName := "./test/test_AddBackgroundMusic.mp4"
	_, err := os.Stat(testFileName)
	if err == nil {
		_ = os.Remove(testFileName)
	}

	var testFFmpeg FFmpeg
	testFFmpeg.SetFilePath("./test/nosound.mp4")
	err = testFFmpeg.AddBackgroundMusic("./test/demo.mp3", testFileName)
	if err != nil {
		t.Fatal(err)
	}
}
