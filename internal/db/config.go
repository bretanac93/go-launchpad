package db

// Config is a struct that holds the configuration for the database. Can be refactored later to support postgres.
type Config struct {
	Path string `env:"PATH" envDefault:"finanzen.db"`
}
