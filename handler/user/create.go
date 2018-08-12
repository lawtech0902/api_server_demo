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
__date__ = '2018/8/10 上午11:38'
*/

// 创建user
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqId(c)})

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 验证
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密password
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 将user插入数据库
	if err := u.CreateUser(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c, nil, resp)
}
