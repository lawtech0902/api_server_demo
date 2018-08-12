package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"go_projects/api_server/handler"
	"go_projects/api_server/pkg/errno"
	"go_projects/api_server/service"
	"go_projects/api_server/util"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/12 下午4:50'
*/

func List(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqId(c)})

	var r ListRequest

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Limit, r.Offset)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
