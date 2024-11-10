package logx

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"path"
)

func NewWriter(dir string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   path.Join(dir, "app.log"), // 日志输出文件
		MaxSize:    2,                         // 日志最大保存2M
		MaxBackups: 10,
		LocalTime:  true,
	}
}
