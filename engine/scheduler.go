package engine

import (
	"context"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/plugger"
)

type Scheduler struct {
	master   bool
	pluggers []plugger.Plugger
}

func (sc *Scheduler) Run(ctx context.Context) {
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
