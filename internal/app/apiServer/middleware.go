package apiServer

import (
	"context"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/model"
	"net/http"
)

const (
	ctxKeyUser ctxKey = iota
)

type ctxKey int8

func (srv *server) authenticateUserMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		// Get user session from request
		session, err := srv.sessionStore.Get(request, sessionName)
		if err != nil {
			srv.error(writer, request, http.StatusInternalServerError, err)
			return
		}

		// Get user id from session
		userId, exist := session.Values[userIdSessionValue]
		if !exist {
			srv.error(writer, request, http.StatusUnauthorized, model.NotAuthenticated)
			return
		}

		// Check user id exist
		user, err := srv.store.User().Find(userId.(int))
		if err != nil {
			srv.error(writer, request, http.StatusUnauthorized, model.NotAuthenticated)
			return
		}

		// Throw user context next
		next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, user)))
	})
}
