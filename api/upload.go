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
		Status:  entity.STATUS_CREATED,
		Name:    object.Name,
		Type:    entity.TYPE_OTHER,
		Hash:    object.Hash,
		Size:    object.Size,
		Ext:     object.Ext,
	}

	_, err = resource.Create()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, Error(err))
	}

	context.JSON(http.StatusOK, Success(nil))
}
