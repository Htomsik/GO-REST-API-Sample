package apiServer

import (
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store/testStore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandeUsersAdd(t *testing.T) {
	// Arrange
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/users", nil)
	server := newServer(testStore.New())

	// Act
	server.ServeHTTP(recorder, request)

	// Assert
	assert.Equal(t, recorder.Code, http.StatusOK)
}
