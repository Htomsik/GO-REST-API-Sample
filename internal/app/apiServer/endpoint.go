package apiServer

import (
	"encoding/json"
	"github.com/Htomsik/GO-REST-API-Sample/internal/app/model"
	"net/http"
)

const (
	sessionName        = "GO-REST-API-Sample"
	userIdSessionValue = "userId"
)

// handleUsersAdd Add new user
func (srv *server) handleUsersAdd() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode request
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			srv.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Add new user
		user := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := srv.store.User().Add(user); err != nil {
			srv.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		user.ClearPrivate()
		srv.respond(w, r, http.StatusCreated, user)
	}
}

// handleSessionsAdd Add new user session
func (srv *server) handleSessionsAdd() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode request
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			srv.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Find user
		user, err := srv.store.User().FindByEmail(req.Email)
		if err != nil || !user.ComparePassword(req.Password) {
			srv.error(w, r, http.StatusUnauthorized, model.EmailOrPasswordIncorrect)
			return
		}

		// Create user session
		session, err := srv.sessionStore.Get(r, sessionName)
		if err != nil {
			srv.error(w, r, http.StatusInternalServerError, err)
		}

		session.Values[userIdSessionValue] = user.ID

		// Add user session into sessions store
		if err = srv.sessionStore.Save(r, w, session); err != nil {
			srv.error(w, r, http.StatusInternalServerError, err)
		}

		srv.respond(w, r, http.StatusOK, nil)
	}
}
