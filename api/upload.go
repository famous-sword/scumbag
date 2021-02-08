package api

import (
	"github.com/famous-sword/scumbag/entity"
	"github.com/famous-sword/scumbag/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var bucket = "public"

func Upload(context *gin.Context) {
	request := context.Request
	body := request.Body
	defer body.Close()
	name := request.Header.Get("Content-Name")
	digest := request.Header.Get("Digest")
	// todo: check hash for fast upload
	_ = strings.Split(digest, "=")[1]

	object := storage.NewObject()
	object.Name = name
	object.Read(body)

	err := storage.Adapter().Put(bucket, object)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, Error(err))
	}

	resource := &entity.Resource{
		Uuid:   object.Id(),
		Status: entity.StatusCreated,
		Name:   object.Name,
		Hash:   object.Hash,
		Ext:    object.Ext,
		Bucket: bucket,
	}

	meta := &entity.Meta{
		Size: object.Size,
	}

	_, err = resource.Create(meta)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, Error(err))
	}

	context.JSON(http.StatusOK, Success(nil))
}
