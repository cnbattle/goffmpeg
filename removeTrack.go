package goffmpeg

import (
	"bytes"
	"os/exec"
)

// RemoveTrack 删除原视频音轨
func (ffmpeg *FFmpeg) RemoveTrack(noSoundVideoName string) error {
	cmd := exec.Command(ffmpeg.GetFFmpegCmd(), "-i", ffmpeg.filePath, "-c:v", "copy", "-an", noSoundVideoName)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
