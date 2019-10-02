package daemon

import (
	"context"

	"gitlab.node-3.net/nadams/zetch/proto"
)

type Server struct{}

func (s *Server) List(_ context.Context, in *proto.ListRequest) (*proto.ListResponse, error) {
	return &proto.ListResponse{
		Servers: []*proto.Server{
			{
				Id:   "1",
				Name: "test",
			},
		},
	}, nil
}
