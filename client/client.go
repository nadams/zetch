package client

import (
	"context"
	"io"
	"net"
	"time"

	"google.golang.org/grpc"

	"gitlab.node-3.net/nadams/zetch/proto"
)

type Client struct {
	socketLoc string
	out       io.Writer
	conn      *grpc.ClientConn
	client    proto.DaemonClient
}

func New(socketLoc string, out io.Writer) *Client {
	return &Client{
		socketLoc: socketLoc,
		out:       out,
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

func (c *Client) Version() (*proto.ListResponse, error) {
	return c.client.List(context.Background(), &proto.ListRequest{})
}
