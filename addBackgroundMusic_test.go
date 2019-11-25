package goffmpeg

import (
	"os"
	"testing"
)

// TestAddBackgroundMusic
func TestAddBackgroundMusic(t *testing.T) {
	testFileName := "./test/test_AddBackgroundMusic.mp4"
	_, err := os.Stat(testFileName)
	if err == nil {
		_ = os.Remove(testFileName)
	}
	testMediaFile := MediaFile{FilePath: "./test/nosound.mp4"}
	err = testMediaFile.AddBackgroundMusic("./test/demo.mp3", testFileName)
	if err != nil {
		t.Fatal(err)
	}
}
