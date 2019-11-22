package goffmpeg

import (
	"bytes"
	"os/exec"
)

// RemoveTrack 删除原视频音轨
func (ffmpeg *FFmpeg) RemoveTrack(noSoundVideoName string) error {
	//去掉原视频音轨
	//E:\anzhuangbao\ffmpeg\bin\ffmpeg -i G:\hi.mp4 -c:v copy -an G:\nosound.mp4
	//添加背景音乐
	//E:\anzhuangbao\ffmpeg\bin\ffmpeg -i G:\nosound.mp4 -i G:\songs.mp3 -t 7.1 -c y copy G:\output.mp4

	cmd := exec.Command(ffmpeg.GetFFmpegCmd(), "-i", ffmpeg.filePath, "-c:v", "copy", "-an", noSoundVideoName)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
