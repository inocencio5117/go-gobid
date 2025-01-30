package api

import (
	"errors"
	"net/http"

	"github.com/inocencio5117/go-gobid/internal/jsonutils"
	"github.com/inocencio5117/go-gobid/internal/services"
	"github.com/inocencio5117/go-gobid/internal/usecase/user"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context(),
		data.UserName,
		data.Email,
		data.Password,
		data.Bio,
	)

	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrUsername) {
			_ = jsonutils.EncodeJson(w, r, http.StatusConflict, map[string]any{
				"error": "email or username already exists",
			})
			return
		}
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"user_id": id,
	})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.LoginUserReq](r)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.AuthenticateUser(r.Context(), data.Email, data.Password)

	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"error": "invalid email or password",
			})
			return
		}
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}

	err = api.Sessions.RenewToken(r.Context())

	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}

	api.Sessions.Put(r.Context(), "AuthenticatedUserId", id)

	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "user authenticated successfully",
	})
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	err := api.Sessions.RenewToken(r.Context())

	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "user logged out successfully",
	})
}
