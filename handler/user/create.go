package user

import (
	"fmt"
	"log"

	. "github.com/apiserver/handler"
	"github.com/apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Printf("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Printf("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Printf("Header Content-Type: %s", contentType)

	log.Printf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
