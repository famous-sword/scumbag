package trancoding

import "github.com/famous-sword/scumbag/storage/warp"

type Transcoder interface {
	Transcode(object *warp.Object) error
}

type Context struct {
	transcoder Transcoder
}

func (context *Context) Execute(object *warp.Object) (err error) {
	return context.transcoder.Transcode(object)
}

func NewContext(transcoder Transcoder) *Context {
	return &Context{
		transcoder: transcoder,
	}
}
