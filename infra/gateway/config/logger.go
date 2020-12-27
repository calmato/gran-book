package config

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger(logPath string, logLevel string) (*zap.Logger, error) {
	outputPaths := []string{}
	if logPath != "" {
		outputPath := fmt.Sprintf("%s/outputs.log", logPath)
		outputPaths = append(outputPaths, outputPath)
	}

	errorOutputPaths := []string{}
	if logPath != "" {
		errorOutputPath := fmt.Sprintf("%s/errors.log", logPath)
		errorOutputPaths = append(errorOutputPaths, errorOutputPath)
	}

	level := getLogLevel(logLevel)

	config := zap.Config{
		Level:         zap.NewAtomicLevelAt(level),
		DisableCaller: true,
		Encoding:      "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
		OutputPaths:      outputPaths,
		ErrorOutputPaths: errorOutputPaths,
	}

	return config.Build()
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
