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
	stdoutBuf  bytes.Buffer
	stdoutPipe io.Reader
	stdout     io.ReadCloser
	stderr     io.ReadCloser
	stdin      io.WriteCloser
	cmd        *exec.Cmd
}

func NewInstance(c Config) *Instance {
	return &Instance{
		conf: c,
	}
}

func (i *Instance) Start() error {
	args := i.conf.Args()
	i.cmd = exec.Command("zandronum-server", args...)

	stdout, err := i.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	i.stdoutPipe = io.TeeReader(stdout, &i.stdoutBuf)

	stderr, err := i.cmd.StderrPipe()
	if err != nil {
		return err
	}

	stdin, err := i.cmd.StdinPipe()
	if err != nil {
		return err
	}

	i.stdout = stdout
	i.stderr = stderr
	i.stdin = stdin

	if err := i.cmd.Run(); err != nil {
		log.Println(err)
	}

	return nil
}

func (i *Instance) Attach(ctx context.Context, out io.Writer) error {
	if _, err := io.Copy(out, &i.stdoutBuf); err != nil {
		return err
	}

	_, err := io.Copy(NewWriter(ctx, out), i.stdoutPipe)
	return err
	//return err
}

func (i *Instance) Conf() Config {
	return i.conf
}
