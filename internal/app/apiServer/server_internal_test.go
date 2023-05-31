package apiServer

import (
	"bytes"
	"encoding/json"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/model"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store/testStore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandeUsersAdd(t *testing.T) {
	// Arrange
	srv := newServer(testStore.New(), sessions.NewCookieStore([]byte("testSecretKey")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "ex@ex.com",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "nonePayload",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid payload params",
			payload: map[string]string{
				"email": "notEmail",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	// Act
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			bytesPayload := &bytes.Buffer{}
			json.NewEncoder(bytesPayload).Encode(testCase.payload)

			request, _ := http.NewRequest(http.MethodPost, "/users", bytesPayload)

			srv.ServeHTTP(recorder, request)

			// Assert
			assert.Equal(t, testCase.expectedCode, recorder.Code)
		})
	}

}

func TestServer_HandeSessions(t *testing.T) {
	// Arrange
	user := model.TestUser(t)

	store := testStore.New()
	store.User().Add(user)

	srv := newServer(store, sessions.NewCookieStore([]byte("testSecretKey")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    user.Email,
				"password": user.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "nonePayload",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid payload params",
			payload: map[string]string{
				"email": "notEmail",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    user.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	// Act
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			bytesPayload := &bytes.Buffer{}
			json.NewEncoder(bytesPayload).Encode(testCase.payload)

			request, _ := http.NewRequest(http.MethodPost, "/sessions", bytesPayload)

			srv.ServeHTTP(recorder, request)

			// Assert
			assert.Equal(t, testCase.expectedCode, recorder.Code)
		})
	}

}
