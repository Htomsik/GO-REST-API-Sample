package testStore_test

import (
	"github.com/Htomsik/GO-REST-API-Sample/internal/app"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store/testStore"

	"github.com/Htomsik/GO-REST-API-Sample/internal/model"
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

func TestUserRepository_FindByEmail(t *testing.T) {
	st := testStore.New()

	// Проверка на то что не находит рандом юзера
	email := "user@ex.com"
	_, err := st.User().FindByEmail(email)

	assert.EqualError(t, err, app.RecordNotFound.Error())

	// Создание + Провекра на то что находит юзера
	user := model.TestUser(t)
	user.Email = email

	// Проверка после создания
	st.User().Add(user)

	user, err = st.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
