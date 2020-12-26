package config

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger(logPath string) (*zap.Logger, error) {
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

	config := zap.Config{
		Level:    zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			NameKey:      "name",
			FunctionKey:  "func",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      outputPaths,
		ErrorOutputPaths: errorOutputPaths,
	}

	return config.Build()
}
