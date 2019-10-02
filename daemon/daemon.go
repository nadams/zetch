package daemon

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	defer func() {
		os.RemoveAll(d.socket)
	}()

	l, err := net.Listen("unix", d.socket)
	if err != nil {
		return err
	}
	defer l.Close()

	d.server = grpc.NewServer()
	proto.RegisterDaemonServer(d.server, &Server{})

	log.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-quit

		log.Println("quitting...")

		d.server.Stop()
	}()

	return d.server.Serve(l)
}
