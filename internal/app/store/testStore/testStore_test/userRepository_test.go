package testStore_test

import (
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/model"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store/testStore"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Add(t *testing.T) {
	st := testStore.New()

	user := model.TestUser(t)
	err := st.User().Add(user)

	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestUserRepository_FindByEmailNotAdded(t *testing.T) {
	// Act
	st := testStore.New()

	// Arrange
	_, err := st.User().FindByEmail("user@ex.com")

	// Assert
	assert.EqualError(t, err, model.RecordNotFound.Error())
}

func TestUserRepository_FindByEmailAdded(t *testing.T) {
	// Act
	st := testStore.New()
	email := "user@ex.com"
	user := model.TestUser(t)
	user.Email = email

	// Arrange
	st.User().Add(user)
	user, err := st.User().FindByEmail(email)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindNotAdded(t *testing.T) {
	// Act
	st := testStore.New()

	// Arrange
	_, err := st.User().Find(0)

	// Assert
	assert.EqualError(t, err, model.RecordNotFound.Error())
}

func TestUserRepository_FindAdded(t *testing.T) {
	// Act
	st := testStore.New()
	user := model.TestUser(t)

	// Arrange
	st.User().Add(user)
	user, err := st.User().Find(user.ID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
