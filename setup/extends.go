package setup

import "github.com/gin-gonic/gin"

type Plugger interface {
	Plug() (err error)
}

type Routable interface {
	ApplyRoutes(router *gin.Engine)
}
