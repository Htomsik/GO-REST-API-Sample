package apiServer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/model"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store/testStore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testCookieSecretKey = []byte("testSecretKey")
)

func TestServer_HandeUsersAdd(t *testing.T) {
	// Arrange
	srv := newServer(testStore.New(), sessions.NewCookieStore(testCookieSecretKey))

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

			request, _ := http.NewRequest(http.MethodPost, usersEndpoint, bytesPayload)

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

	srv := newServer(store, sessions.NewCookieStore([]byte(testCookieSecretKey)))

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

			request, _ := http.NewRequest(http.MethodPost, sessionsEndpoint, bytesPayload)

			srv.ServeHTTP(recorder, request)

			// Assert
			assert.Equal(t, testCase.expectedCode, recorder.Code)
		})
	}

}

func TestServer_AuthenticateUser(t *testing.T) {
	// Arrange
	store := testStore.New()
	user := model.TestUser(t)

	store.User().Add(user)

	srv := newServer(store, sessions.NewCookieStore(testCookieSecretKey))

	// Create new cookie for response
	sc := securecookie.New(testCookieSecretKey, nil)
	dummyHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				userIdSessionValue: user.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	// Act
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodGet, homeEndpoint, nil)

			// Generate cookie session key for response
			cookie, _ := sc.Encode(sessionName, testCase.cookieValue)
			request.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookie))

			srv.authenticateUserMiddleWare(dummyHandler).ServeHTTP(recorder, request)

			// Assert
			assert.Equal(t, testCase.expectedCode, recorder.Code)
		})
	}
}

func TestServer_ActiveUser(t *testing.T) {
	// Arrange
	store := testStore.New()
	ActiveUser := model.TestUser(t)
	NotActiveUser := model.TestUser(t)

	NotActiveUser.Email = "NotActiveUser@ex.com"

	store.User().Add(ActiveUser)
	store.User().Add(NotActiveUser)

	store.User().Deactivate(NotActiveUser.ID)

	srv := newServer(store, sessions.NewCookieStore(testCookieSecretKey))

	// Create new cookie for response
	sc := securecookie.New(testCookieSecretKey, nil)
	dummyHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "Active",
			cookieValue: map[interface{}]interface{}{
				userIdSessionValue: ActiveUser.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "not Active",
			cookieValue: map[interface{}]interface{}{
				userIdSessionValue: NotActiveUser.ID,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "nil",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "notExist",
			cookieValue: map[interface{}]interface{}{
				userIdSessionValue: -1,
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	// Act
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodGet, accountEndPoint+accountActiveEndPoint+accountWhoAmIEndPoint, nil)

			// Generate cookie session key for response
			cookie, _ := sc.Encode(sessionName, testCase.cookieValue)
			request.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookie))

			srv.authenticateUserMiddleWare(srv.activeUserMiddleWare(dummyHandler)).ServeHTTP(recorder, request)

			// Assert
			assert.Equal(t, testCase.expectedCode, recorder.Code)
		})
	}
}
