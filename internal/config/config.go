package config

import (
	"github.com/bretanac93/finanzen/internal/db"
	"github.com/bretanac93/finanzen/internal/logger"
	"github.com/bretanac93/finanzen/internal/server"
)

type Config struct {
	DB     *db.Config     `env:", prefix=DB_"`
	Logger *logger.Config `env:", prefix=LOGGER_"`
	Server *server.Config `env:", prefix=SERVER_"`
}
