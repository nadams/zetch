package doom

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"syscall"
	"time"
)

type Instance struct {
	conf      Config
	stdoutpr  *io.PipeReader
	stdoutpw  *io.PipeWriter
	stdinpr   *io.PipeReader
	stdinpw   *io.PipeWriter
	stdoutBuf *bytes.Buffer
	cmd       *exec.Cmd
}

func NewInstance(c Config) *Instance {
	stdoutpr, stdoutpw := io.Pipe()
	stdinpr, stdinpw := io.Pipe()
	stdoutBuf := new(bytes.Buffer)

	go func() {
		io.Copy(ioutil.Discard, stdoutpr)
	}()

	return &Instance{
		conf:      c,
		stdoutpr:  stdoutpr,
		stdoutpw:  stdoutpw,
		stdinpr:   stdinpr,
		stdinpw:   stdinpw,
		stdoutBuf: stdoutBuf,
	}
}

func (i *Instance) Running() bool {
	return i.cmd != nil && (i.cmd.ProcessState == nil || i.cmd.ProcessState.Exited())
}

func (i *Instance) Stop() error {
	if i.cmd != nil && i.cmd.ProcessState == nil {
		if err := i.cmd.Process.Signal(syscall.SIGTERM); err != nil {
			return err
		}

		for {
			if i.cmd.ProcessState != nil {
				break
			}

			time.Sleep(time.Millisecond * 250)
		}
	}

	return nil
}

func (i *Instance) Start() error {
	if i.cmd == nil || i.cmd.ProcessState != nil {
		i.cmd = exec.Command("zandronum-server", i.conf.Args()...)
		i.cmd.Stdout = io.MultiWriter(i.stdoutpw, i.stdoutBuf)
		i.cmd.Stdin = i.stdinpr

		return i.cmd.Run()
	}

	return nil
}

func (i *Instance) Attach(ctx context.Context, out io.WriteCloser, in io.ReadCloser) error {
	copyStdout := func(out io.Writer) error {
		_, err := io.Copy(out, bytes.NewBuffer(i.stdoutBuf.Bytes()))

		return err
	}

	if i.cmd == nil || i.cmd.ProcessState != nil {
		defer out.Close()

		return copyStdout(out)
	}

	go func() {
		if _, err := io.Copy(i.stdinpw, NewReader(ctx, in)); err != nil {
			log.Println(err)
		}
	}()

	if err := copyStdout(out); err != nil {
		return err
	}

	_, err := io.Copy(NewWriter(ctx, out), i.stdoutpr)
	return err
}

func (i *Instance) Conf() Config {
	return i.conf
}
