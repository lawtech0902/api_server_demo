package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/11 下午5:24'
*/

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqId(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
