package api

import (
	"github.com/famous-sword/scumbag/stroage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(context *gin.Context) {
	request := context.Request
	body := request.Body
	name := request.Header.Get("Content-Name")
	// todo: check hash for fast upload
	_ = request.Header.Get("Content-Hash")

	object := stroage.NewObject()
	object.Name = name
	object.Read(body)

	err := stroage.Adapter().Put(object)

	if err != nil {
		context.JSON(http.StatusInternalServerError, Error(err))
	} else {
		context.JSON(http.StatusOK, Success(nil))
	}
}
