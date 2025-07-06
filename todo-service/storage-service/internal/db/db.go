package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// NewConnection opens and returns a *sql.DB connection.
func NewConnection(dbPath string, logger *slog.Logger) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database at %s: %w", dbPath, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database at %s: %w", dbPath, err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	logger.Info("Database connection established", "path", dbPath)
	return db, nil
}

// In a real project, this would be handled by a proper migration tool.
func ApplySchema(db *sql.DB, logger *slog.Logger) error {
	tasksTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := db.ExecContext(ctx, tasksTableSQL); err != nil {
		return fmt.Errorf("failed to create tasks table: %w", err)
	}
	logger.Debug("Tasks table ensured")

	return nil
}
