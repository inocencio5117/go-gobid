package api

import (
	"net/http"

	"github.com/inocencio5117/go-gobid/internal/jsonutils"
)

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "must be authenticated",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
