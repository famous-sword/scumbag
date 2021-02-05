package api

import (
	"github.com/gin-gonic/gin"
)

func Error(err error) gin.H {
	return gin.H{
		"code":    0,
		"message": err.Error(),
	}
}

func Success(data interface{}) gin.H {
	return gin.H{
		"code":    1,
		"message": "success",
		"data":    data,
	}
}
