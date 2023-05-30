package store

import "github.com/Htomsik/GO-REST-API-Sample/internal/model"

// UserRepository ...
type UserRepository interface {
	Add(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
