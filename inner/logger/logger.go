package logger

import (
	"log"
	"log/slog"
	"os"
	"sync"
)

type LogLevel string

const (
	LogInfo  LogLevel = "info"
	LogDebug LogLevel = "debug"
	LogError LogLevel = "error"

	DEFAULTLOGLEVEL = LogInfo
)

var (
	once sync.Once
)

// MustInitLogger setting up default slog
func MustInitLogger(level LogLevel) {
	once.Do(func() {
		var clog *slog.Logger
		switch level {
		case LogDebug:
			clog = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case LogInfo:
			clog = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		case LogError:
			clog = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
		default:
			log.Fatalf("Unknown log level: %s", level)
		}
		slog.SetDefault(clog)
		slog.Info("Logger Initialized", "level", level)
	})
}

// Err Using to add error to the slog as attribute
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
