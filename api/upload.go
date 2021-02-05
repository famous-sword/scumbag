package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
