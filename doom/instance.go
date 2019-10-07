package doom

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

type Instance struct {
	conf   Config
	id     string
	stdout io.ReadCloser
	stderr io.ReadCloser
	cmd    *exec.Cmd
}

func NewInstance(c Config, id string) *Instance {
	return &Instance{
		conf: c,
		id:   id,
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

	i.stdout = stdout
	i.stderr = stderr

	if err := i.cmd.Run(); err != nil {
		log.Println(err)
	}

	return nil
}

func (i *Instance) Attach(out chan<- string) error {
	scanner := bufio.NewScanner(i.stdout)

	for scanner.Scan() {
		out <- scanner.Text()
	}

	return nil
}

func (i *Instance) Conf() Config {
	return i.conf
}

func (i *Instance) ID() string {
	return i.id
}
