package controllers

import (
	"net/http"

	"github.com/LuD1161/posts_api/fullstack/api/responses"
)

// Home : Home route function defined here
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this Awesome API")
}
