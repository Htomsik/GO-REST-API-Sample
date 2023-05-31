package apiServer

import (
	"encoding/json"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// server ...
type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// newServer ...
func newServer(store store.Store) *server {
	srv := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	srv.configureRouter()

	return srv
}

func (srv *server) configureRouter() {
	srv.router.HandleFunc("/users", srv.handleUsersAdd()).Methods(http.MethodPost)
}

func (srv *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	srv.router.ServeHTTP(writer, request)
}

// error called if there are any errors in the request
func (srv *server) error(writer http.ResponseWriter, request *http.Request, code int, err error) {
	srv.respond(writer, request, code, map[string]string{"error": err.Error()})
}

// respond respond on request
func (srv *server) respond(writer http.ResponseWriter, request *http.Request, code int, data interface{}) {
	writer.WriteHeader(code)

	if data != nil {
		json.NewEncoder(writer).Encode(data)
	}
}
