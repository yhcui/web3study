package initialize

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() (log *slog.Logger) {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	slog.SetDefault(Logger)
	return Logger
}
