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
	var testFFmpeg FFmpeg
	testFFmpeg.filePath = "./test/demo.mp4"
	err = testFFmpeg.RemoveTrack(testFileName)
	if err != nil {
		t.Fatal(err)
	}
}
