package messages

import (
	"encoding/json"
	"net/http"
)

// Errors

// Errors struct carries a slice of error struct which in turn are error messages that cnfrm with the json+vdn spec
type Errors struct {
	Errors []*Error `json:"errors"`
}

// Error struct carries deatiled error messages spoken of above
type Error struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// WriteError is a convenience function to write an error struct back to the requester
func WriteError(w http.ResponseWriter, err *Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(Errors{[]*Error{err}})
}

var (
	//Success 400 shorthand
	Success = &Error{"success", 200, "Success", "Action Performed successfully"}

	//ErrBadRequest 400 shorthand
	ErrBadRequest = &Error{"bad_request", 400, "Bad request", "Request body is not well-formed. It must be JSON."}
	//ErrNotAcceptable 406 shorthand
	ErrNotAcceptable = &Error{"not_acceptable", 406, "Not Acceptable", "Accept header must be set to 'application/vnd.api+json'."}
	//ErrUnsupportedMediaType 415 shorthand
	ErrUnsupportedMediaType = &Error{"unsupported_media_type", 415, "Unsupported Media Type", "Content-Type header must be set to: 'application/vnd.api+json'."}
	//ErrInternalServer 500 shorthand
	ErrInternalServer = &Error{"internal_server_error", 500, "Internal Server Error", "Something went wrong."}
	//ErrNoAuth 401 shorthand
	ErrNoAuth = &Error{"unauthorised", 401, "Unauthorised Access", "Not authenticated. Please login."}
	//ErrBadToken 401 shorthand
	ErrBadToken = &Error{"unauthorised", 401, "Unauthorised Access", "Not authenticated. Invalid Token."}
	//ErrNotFound 404 shorthand
	ErrNotFound = &Error{"not_found", 404, "Resource not Found", "Requested resource could not be found"}
	//ErrWrongPassword is a shorthand with no error code
	ErrWrongPassword = &Error{"wrong_password", http.StatusNotAcceptable, "Wrong Password", "Wrong Password"}
)
