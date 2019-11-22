package goffmpeg

import (
	"testing"
)

// TestGetDuration
func TestGetDuration(t *testing.T) {
	var testFFmpeg FFmpeg
	testFFmpeg.filePath = "./test/demo.mp4"
	duration, err := testFFmpeg.GetDuration()
	if err != nil {
		t.Fatal(err)
	}
	if duration != 8.545 {
		t.Fatal("TestGetDuration error: is not 8.545")
	}
}
