package apiServer

import "net/http"

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
	accountActiveEndpoint = accountEndpoint + "/active"
	accountWhoAmIEndpoint = accountActiveEndpoint + "/whoami"
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
}

// configureAccountActiveEndpoints Account endpoints with authentication + active middleware
func (srv *server) configureAccountActiveEndpoints() {
	accountActive := srv.router.PathPrefix(accountActiveEndpoint).Subrouter()
	accountActive.Use(srv.activeUserMiddleWare)

	srv.router.HandleFunc(accountWhoAmIEndpoint, srv.handleWhoAmI()).Methods(http.MethodGet)
}
