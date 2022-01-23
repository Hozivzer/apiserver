package user

import (
	"strconv"

	. "github.com/apiserver/handler"
	"github.com/apiserver/model"
	"github.com/apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
