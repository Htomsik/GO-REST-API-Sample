package apiServer

import (
	_ "github.com/Htomsik/GO-REST-API-Sample/docs"
	_ "github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// Account endpoints
const (
	accountEndpoint         = "/account"
	accountActivateEndpoint = "/activate"
	accountSessionEndpoint  = "/session"
)

// Account/Active endpoints
const (
	accountActiveEndpoint     = accountEndpoint + "/active"
	accountWhoAmIEndpoint     = "/who"
	accountDeactivateEndpoint = "/deactivate"
)

// configureOtherEndpoints public endpoints
func (srv *server) configureOtherEndpoints() {
	srv.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

// configureAccountEndpoint Account endpoints with authentication middleware
func (srv *server) configureAccountEndpoint() {
	account := srv.router.PathPrefix(accountEndpoint).Subrouter()

	account.HandleFunc(accountActivateEndpoint, srv.handleAccountActivate()).Methods(http.MethodPut)
	account.HandleFunc("", srv.handleAccountAdd()).Methods(http.MethodPost)
	account.HandleFunc(accountSessionEndpoint, srv.handleSessionsAdd()).Methods(http.MethodPost)
}

// configureAccountActiveEndpoints Account endpoints with authentication + active middleware
func (srv *server) configureAccountActiveEndpoints() {
	accountActive := srv.router.PathPrefix(accountActiveEndpoint).Subrouter()
	accountActive.Use(srv.authenticateUserMiddleWare)
	accountActive.Use(srv.activeUserMiddleWare)

	accountActive.HandleFunc(accountWhoAmIEndpoint, srv.handleWho()).Methods(http.MethodGet)
	accountActive.HandleFunc(accountDeactivateEndpoint, srv.handleAccountDeactivate()).Methods(http.MethodPut)
}
