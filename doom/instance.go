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
	conf   Config
	id     string
	stdout *circbuf.Buffer
	stderr *circbuf.Buffer
	cmd    *exec.Cmd
}

func NewInstance(c Config, id string) *Instance {
	return &Instance{
		conf: c,
		id:   id,
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

	args := i.conf.Args()
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

func (i *Instance) Conf() Config {
	return i.conf
}

func (i *Instance) ID() string {
	return i.id
}
