package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open Открытие соединеня с бд
func (store *Store) Open() error {
	db, err := sql.Open(store.config.DatabaseType, store.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	store.db = db

	return nil
}

// Close Закрытие соединения с бд
func (store *Store) Close() error {
	if err := store.db.Close(); err != nil {
		return err
	}

	return nil
}
