package sqlite

import (
	"database/sql"
	"fmt"
	// "os"
	// "path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	// if storagePath == "" {
	// 	return nil, fmt.Errorf("%s: storage path is empty", op)
	// }

	// if dir := filepath.Dir(storagePath); dir != "." {
	// 	if err := os.MkdirAll(dir, 0o755); err != nil {
	// 		return nil, fmt.Errorf("%s: create storage dir: %w", op, err)
	// 	}
	// }

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
	id INTEGER PRIMARY KEY,
	alias TEXT NOT NULL UNIQUE,
	url TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if _, err = stmt.Exec(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
