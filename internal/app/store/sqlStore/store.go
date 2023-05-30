package sqlStore

import (
	"database/sql"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store"
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (store *Store) User() store.UserRepository {
	if store.userRepository != nil {
		return store.userRepository
	}

	store.userRepository = &UserRepository{
		store: store,
	}

	return store.userRepository
}
