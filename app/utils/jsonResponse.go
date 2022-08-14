package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonSuccessResponse(c *gin.Context, data interface{}, msg string, code int) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"code": code,
		"data": data,
	})
}
