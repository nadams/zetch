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
		conf := instance.Conf()
		servers = append(servers, &proto.Server{
			Id:       instance.ID(),
			Name:     conf.Name,
			Port:     fmt.Sprintf("%d", conf.Port),
			Pwads:    conf.WADs,
			GameType: string(conf.Mode),
		})
	}

	return &proto.ListResponse{Servers: servers}, nil
}
