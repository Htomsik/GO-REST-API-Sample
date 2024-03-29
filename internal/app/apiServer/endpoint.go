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

// handleAccountAdd Create Account
// @Summary      Add account
// @Description  Create new account
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param		 User 	body 		model.UserShort 	true 	"user information"
// @Success      201  	{object} 	model.User
// @Failure      422
// @Failure      400
// @Router       /account [post]
func (srv *server) handleAccountAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode request
		req := &model.UserShort{}

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

// handleSessionsAdd Authorize
// @Summary      Authorize into account
// @Description  Authorize into account
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param		 User 	body	model.UserShort 	true 	"user information"
// @Failure      401
// @Failure      400
// @Failure      500
// @Success      200
// @Router       /account/session [post]
func (srv *server) handleSessionsAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode request
		req := &model.UserShort{}
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

// handleWho info about current authorised user
// @Summary      Account info
// @Description  info about current user
// @Tags         Account/Active
// @Accept       json
// @Produce      json
// @Success      200	{object}	model.User
// @Failure      401
// @Router       /account/active/who [get]

func (srv *server) handleWho() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		srv.respond(writer, request, http.StatusOK, request.Context().Value(ctxKeyUser).(*model.User))
	}
}

// handleAccountDeactivate deactivate current active account
// @Summary      Deactivate account
// @Description  Only deactivate, not delete
// @Tags         Account/Active
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      422
// @Failure      401
// @Router       /account/active/deactivate [post]

func (srv *server) handleAccountDeactivate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxKeyUser).(*model.User)

		if err := srv.store.User().Deactivate(user.ID); err != nil {
			srv.error(w, r, http.StatusUnprocessableEntity, err)
		}

		srv.respond(w, r, http.StatusOK, nil)
	}
}

// handleAccountActivate activate nonactive account
// @Summary      Activate account
// @Description  Activate only deactivated accounts
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param		 User 	body	model.UserShort 	true 	"user information"
// @Success      200
// @Failure      422
// @Router       /account/activate [post]
func (srv *server) handleAccountActivate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxKeyUser).(*model.User)

		if err := srv.store.User().Activate(user.ID); err != nil {
			srv.error(w, r, http.StatusUnprocessableEntity, err)
		}

		srv.respond(w, r, http.StatusOK, nil)
	}
}
