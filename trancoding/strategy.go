package trancoding

type Transcoder interface {
	Transcode() error
}

type Context struct {
	transcoder Transcoder
}
