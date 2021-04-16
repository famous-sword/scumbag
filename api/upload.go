package api

import (
	"github.com/famous-sword/scumbag/setup"
	"github.com/famous-sword/scumbag/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Uploader struct{}

func (u *Uploader) ApplyRoutes(router *gin.Engine) {
	router.PUT("/upload", u.upload)
}

func (u *Uploader) upload(context *gin.Context) {
	request := context.Request
	body := request.Body
	defer body.Close()
	name := request.Header.Get("Content-Name")
	digest := request.Header.Get("Digest")
	// todo: check hash for fast upload
	_ = strings.Split(digest, "=")[1]

	err := storage.Put("public", name, body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, Error.WithMessage(err.Error()))
	}

	context.JSON(http.StatusOK, Success)
}

func NewUploader() setup.Routable {
	return &Uploader{}
}
