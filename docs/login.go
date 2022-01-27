package docs

import (
	"github.com/apiserver/handler"
	"github.com/apiserver/handler/user"
)

// swagger:route POST /login login userLoginResponse
// 用户登录.
// responses:
//   200: userLoginResponse
//   default: errResponse

// This text will appear as description of your request body.
// swagger:parameters userLoginResponse
type userParamsWrapperLogin struct {
	// in:body
	Body user.CreateRequest
}

// This text will appear as description of your response body.
// swagger:response userLoginResponse
type getUserResponseWrapperLogin struct {
	// in:body
	Body handler.Response
}

// This text will appear as description of your error response body.
// swagger:response errResponse
type errResponseWrapperLogin struct {
	// Error code.
	Code int `json:"code"`

	// Error message.
	Message string `json:"message"`
}
