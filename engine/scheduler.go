package engine

import (
	"context"
	"fmt"
	"github.com/famous-sword/scumbag/api"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/plugger"
	"github.com/gin-gonic/gin"
)

const Version = "0.0.1-dev"

type Scheduler struct {
	master     bool
	pluggers   []plugger.Plugger
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

	sc.httpServer = api.Uploader()

	return nil
}

func (sc *Scheduler) registerKernelPlugger() {
	kernelStarters := []plugger.Plugger{
		config.NewPlugger(),
	}

	sc.pluggers = append(kernelStarters, sc.pluggers...)
}

func (sc *Scheduler) Register(plugger plugger.Plugger) {
	sc.register(plugger)
}

func (sc *Scheduler) register(plugger plugger.Plugger) {
	sc.pluggers = append(sc.pluggers, plugger)
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func address() string {
	return fmt.Sprintf("%s:%d", config.String("web.host"), config.Integer("web.port"))
}
