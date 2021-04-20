package engine

import (
	"context"
	"fmt"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/foundation"
	"github.com/gin-gonic/gin"
)

const Version = "0.0.1-dev"

type Scheduler struct {
	master       bool
	nodes        []node
	bootstrapper []foundation.Bootable
	routes       []foundation.Routable
	httpServer   *gin.Engine
}

// Run run scheduler with a context
func (sc *Scheduler) Run(ctx context.Context) {
	go func() {
		sc.httpServer.Run(address())
	}()

	<-ctx.Done()
}

// IsMaster judge scheduler as s master node
func (sc *Scheduler) IsMaster() bool {
	return sc.master
}

// Bootstrap scheduler set up, apply kernel components
func (sc *Scheduler) Bootstrap() error {
	sc.setupKernel()

	for _, plugger := range sc.bootstrapper {
		if err := plugger.Bootstrap(); err != nil {
			return err
		}
	}

	for _, route := range sc.routes {
		route.Register(sc.httpServer)
	}

	return nil
}

// Register can register a feature
// feature can be a foundation.Bootable
// can also be a foundation.Routable
func (sc *Scheduler) Register(feature interface{}) {
	sc.register(feature)
}

func (sc *Scheduler) setupKernel() {
	bootstrappers := []foundation.Bootable{
		config.NewBootstrapper(),
	}

	sc.bootstrapper = append(bootstrappers, sc.bootstrapper...)
}

func (sc *Scheduler) register(feature interface{}) {
	if plugger, ok := feature.(foundation.Bootable); ok {
		sc.bootstrapper = append(sc.bootstrapper, plugger)
	}

	if routes, ok := feature.(foundation.Routable); ok {
		sc.routes = append(sc.routes, routes)
	}
}

func NewScheduler() *Scheduler {
	s := &Scheduler{}
	s.httpServer = gin.Default()

	return s
}

func address() string {
	return fmt.Sprintf("%s:%d", config.String("web.host"), config.Integer("web.port"))
}
