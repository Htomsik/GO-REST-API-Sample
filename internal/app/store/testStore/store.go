package testStore

import (
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/model"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

// User ...
func (store *Store) User() store.UserRepository {
	if store.userRepository != nil {
		return store.userRepository
	}

	store.userRepository = &UserRepository{
		store: store,
		users: make(map[int]*model.User),
	}

	return store.userRepository
}
