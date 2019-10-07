package client

import (
	"context"
	"io"
	"net"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
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

type ListResponse struct {
	servers []*proto.Server
}

func (l *ListResponse) Out(w io.Writer) {
	t := tablewriter.NewWriter(w)
	t.SetHeader([]string{"name", "mode", "pwads", "port"})

	for _, server := range l.servers {
		t.Append([]string{server.Name, server.GameType, strings.Join(server.Pwads, ", "), server.Port})
	}

	t.Render()
}

func (c *Client) List() (*ListResponse, error) {
	resp, err := c.client.List(context.Background(), &proto.ListRequest{})
	if err != nil {
		return nil, err
	}

	return &ListResponse{servers: resp.Servers}, nil
}
