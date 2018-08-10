package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"go_projects/api_server/handler"
	"go_projects/api_server/pkg/errno"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/10 上午11:38'
*/

func Create(c *gin.Context) {
	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("password is empty"), nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	handler.SendResponse(c, nil, rsp)
}
