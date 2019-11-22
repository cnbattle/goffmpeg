package goffmpeg

import "errors"

type FFmpeg struct {
	ffmpegCmd  string
	ffprobeCmd string
	ffplayCmd  string
	FilePath   string
}

func (ffmpeg *FFmpeg) SetFilePath(filePath string) {
	ffmpeg.FilePath = filePath
}

func (ffmpeg *FFmpeg) GetFilePath() (string, error) {
	if ffmpeg.FilePath == "" {
		return "", errors.New("file path is empty")
	}
	return ffmpeg.FilePath, nil
}

func (ffmpeg *FFmpeg) SetFFmpegCmd(cmd string) {
	ffmpeg.ffmpegCmd = cmd
}

func (ffmpeg *FFmpeg) GetFFmpegCmd() string {
	if ffmpeg.ffmpegCmd == "" {
		return "ffmpeg"
	}
	return ffmpeg.ffmpegCmd
}

func (ffmpeg *FFmpeg) SetFFprobeCmd(cmd string) {
	ffmpeg.ffprobeCmd = cmd
}

func (ffmpeg *FFmpeg) GetFFprobeCmd() string {
	if ffmpeg.ffprobeCmd == "" {
		return "ffprobe"
	}
	return ffmpeg.ffmpegCmd
}

func (ffmpeg *FFmpeg) SetFFplayCmd(cmd string) {
	ffmpeg.ffplayCmd = cmd
}

func (ffmpeg *FFmpeg) GetFFplayCmd() string {
	if ffmpeg.ffplayCmd == "" {
		return "ffplay"
	}
	return ffmpeg.ffplayCmd
}
