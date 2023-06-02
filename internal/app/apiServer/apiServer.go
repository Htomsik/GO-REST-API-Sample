package apiServer

import (
	"database/sql"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store/sqlStore"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

const migrationsPath = "file://migrations"

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseType, config.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlStore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionsKey))

	srv := newServer(store, sessionStore)

	srv.logger.Infof("Server started on port %v", config.Port)

	return http.ListenAndServe(config.Port, srv)
}

// newDB ...
func newDB(databaseType string, databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open(databaseType, databaseUrl)

	if err != nil {
		return nil, err
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		return nil, err
	}

	migration, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		databaseType,
		driver,
	)

	if err != nil {
		return nil, err
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return db, db.Ping()
}
