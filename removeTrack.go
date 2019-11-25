package goffmpeg

import (
	"bytes"
	"os/exec"
)

// RemoveTrack 删除原视频音轨
func (model *MediaFile) RemoveTrack(noSoundVideoName string) error {
	cmd := exec.Command(model.GetFFmpegCmd(), "-i", model.FilePath, "-c:v", "copy", "-an", noSoundVideoName)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
