package apiServer

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(test *testing.T) {

	api := New(NewConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	api.handleHello().ServeHTTP(recorder, request)

	assert.Equal(test, recorder.Body.String(), "Hello")
}
