package plugger

type Plugger interface {
	Plug() (err error)
}
