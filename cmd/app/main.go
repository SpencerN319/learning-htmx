package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/SpencerN319/learning-htmx/env"
	"github.com/SpencerN319/learning-htmx/server"
)

func init() {
	parseLogLevel := func(l string) slog.Level {
		switch l {
		default:
			return slog.LevelInfo
		}
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: parseLogLevel(env.Getenv("LOG_LEVEL", "INFO")),
	})))
}

func main() {
	svr :=  server.New(context.Background())
	defer svr.Close()
	svr.ListenAndServe()
}
