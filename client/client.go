package client

import (
	"context"
	"fmt"
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
	t.SetHeader([]string{"id", "name", "mode", "pwads", "port"})

	for _, server := range l.servers {
		t.Append([]string{server.Id, server.Name, server.GameType, strings.Join(server.Pwads, ", "), server.Port})
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

func (c *Client) Attach(id string, out io.Writer) error {
	resp, err := c.client.Attach(context.Background(), &proto.AttachRequest{Id: id})
	if err != nil {
		return err
	}

	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fmt.Fprintln(out, msg.Msg)
	}

	return nil
}
