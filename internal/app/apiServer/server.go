package apiServer

import (
	"encoding/json"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/store"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	usersEndpoint    = "/users"
	sessionsEndpoint = "/sessions"
	homeEndpoint     = "/"
)

// Account endpoints
const (
	accountEndPoint           = "/account"
	accountActiveEndPoint     = "/active"
	accountWhoAmIEndPoint     = "/whoami"
	accountDeactivateEndPoint = "/deactivate"
)

// server ...
type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

// newServer ...
func newServer(store store.Store, sessionStore sessions.Store) *server {
	srv := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	srv.configureRouter()

	return srv
}

func (srv *server) configureRouter() {
	// Add access with different domains
	srv.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	// Add requestID for all endpoints
	srv.router.Use(srv.requestIDMiddleWare)

	// Add logger middleware for all endpoints
	srv.router.Use(srv.logRequestMiddleWare)

	// Public endpoints
	srv.router.HandleFunc(usersEndpoint, srv.handleUsersAdd()).Methods(http.MethodPost)
	srv.router.HandleFunc(sessionsEndpoint, srv.handleSessionsAdd()).Methods(http.MethodPost)

	// Account endPoints with authentication middleware
	account := srv.router.PathPrefix(accountEndPoint).Subrouter()
	account.Use(srv.authenticateUserMiddleWare)

	// Account endPoints with authentication middleware
	accountActive := srv.router.PathPrefix(accountEndPoint + accountActiveEndPoint).Subrouter()
	accountActive.Use(srv.activeUserMiddleWare)

	accountActive.HandleFunc(accountWhoAmIEndPoint, srv.handleWhoAmI()).Methods(http.MethodGet)
}

func (srv *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	srv.router.ServeHTTP(writer, request)
}

// error called if there are any errors in the request
func (srv *server) error(writer http.ResponseWriter, request *http.Request, code int, err error) {
	srv.respond(writer, request, code, map[string]string{"error": err.Error()})
}

// respond on request
func (srv *server) respond(writer http.ResponseWriter, request *http.Request, code int, data interface{}) {
	writer.WriteHeader(code)

	if data != nil {
		json.NewEncoder(writer).Encode(data)
	}
}
