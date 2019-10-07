package daemon

import (
	"context"
	"fmt"

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
		servers = append(servers, &proto.Server{
			Name:     instance.Conf.Name,
			Port:     fmt.Sprintf("%d", instance.Conf.Port),
			Pwads:    instance.Conf.WADs,
			GameType: string(instance.Conf.Mode),
		})
	}

	return &proto.ListResponse{Servers: servers}, nil
}
