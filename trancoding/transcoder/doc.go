package transcoder

import (
	"github.com/famous-sword/scumbag/storage/warp"
	"os/exec"
	"path/filepath"
)

type DocTranscoder struct{}

func (doc DocTranscoder) Transcode(object *warp.Object) error {
	dir := filepath.Join(object.Hash, object.Hash, object.Name)
	src := filepath.Dir(dir)

	args := []string{
		"--invisible",
		"--convert-to",
		"jpg",
		"--outdir",
		dir,
		src,
	}

	command := exec.Command("soffice", args...)

	command.Run()

	return nil
}
