package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go_projects/api_server/pkg/errno"
	"github.com/lexkong/log"
	"fmt"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/10 上午11:38'
*/

func Create(c *gin.Context) {
	var (
		r struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		err error
	)

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind}) // 返回json格式的response
		return
	}

	log.Debugf("username is [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).
			Add("This is add message.")
		log.Errorf(err, "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	code, message := errno.Decode(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
