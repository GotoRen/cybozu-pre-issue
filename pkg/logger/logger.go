// Package logger provides logger methods.
//
//nolint:gomnd
package logger

import (
	"fmt"
	"os"
	"path/filepath"

	local "github.com/GotoRen/cybozu-pre-issue/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitZap provides logging with zap.
func InitZap(cfg *local.Config) error {
	logLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)

	// Standard output
	stdCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(config()),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)

	f, err := setFile()
	if err != nil {
		return err
	}

	// Output as a log file
	logCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(config()),
		zapcore.AddSync(f),
		logLevel,
	)

	logger := zap.New(zapcore.NewTee(
		logCore,
	))

	if cfg.DebugMode {
		logger = zap.New(zapcore.NewTee(
			stdCore,
			logCore,
		))
	}

	zap.ReplaceGlobals(logger)

	return nil
}

// setFile return the location where the log file will be placed.
func setFile() (*os.File, error) {
	dirPath := "logs/"
	fileName := "log.json"
	content := filepath.Join(dirPath, fileName)

	if _, err := os.Stat(content); err != nil {
		if os.IsNotExist(err) {
			if _, err := os.Create(content); err != nil {
				return nil, fmt.Errorf("failed to create the logging file: %w", err)
			}
		}
	}

	f, err := os.OpenFile(content, os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return nil, fmt.Errorf("failed to open the logging file: %w", err)
	}

	return f, nil
}

// config returns EncoderConfig for production environments.
func config() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()

	cfg.MessageKey = "msg"
	cfg.LevelKey = "level"
	cfg.NameKey = "name"
	cfg.TimeKey = "timestamp"
	cfg.CallerKey = "caller"
	cfg.FunctionKey = "func"
	cfg.StacktraceKey = "stacktrace"
	cfg.LineEnding = "\n"
	cfg.EncodeTime = zapcore.EpochTimeEncoder
	cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	cfg.EncodeDuration = zapcore.SecondsDurationEncoder
	cfg.EncodeCaller = zapcore.ShortCallerEncoder

	return cfg
}

// LogDebug is Key-value format debug log.
func LogDebug(msg string, kv ...interface{}) {
	zap.S().Debugw(msg, kv...)
}

// LogErr is Key-value format error log.
func LogErr(msg string, kv ...interface{}) {
	zap.S().Errorw(msg, kv...)
}
