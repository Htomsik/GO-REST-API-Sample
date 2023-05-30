package store_test

import (
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store"
	"github.com/Htomsik/GO-REST-API-Sample/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Add(t *testing.T) {
	st, teardown := store.TestStore(t, databaseType, databaseURL)
	defer teardown("users")

	user, err := st.User().Add(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	st, teardown := store.TestStore(t, databaseType, databaseURL)
	defer teardown("users")

	email := "user@ex.com"

	_, err := st.User().FindByEmail(email)

	assert.Error(t, err)

	user := model.TestUser(t)
	user.Email = email

	// Проверка после создания
	st.User().Add(user)

	user, err = st.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
