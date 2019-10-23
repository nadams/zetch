package daemon

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/nadams/zetch/doom"
	"github.com/nadams/zetch/proto"
)

type Server struct {
	d *Daemon
}

func NewServer(d *Daemon) *Server {
	return &Server{d: d}
}

func (s *Server) List(_ context.Context, in *proto.ListRequest) (*proto.ListResponse, error) {
	s.d.m.Lock()
	defer s.d.m.Unlock()

	servers := make([]*proto.Server, 0, len(s.d.instances))
	for _, instance := range s.d.instances {
		conf := instance.Conf()
		servers = append(servers, &proto.Server{
			Name:     conf.Name,
			HostName: conf.Hostname,
			Port:     fmt.Sprintf("%d", conf.Port),
			Pwads:    conf.WADs,
			GameType: string(conf.Mode),
			Status: func() string {
				if instance.Running() {
					return "running"
				}
				return "stopped"
			}(),
		})
	}

	return &proto.ListResponse{Servers: servers}, nil
}

func (s *Server) Stop(ctx context.Context, in *proto.StopRequest) (*proto.StopResponse, error) {
	s.d.m.Lock()
	defer s.d.m.Unlock()

	for _, i := range s.d.instances {
		for _, name := range in.Names {
			if i.Conf().Name == name {
				log.Println("stopping server", i.Stop())
			}
		}
	}

	return &proto.StopResponse{}, nil
}

func (s *Server) Start(ctx context.Context, in *proto.StartRequest) (*proto.StartResponse, error) {
	s.d.m.Lock()
	defer s.d.m.Unlock()

	for _, i := range s.d.instances {
		for _, name := range in.Names {
			if i.Conf().Name == name {
				log.Println("starting server")

				go func() {
					if err := i.Start(); err != nil {
						log.Println(err)
					}
				}()
			}
		}
	}

	return &proto.StartResponse{}, nil
}

func (s *Server) Restart(ctx context.Context, in *proto.RestartRequest) (*proto.RestartResponse, error) {
	if _, err := s.Stop(ctx, &proto.StopRequest{Names: in.Names}); err != nil {
		return nil, err
	}
	if _, err := s.Start(ctx, &proto.StartRequest{Names: in.Names}); err != nil {
		return nil, err
	}
	return &proto.RestartResponse{}, nil
}

func (s *Server) Attach(stream proto.Daemon_AttachServer) error {
	var instance *doom.Instance
	recv, err := stream.Recv()
	if err != nil {
		return err
	}

	for _, i := range s.d.instances {
		if i.Conf().Name == recv.Name {
			instance = i
			break
		}
	}

	if instance != nil {
		var wg sync.WaitGroup
		r, stdout := io.Pipe()
		defer r.Close()
		defer stdout.Close()

		go func() {
			<-stream.Context().Done()

			stdout.Close()
			r.Close()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				stream.Send(&proto.ServerOutput{
					Msg: scanner.Text(),
				})
			}
		}()

		stdin, stdinW := io.Pipe()
		go func() {
			for {
				in, err := stream.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					log.Println(err)
					return
				}
				if _, err := io.WriteString(stdinW, in.Msg); err != nil {
					log.Println(err)
				}
			}
		}()

		if err := instance.Attach(stream.Context(), stdout, stdin); err != nil {
			return err
		}

		log.Println("waiting for wg")
		wg.Wait()
		log.Println("done waiting for wg")
	}

	return nil
}
