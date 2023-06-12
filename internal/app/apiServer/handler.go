package apiServer

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Account endpoints
const (
	accountEndpoint = "/account"
)

// Other endpoints
const (
	usersEndpoint    = "/users"
	sessionsEndpoint = "/sessions"
)

// Account/Active endpoints
const (
	accountActiveEndpoint     = accountEndpoint + "/active"
	accountWhoAmIEndpoint     = "/whoami"
	accountDeactivateEndpoint = "/deactivate"
)

// configureOtherEndpoints public endpoints
func (srv *server) configureOtherEndpoints() {
	srv.router.HandleFunc(usersEndpoint, srv.handleUsersAdd()).Methods(http.MethodPost)
	srv.router.HandleFunc(sessionsEndpoint, srv.handleSessionsAdd()).Methods(http.MethodPost)
}

// configureAccountEndpoint Account endpoints with authentication middleware
func (srv *server) configureAccountEndpoint() {
	account := srv.router.PathPrefix(accountEndpoint).Subrouter()
	account.Use(srv.authenticateUserMiddleWare)

	srv.configureAccountActiveEndpoints(account)
}

// configureAccountActiveEndpoints Account endpoints with authentication + active middleware
func (srv *server) configureAccountActiveEndpoints(router *mux.Router) {
	accountActive := srv.router.PathPrefix(accountActiveEndpoint).Subrouter()
	accountActive.Use(srv.authenticateUserMiddleWare)
	accountActive.Use(srv.activeUserMiddleWare)

	accountActive.HandleFunc(accountWhoAmIEndpoint, srv.handleWhoAmI()).Methods(http.MethodGet)
	accountActive.HandleFunc(accountDeactivateEndpoint, srv.handleAccountDeactivate()).Methods(http.MethodPut)
}
