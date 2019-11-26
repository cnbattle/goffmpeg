package goffmpeg

import (
	"strings"
	"testing"
)

// TestGetFormatName
func TestGetFormatName(t *testing.T) {
	testMediaFile := MediaFile{FilePath: "./testdata/demo.mp4"}
	formatName, err := testMediaFile.GetFormatName()
	if err != nil {
		t.Fatal(err)
	}
	if strings.EqualFold(formatName, "mov,mp4,m4a,3gp,3g2,mj2") {
		t.Fatal("TestGetFormatName error,is:" + formatName)
	}
}
