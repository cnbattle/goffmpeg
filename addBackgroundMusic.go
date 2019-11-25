package goffmpeg

import (
	"bytes"
	"errors"
	"os/exec"
)

// AddBackgroundMusic 添加背景音乐
func (model *MediaFile) AddBackgroundMusic(backgroundMusicFile, saveVideoFile string) error {
	// 获取视频时长
	videoLen, err := model.GetDuration()
	if err != nil {
		return err
	}
	// 获取音频时长
	testMp3MediaFile := MediaFile{FilePath: backgroundMusicFile}
	mp3Len, err := testMp3MediaFile.GetDuration()
	if err != nil {
		return err
	}
	if videoLen > mp3Len {
		return errors.New("视频长度大于音频")
	}
	cmd := exec.Command(model.GetFFmpegCmd(), "-i", model.FilePath, "-i", backgroundMusicFile, "-shortest", saveVideoFile)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return errors.New(buf.String() + err.Error())
	}
	return nil
}
