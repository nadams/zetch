package daemon

import (
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/nadams/zetch/proto"
)

type Daemon struct {
	socket string
	server *grpc.Server
}

func New(socket string) *Daemon {
	return &Daemon{
		socket: socket,
	}
}

func (d *Daemon) Listen() error {
	if err := os.RemoveAll(d.socket); err != nil {
		return err
	}

	l, err := net.Listen("unix", d.socket)
	if err != nil {
		return err
	}
	defer l.Close()

	d.server = grpc.NewServer()
	proto.RegisterDaemonServer(d.server, &Server{})

	return d.server.Serve(l)
}
