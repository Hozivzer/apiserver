package docs

import "github.com/apiserver/handler"

// swagger:route GET /v1/user/{name} user getUserRequest
// Get a user from memory.
// responses:
//   200: getUserResponse
//   default: errResponse

// This text will appear as description of your request url path.
// swagger:parameters getUserRequest
type getUserParamsWrapper struct {
	// in:path
	Name string `json:"name"`
}

// This text will appear as description of your response body.
// swagger:response getUserResponse
type getUserResponseWrapper struct {
	// in:body
	Body handler.Response
}

// This text will appear as description of your error response body.
// swagger:response errResponse
type errResponseWrapper struct {
	// Error code.
	Code int `json:"code"`

	// Error message.
	Message string `json:"message"`
}
