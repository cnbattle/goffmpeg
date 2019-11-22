package goffmpeg

import (
	"bytes"
	"errors"
	"os/exec"
)

// AddBackgroundMusic 添加背景音乐
func (ffmpeg *FFmpeg) AddBackgroundMusic(backgroundMusicFile, saveVideoFile string) error {
	// 获取视频时长
	videoLen, err := ffmpeg.GetDuration()
	if err != nil {
		return err
	}
	// 获取音频时长
	mp3FFmpeg := FFmpeg{
		filePath: backgroundMusicFile,
	}
	mp3Len, err := mp3FFmpeg.GetDuration()
	if err != nil {
		return err
	}
	if videoLen > mp3Len {
		return errors.New("视频长度大于音频")
	}
	cmd := exec.Command(ffmpeg.GetFFmpegCmd(), "-i", ffmpeg.filePath, "-i", backgroundMusicFile, "-shortest", saveVideoFile)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return errors.New(buf.String() + err.Error())
	}
	return nil
}
