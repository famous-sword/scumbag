package engine

var pluggers []Plugger

type Plugger interface {
	Plug() (err error)
}

func Bootstrap() error {
	registerKernelPlugger()

	for _, starter := range pluggers {
		if err := starter.Plug(); err != nil {
			return err
		}
	}

	return nil
}

func registerKernelPlugger() {
	kernelStarters := []Plugger{
		NewConfigPlugger(),
	}

	pluggers = append(kernelStarters, pluggers...)
}

func Register(plugger Plugger) {
	register(plugger)
}

func register(plugger Plugger) {
	pluggers = append(pluggers, plugger)
}
