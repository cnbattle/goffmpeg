package goffmpeg

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

// GetDuration 获取视频时长
func (model *MediaFile) GetDuration() (float64, error) {
	cmd := exec.Command(model.GetFFprobeCmd(), "-v", "error", "-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1", model.FilePath)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return 0, errors.New(fmt.Sprintf("%s: %s", "could not generate frame", err.Error()))
	}
	return strconv.ParseFloat(buf.String()[0:len(buf.String())-2], 64)
}
