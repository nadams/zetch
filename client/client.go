package client

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/nadams/zetch/proto"
)

type Client struct {
	socketLoc string
	conn      *grpc.ClientConn
	client    proto.DaemonClient
}

func New(socketLoc string) *Client {
	return &Client{
		socketLoc: socketLoc,
	}
}

func (c *Client) Open() error {
	conn, err := grpc.Dial(
		c.socketLoc,
		grpc.WithInsecure(),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}),
	)
	if err != nil {
		return err
	}

	c.conn = conn
	c.client = proto.NewDaemonClient(c.conn)

	return nil
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) List() (*proto.ListResponse, error) {
	return c.client.List(context.Background(), &proto.ListRequest{})
}
