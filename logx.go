package logx

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func MustNew(c Conf) *zap.Logger {
	var level zapcore.Level
	var writer zapcore.WriteSyncer

	if c.Output {
		writer = zapcore.AddSync(os.Stdout)
		level = zapcore.DebugLevel
	} else {
		err := os.MkdirAll(c.Dir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		var w = NewWriter(c.Dir)
		writer = zapcore.AddSync(w)
		level = zapcore.InfoLevel
	}

	var e = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
	})

	return zap.New(zapcore.NewCore(e, writer, level))
}
