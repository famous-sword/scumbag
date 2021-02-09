package trancoding

type Transcoder interface {
	Transcode() error
}

type Context struct {
	transcoder Transcoder
}

func (context *Context) Execute () (err error) {
	return context.transcoder.Transcode()
}
