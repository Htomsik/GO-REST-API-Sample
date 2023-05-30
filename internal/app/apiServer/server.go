package apiServer

import (
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

func (srv *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	srv.router.ServeHTTP(writer, request)
}

func (srv *server) configureRouter() {
	srv.router.HandleFunc("/users", srv.handleUsersAdd()).Methods(http.MethodPost)
}

func (srv *server) handleUsersAdd() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
