package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
	t.SetHeader([]string{"name", "hostname", "mode", "pwads", "port", "status"})

	for _, server := range l.servers {
		t.Append([]string{server.Name, server.HostName, server.GameType, strings.Join(server.Pwads, ", "), server.Port, server.Status})
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

func (c *Client) Attach(name string, in io.Reader, out io.Writer) error {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-quit
		log.Println("cancelling context")
		cancel()
	}()

	attachClient, err := c.client.Attach(ctx)
	if err != nil {
		return err
	}

	attachClient.Send(&proto.AttachRequest{Name: name})

	go func() {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			msg := scanner.Text()
			attachClient.Send(&proto.AttachRequest{Name: name, Msg: msg + "\n"})
		}
	}()

	for {
		msg, err := attachClient.Recv()
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

func (c *Client) Stop(name string) error {
	_, err := c.client.Stop(context.Background(), &proto.StopRequest{Name: name})
	return err
}
