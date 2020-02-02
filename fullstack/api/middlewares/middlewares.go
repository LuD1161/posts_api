package middlewares

import (
	"errors"
	"net/http"

	"github.com/LuD1161/posts_api/fullstack/api/auth"
	"github.com/LuD1161/posts_api/fullstack/api/responses"
)

// SetMiddleWareJSON : Set response header to application/json header
func SetMiddleWareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddleWareAuthentication : Check for authenticated users
func SetMiddleWareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
