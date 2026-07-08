package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}
	api := application{
		config: cfg,
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
