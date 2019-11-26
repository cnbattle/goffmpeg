package goffmpeg

import (
	"testing"
)

// TestGetDuration
func TestGetDuration(t *testing.T) {
	testMediaFile := MediaFile{FilePath: "./testdata/demo.mp4"}
	duration, err := testMediaFile.GetDuration()
	if err != nil {
		t.Fatal(err)
	}
	if duration != 8.545 {
		t.Fatal("TestGetDuration error: is not 8.545")
	}
}
