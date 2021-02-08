package trancoding

type Transcoder interface {
	Transcode() error
}
