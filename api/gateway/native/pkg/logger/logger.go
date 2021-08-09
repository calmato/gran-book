package logger

import (
	"fmt"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewGinMiddleware - gin-gonic/gin 用のミドルウェアを生成
func NewGinMiddleware(logPath string, logLevel string) (gin.HandlerFunc, error) {
	logger, err := New(logPath, logLevel)
	if err != nil {
		return nil, err
	}

	return ginzap.Ginzap(logger, time.RFC3339, true), nil
}

// New - Loggerの生成
func New(logPath string, logLevel string) (*zap.Logger, error) {
	level := getLogLevel(logLevel)

	encoderConfig := zapcore.EncoderConfig{
		CallerKey:      "caller",
		LevelKey:       "level",
		MessageKey:     "msg",
		NameKey:        "name",
		StacktraceKey:  "stacktrace",
		TimeKey:        "time",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		level,
	)

	// logPathが""のとき、標準出力のみ
	if logPath == "" {
		logger := zap.New(zapcore.NewTee(
			consoleCore,
		))

		return logger, nil
	}

	// logPathが""でないとき、ファイル出力も追加
	outputPath := fmt.Sprintf("%s/outputs.log", logPath)
	file, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}

	logCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(file),
		level,
	)

	logger := zap.New(zapcore.NewTee(
		consoleCore,
		logCore,
	))

	return logger, nil
}

func getLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
