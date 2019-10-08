package doom

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

type Instance struct {
	conf   Config
	stdout io.ReadCloser
	stderr io.ReadCloser
	stdin  io.WriteCloser
	cmd    *exec.Cmd
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

func (i *Instance) Attach(in <-chan string, out chan<- string) error {
	go func() {
		scanner := bufio.NewScanner(i.stdout)

		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	for msg := range in {
		io.WriteString(i.stdin, msg+"\n")
	}

	return nil
}

func (i *Instance) Conf() Config {
	return i.conf
}
