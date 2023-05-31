package apiServer

import (
	"database/sql"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store/sqlStore"
	"github.com/gorilla/sessions"
	"net/http"
)

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

	return db, db.Ping()
}
