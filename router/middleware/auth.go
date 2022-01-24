package middleware

import (
	"github.com/apiserver/handler"
	"github.com/apiserver/pkg/errno"
	"github.com/apiserver/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
