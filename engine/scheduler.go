package engine

import (
	"context"
	"fmt"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/setup"
	"github.com/gin-gonic/gin"
)

const Version = "0.0.1-dev"

type Scheduler struct {
	master     bool
	nodes      []node
	pluggers   []setup.Plugger
	routes     []setup.Routable
	httpServer *gin.Engine
}

func (sc *Scheduler) Run(ctx context.Context) {
	go func() {
		sc.httpServer.Run(address())
	}()

	<-ctx.Done()
}

func (sc *Scheduler) IsMaster() bool {
	return sc.master
}

func (sc *Scheduler) Bootstrap() error {
	sc.registerKernelPlugger()

	for _, plugger := range sc.pluggers {
		if err := plugger.Plug(); err != nil {
			return err
		}
	}

	for _, route := range sc.routes {
		route.ApplyRoutes(sc.httpServer)
	}

	return nil
}

func (sc *Scheduler) registerKernelPlugger() {
	kernelStarters := []setup.Plugger{
		config.NewPlugger(),
	}

	sc.pluggers = append(kernelStarters, sc.pluggers...)
}

func (sc *Scheduler) Register(feature interface{}) {
	sc.register(feature)
}

func (sc *Scheduler) register(feature interface{}) {
	if s, ok := feature.(setup.Plugger); ok {
		sc.pluggers = append(sc.pluggers, s)
	}

	if s, ok := feature.(setup.Routable); ok {
		sc.routes = append(sc.routes, s)
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
