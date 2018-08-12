package user

import (
	"github.com/gin-gonic/gin"
	"go_projects/api_server/handler"
	"go_projects/api_server/model"
	"go_projects/api_server/pkg/errno"
	"strconv"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/12 下午4:12'
*/

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
