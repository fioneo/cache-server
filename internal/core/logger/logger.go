package core_logger

import (
	"log/slog"
	"os"
)

func Init() {
	level := slog.LevelInfo

	if os.Getenv("LEVEL") == "debug" {
		level = slog.LevelDebug
	}
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}),
	)

	slog.SetDefault(log)
}
