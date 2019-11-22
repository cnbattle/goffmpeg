package goffmpeg

import (
	"strings"
	"testing"
)

// TestGetFormatName
func TestGetFormatName(t *testing.T) {
	var testFFmpeg FFmpeg
	testFFmpeg.filePath = "./test/test.mp4"
	formatName, err := testFFmpeg.GetFormatName()
	if err != nil {
		t.Fatal(err)
	}
	if strings.EqualFold(formatName, "mov,mp4,m4a,3gp,3g2,mj2") {
		t.Fatal("TestGetFormatName error,is:" + formatName)
	}
}
