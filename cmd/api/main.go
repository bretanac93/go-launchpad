package main

import (
	"context"
	"log/slog"

	"github.com/bretanac93/finanzen/internal/config"
	"github.com/bretanac93/finanzen/internal/db"
	"github.com/bretanac93/finanzen/internal/logger"
	"github.com/bretanac93/finanzen/internal/server"
	"github.com/bretanac93/finanzen/internal/users"

	"github.com/sethvargo/go-envconfig"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()

	var cfg config.Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		panic(err)
	}

	logger.Init(*cfg.Logger)

	dbConn, closeConn, err := db.Open(*cfg.DB)
	if err != nil {
		panic(err)
	}
	defer closeConn()

	slog.InfoContext(ctx, "Connected to database")

	_ = dbConn

	s := server.New(*cfg.Server)

	usersHandler := &users.UsersHandler{}

	s.AddHandlers(usersHandler)

	s.Run()
}
