package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found")
	}

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: os.Getenv("GOOSE_DBSTRING"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	api := &application{
		config: cfg,
		logger: logger,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
}
