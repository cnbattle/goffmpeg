package goffmpeg

import (
	"testing"
)

// TestGetDuration
func TestRemoveTrack(t *testing.T) {
	var testFFmpeg FFmpeg
	testFFmpeg.filePath = "./test/test.mp4"
	err := testFFmpeg.RemoveTrack("./test/test_nosound.mp4")
	if err != nil {
		t.Fatal(err)
	}
}
