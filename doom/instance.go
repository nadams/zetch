package doom

import (
	"bytes"
	"context"
	"io"
	"log"
	"os/exec"
	"syscall"

	"github.com/hashicorp/go-multierror"
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
	var errs *multierror.Error

	if i.stdoutpr != nil {
		if err := i.stdoutpr.Close(); err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	if i.stdinpw != nil {
		if err := i.stdinpw.Close(); err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	if i.cmd != nil {
		if err := i.cmd.Process.Signal(syscall.SIGTERM); err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	return errs.ErrorOrNil()
}

func (i *Instance) Start() error {
	i.cmd = exec.Command("zandronum-server", i.conf.Args()...)
	i.cmd.Stdout = io.MultiWriter(i.stdoutpw, i.stdoutBuf)
	i.cmd.Stdin = i.stdinpr

	if err := i.cmd.Run(); err != nil {
		log.Println(err)
	}

	return nil
}

func (i *Instance) Attach(ctx context.Context, out io.Writer, in io.Reader) error {
	go func() {
		if _, err := io.Copy(i.stdinpw, NewReader(ctx, in)); err != nil {
			log.Println(err)
		}
	}()
	buf := bytes.NewBuffer(i.stdoutBuf.Bytes())
	if _, err := io.Copy(out, buf); err != nil {
		return err
	}

	_, err := io.Copy(NewWriter(ctx, out), i.stdoutpr)
	return err
}

func (i *Instance) Conf() Config {
	return i.conf
}
