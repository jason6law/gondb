// Package classification gondb API.
//
// The purpose of this service is to provide an application
// that is storing structured data on some key-value store,
// like redis.
//
//		Schemes: http
// 		Host: localhost
//		BasePath: /gondb
//     	Version: 0.0.1
//
// swagger:meta
package service

import "net/http"

// swagger:parameters getSingleUser
type GetUserParam struct {
	// an id of user info
	//
	// Required: true
	// in:query
	Id int `json:"id"`
}

func GetOneUser(w http.ResponseWriter,r http.Request) {
	// swagger:route GET /user users getSingleUser
	//
	// get a user by userID
	//
	// This will show a user info
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: UserResponse
	vals := r.URL.Query()

}