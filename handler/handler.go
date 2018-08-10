package handler

import (
	"github.com/gin-gonic/gin"
	"go_projects/api_server/pkg/errno"
	"net/http"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/10 下午12:48'
*/

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.Decode(err)

	// 永远返回http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
