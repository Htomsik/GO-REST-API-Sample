package sqlStore

import (
	"database/sql"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app"
	"github.com/Htomsik/GO-REST-API-Sample/internal/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Add ...
func (repository *UserRepository) Add(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeAdd(); err != nil {
		return err
	}

	return repository.store.db.QueryRow(
		"INSERT INTO users(email,encryptedPassword) values ($1,$2) RETURNING id",
		user.Email,
		user.EncryptedPassword,
	).Scan(&user.ID)
}

// FindByEmail ...
func (repository *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	err := repository.store.db.QueryRow(
		"SELECT id,email,encryptedPassword FROM users WHERE email = $1",
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.EncryptedPassword,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			err = app.RecordNotFound
		}
	}

	return user, err
}
