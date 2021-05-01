package foundation

import "github.com/gin-gonic/gin"

// Bootable initializes resources
// by implementing the `Bootstrap` method
type Bootable interface {
	Bootstrap() (err error)
}

// Routable register routes by
// implementing the `Apply` method
type Routable interface {
	Apply(router *gin.Engine)
}
