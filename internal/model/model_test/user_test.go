package model_test

import (
	"github.com/Htomsik/GO-REST-API-Sample/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestUser_BeforeAdd ...
func TestUser_BeforeAdd(t *testing.T) {
	user := model.TestUser(t)

	assert.NoError(t, user.BeforeAdd())
	assert.NotEmpty(t, user.EncryptedPassword)
}
