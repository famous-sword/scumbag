package api

import "github.com/gin-gonic/gin"

func Uploader() *gin.Engine {
	app := gin.Default()
	app.PUT("/upload", Upload)

	return app
}
