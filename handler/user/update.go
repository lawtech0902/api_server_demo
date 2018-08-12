package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"go_projects/api_server/handler"
	"go_projects/api_server/model"
	"go_projects/api_server/pkg/errno"
	"go_projects/api_server/util"
	"strconv"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/12 下午4:14'
*/

func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqId(c)})

	userId, _ := strconv.Atoi(c.Param("id"))

	// 绑定用户数据
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}

	u.Id = uint64(userId)

	// 验证
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密密码
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 保存修改
	if err := u.UpdateUser(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
