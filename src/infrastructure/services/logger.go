package services

import (
	"github.com/aso779/go-ddd-example/application/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(conf *config.Config) *zap.Logger {
	logger, _ := zap.Config{
		Encoding:          conf.Logs.Format,
		Level:             zap.NewAtomicLevelAt(LogLevel(conf.Logs.Level)),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: true,
		Development:       false,
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:    "level",
			MessageKey:  "msg",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}.Build()

	defer logger.Sync()

	return logger
}

func LogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.WarnLevel
	}
}
