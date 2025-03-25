package logger

import (
	"log/slog"
	"os"
)

func Init(cfg Config) {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     parseLevel(cfg.Level),
	}

    var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
    if cfg.Output == "json" {
        handler = slog.NewJSONHandler(os.Stdout, opts)
    } 

    l := slog.New(handler)
    slog.SetDefault(l)
}

// parseLevel parses a string into a slog.Level. Falls back to [slog.LevelDebug] if the string is invalid.
func parseLevel(level string) slog.Level {
	var l slog.Level
	if err := l.UnmarshalText([]byte(level)); err != nil {
		return slog.LevelDebug
	}

	return l
}
