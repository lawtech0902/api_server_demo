package user

import (
	"github.com/gin-gonic/gin"
	"go_projects/api_server/model"
	"go_projects/api_server/handler"
	"go_projects/api_server/pkg/errno"
	"go_projects/api_server/pkg/auth"
	"go_projects/api_server/pkg/token"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/12 下午9:27'
*/

func Login(c *gin.Context) {
	// 将数据与user struct进行绑定
	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Get the user information by the login username.
	d, err := model.GetUser(u.Username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// sign the jwt
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	handler.SendResponse(c, nil, model.Token{Token: t})
}
