package api

import (
	"github.com/famous-sword/scumbag/entity"
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
		context.AbortWithStatusJSON(http.StatusInternalServerError, Error(err))
	}

	resource := &entity.Resource{
		MediaId: object.Id(),
		Status:  entity.StatusCreated,
		Name:    object.Name,
		Type:    entity.TypeOther,
		Hash:    object.Hash,
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
