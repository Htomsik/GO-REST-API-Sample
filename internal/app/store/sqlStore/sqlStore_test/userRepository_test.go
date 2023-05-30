package sqlStore_test

import (
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store/sqlStore"
	"github.com/Htomsik/GO-REST-API-Sample/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Add(t *testing.T) {
	// Arrange
	db, teardown := sqlStore.TestDb(t, databaseType, databaseURL)
	defer teardown("users")

	// Act
	st := sqlStore.New(db)
	user := model.TestUser(t)
	err := st.User().Add(user)

	// Assert
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestUserRepository_FindByEmailRandom(t *testing.T) {
	// Arrange
	db, teardown := sqlStore.TestDb(t, databaseType, databaseURL)
	defer teardown("users")
	st := sqlStore.New(db)

	// Act
	_, err := st.User().FindByEmail("user@ex.com")

	// Assert
	assert.EqualError(t, err, store.RecordNotFound.Error())
}

func TestUserRepository_FindByEmailAdded(t *testing.T) {
	// Arrange
	db, teardown := sqlStore.TestDb(t, databaseType, databaseURL)
	defer teardown("users")
	st := sqlStore.New(db)
	user := model.TestUser(t)
	user.Email = "user@ex.com"

	// Act
	st.User().Add(user)
	user, err := st.User().FindByEmail(user.Email)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
