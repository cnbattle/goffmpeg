package goffmpeg

// MediaFile 住结构体
type MediaFile struct {
	ffmpegCmd  string
	ffprobeCmd string
	ffplayCmd  string
	// 处理文件的路径
	FilePath string
}

// SetFFmpegCmd 设置FFmpeg命令
func (model *MediaFile) SetFFmpegCmd(cmd string) *MediaFile {
	model.ffmpegCmd = cmd
	return model
}

// GetFFmpegCmd 获取FFmpeg命令
func (model *MediaFile) GetFFmpegCmd() string {
	if model.ffmpegCmd == "" {
		return "ffmpeg"
	}
	return model.ffmpegCmd
}

// SetFFprobeCmd 设置FFprobe命令
func (model *MediaFile) SetFFprobeCmd(cmd string) *MediaFile {
	model.ffprobeCmd = cmd
	return model
}

// GetFFprobeCmd 获取FFprobe命令
func (model *MediaFile) GetFFprobeCmd() string {
	if model.ffprobeCmd == "" {
		return "ffprobe"
	}
	return model.ffmpegCmd
}

// SetFFplayCmd 设置FFplay命令
func (model *MediaFile) SetFFplayCmd(cmd string) *MediaFile {
	model.ffplayCmd = cmd
	return model
}

// GetFFplayCmd 获取FFplay命令
func (model *MediaFile) GetFFplayCmd() string {
	if model.ffplayCmd == "" {
		return "ffplay"
	}
	return model.ffplayCmd
}
