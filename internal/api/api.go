package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/inocencio5117/go-gobid/internal/services"
)

type Api struct {
	Router      *chi.Mux
	UserService services.UserService
}
