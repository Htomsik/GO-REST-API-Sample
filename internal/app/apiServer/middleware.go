package apiServer

import (
	"context"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/model"
	"github.com/google/uuid"
	"net/http"
)

const (
	ctxKeyUser ctxKey = iota
	ctxRequestId
)

const (
	requestIdHeader = "Request-ID"
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

func (srv *server) setRequestIDMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Generate new guid
		guid := uuid.New().String()

		// Set guid to header
		writer.Header().Set(requestIdHeader, guid)

		// Throw request id next
		next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxRequestId, guid)))
	})
}
