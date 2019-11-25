package goffmpeg

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

// GetFormatName 获取文件格式
func (model *MediaFile) GetFormatName() (string, error) {
	cmd := exec.Command(model.GetFFprobeCmd(), "-v", "error", "-show_entries", "format=format_name",
		"-of", "default=noprint_wrappers=1:nokey=1", model.FilePath)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return "", errors.New(fmt.Sprintf("%s-%s-%s", "could not generate frame", err.Error(), buf.String()))
	}
	return buf.String(), nil
}
