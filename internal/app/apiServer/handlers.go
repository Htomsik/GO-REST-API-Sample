package apiServer

import (
	"encoding/json"
	"github.com/Htomsik/GO-REST-API-Sample/internal/model"
	"net/http"
)

// handleUsersAdd Add new user
func (srv *server) handleUsersAdd() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			srv.error(w, r, http.StatusBadRequest, err)
			return
		}

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
