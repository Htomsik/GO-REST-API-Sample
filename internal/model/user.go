package model

import "golang.org/x/crypto/bcrypt"

// User ...
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

// BeforeAdd ...
func (user *User) BeforeAdd() error {
	if len(user.Password) > 0 {
		encryptedPassword, err := encryptString(user.Password)

		if err != nil {
			return err
		}

		user.EncryptedPassword = encryptedPassword
	}

	return nil
}

// encryptString
func encryptString(text string) (string, error) {
	encryptedText, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(encryptedText), err
}
