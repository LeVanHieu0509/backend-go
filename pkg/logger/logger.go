package logger

import (
	"os"

	"github.com/LeVanHieu0509/backend-go/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.LogLevel

	// debug -> info -> warning -> error -> fatal -> panic
	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel

	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}

}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// convert 1719303947.491699 --> 2024-06-25T15:25:47.491+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// convert ts --> Time
	encodeConfig.TimeKey = "time"
	//from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//caller":"cli/main.log.go:23
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}
