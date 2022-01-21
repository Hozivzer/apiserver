package user

import (
	"log"

	. "github.com/apiserver/handler"
	"github.com/apiserver/model"
	"github.com/apiserver/pkg/errno"
	"github.com/apiserver/util"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	log.Printf("User Create function called." + "X-Request-Id:" + util.GetReqID(c))
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	//验证数据
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//加密密码
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//插入数据库
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
