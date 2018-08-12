package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"go_projects/api_server/handler"
	"go_projects/api_server/model"
	"go_projects/api_server/pkg/errno"
	"go_projects/api_server/util"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/12 下午5:02'
*/

func Get(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqId(c)})

	username := c.Param("username")

	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, user)
}
