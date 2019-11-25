package goffmpeg

import (
	"os"
	"testing"
)

// TestRemoveTrack
func TestRemoveTrack(t *testing.T) {
	testFileName := "./test/test_nosound.mp4"
	_, err := os.Stat(testFileName)
	if err == nil {
		_ = os.Remove(testFileName)
	}
	testMediaFile := MediaFile{FilePath: "./test/demo.mp4"}
	err = testMediaFile.RemoveTrack(testFileName)
	if err != nil {
		t.Fatal(err)
	}
}
