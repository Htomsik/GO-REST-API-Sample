package testStore

import (
	"github.com/Htomsik/GO-REST-API-Sample/internal/app"
	"github.com/Htomsik/GO-REST-API-Sample/internal/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store

	users map[string]*model.User
}

// Add ...
func (repository *UserRepository) Add(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeAdd(); err != nil {
		return err
	}

	repository.users[user.Email] = user

	user.ID = len(repository.users)

	return nil
}

// FindByEmail ...
func (repository *UserRepository) FindByEmail(email string) (*model.User, error) {
	user, exist := repository.users[email]

	if !exist {
		return nil, app.RecordNotFound
	}

	return user, nil
}
