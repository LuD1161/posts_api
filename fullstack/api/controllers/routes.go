package controllers

import (
	"github.com/LuD1161/posts_api/fullstack/api/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddleWareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/login", middlewares.SetMiddleWareJSON(s.Login)).Methods("POST")

	//User routes
	s.Router.HandleFunc("/user", middlewares.SetMiddleWareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/user", middlewares.SetMiddleWareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddleWareJSON(s.GetUser)).Methods("GET")
	// swagger:operation GET /user/{id} users GetUser
	// ---
	// summary: Returns user details
	// description: Return user details
	// parameters:
	// - name: id
	//   in: path
	//   description: user id
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
	//   "404":
	//     "$ref": "#/responses/notFound"
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddleWareJSON(middlewares.SetMiddleWareAuthentication(s.UpdateUser))).Methods("PUT")
	// swagger:operation PUT /user/{id} users UpdateUser
	// ---
	// summary: Update user details
	// description: Update user details
	// parameters:
	// - name: email
	//   in: body
	//   description: email to update
	//   type: string
	//   required: false
	// - name: password
	//   in: body
	//   description: password to update
	//   type: string
	//   required: false
	// - name: nickname
	//   in: body
	//   description: nickname to update
	//   type: string
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
	//   "404":
	//     "$ref": "#/responses/notFound"
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddleWareAuthentication(s.DeleteUser)).Methods("DELETE")
	// swagger:operation DELETE /user/{id} users DeleteUser
	// ---
	// summary: Delete user details
	// description: Delete user details
	// parameters:
	// - name: id
	//   in: path
	//   description: user id
	//   type: string
	//   required: true
	// responses:
	//   "204":
	//   "404":
	//     "$ref": "#/responses/notFound"

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddleWareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddleWareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddleWareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddleWareJSON(middlewares.SetMiddleWareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddleWareAuthentication(s.DeletePost)).Methods("DELETE")
}
