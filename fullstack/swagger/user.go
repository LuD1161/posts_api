package swagger

import (
	model "github.com/LuD1161/posts_api/fullstack/api/models"
)

// HTTP status code 200 and user model in data
// swagger:response userResp
type swaggUserResp struct {
	// in:body
	Body struct {
		// HTTP status code 200
		// User model
		model.User
	}
}

// HTTP status code 200 and an array of user models in data
// swagger:response usersResp
type swaggUsersResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - Status OK
		// Code int `json:"code"`
		// Array of user models
		Data []model.User `json:"data"`
	}
}
