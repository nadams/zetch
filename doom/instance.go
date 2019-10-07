package doom

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/armon/circbuf"
)

const bufferSize = 1024 * 1024 // 1mb

type Instance struct {
	Conf   Config
	stdout *circbuf.Buffer
	stderr *circbuf.Buffer
	cmd    *exec.Cmd
}

func NewInstance(c Config) *Instance {
	return &Instance{
		Conf: c,
	}
}

func (i *Instance) Start() error {
	stdout, err := circbuf.NewBuffer(bufferSize)
	if err != nil {
		return err
	}

	stderr, err := circbuf.NewBuffer(bufferSize)
	if err != nil {
		return err
	}

	i.stdout = stdout
	i.stderr = stderr

	args := i.Conf.Args()
	i.cmd = exec.Command("zandronum-server", args...)
	i.cmd.Stdout = i.stdout
	i.cmd.Stderr = i.stderr

	if err := i.cmd.Run(); err != nil {
		log.Println(err)
	}

	return nil
}

func (i *Instance) Attach() error {
	if _, err := io.Copy(os.Stdout, bytes.NewReader(i.stdout.Bytes())); err != nil {
		return err
	}

	return nil
}
