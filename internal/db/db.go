package db

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Open(cfg Config) (*sql.DB, func(), error) {
	if err := createDBFile(cfg.Path); err != nil {
		return nil, nil, fmt.Errorf("error while creating db: %w", err)
	}

	client, err := sql.Open("sqlite3", cfg.Path)
	if err != nil {
		return nil, nil, fmt.Errorf("error while opening db connection: %w", err)
	}
	closeFunc := func() {
		_ = client.Close()
	}

	if err = runMigrations(client); err != nil {
		return nil, nil, fmt.Errorf("error running migrations: %w", err)
	}

	return client, closeFunc, nil
}

func createDBFile(databasePath string) error {
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		// Create the sqlite database file if it does not exist
		_, err := os.Create(databasePath)
		if err != nil {
			return fmt.Errorf("error while creating the sqlite database file: %w", err)
		}
	}

	return nil
}

func runMigrations(conn *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	if err := goose.Up(conn, "migrations"); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	return nil
}
