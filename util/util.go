package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-rRequest-Id")

	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

func GenShortId() (string, error) {
	return shortid.Generate()
}
