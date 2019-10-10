package doom

import (
	"bytes"
	"context"
	"io"
	"log"
	"os/exec"
)

type Instance struct {
	conf       Config
	stdoutBuf  *bytes.Buffer
	stdoutPipe io.Reader
	stdinPipe  io.Writer
	cmd        *exec.Cmd
}

func NewInstance(c Config) *Instance {
	return &Instance{
		conf:      c,
		stdoutBuf: &bytes.Buffer{},
	}
}

func (i *Instance) Start() error {
	pr, pw := io.Pipe()
	defer pr.Close()

	stdout := io.MultiWriter(pw, i.stdoutBuf)
	i.cmd = exec.Command("zandronum-server", i.conf.Args()...)
	i.cmd.Stdout = stdout
	i.stdoutPipe = pr

	stdin, err := i.cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer stdin.Close()

	i.stdinPipe = stdin

	if err := i.cmd.Run(); err != nil {
		log.Println(err)
	}

	return nil
}

func (i *Instance) Attach(ctx context.Context, out io.Writer) error {
	buf := bytes.NewBuffer(i.stdoutBuf.Bytes())
	if _, err := io.Copy(out, buf); err != nil {
		return err
	}

	_, err := io.Copy(NewWriter(ctx, out), i.stdoutPipe)
	return err
}

func (i *Instance) Conf() Config {
	return i.conf
}
