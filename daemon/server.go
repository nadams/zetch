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
		})
	}

	return &proto.ListResponse{Servers: servers}, nil
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
		r, w := io.Pipe()
		defer r.Close()
		defer w.Close()

		go func() {
			<-stream.Context().Done()

			w.Close()
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

		if err := instance.Attach(stream.Context(), w); err != nil {
			return err
		}

		log.Println("waiting for wg")
		wg.Wait()
		log.Println("done waiting for wg")

		//  out := make(chan string)
		//  in := make(chan string)

		//  go func() {
		//    _ = instance.Attach(in, out)
		//  }()

		//  go func() {
		//    for {
		//      req, err := stream.Recv()
		//      if err != nil {
		//        log.Println(err)
		//        return
		//      }

		//      if req.Msg != "" {
		//        in <- req.Msg
		//      }
		//    }
		//  }()

		//  for {
		//    select {
		//    case <-stream.Context().Done():
		//      close(out)
		//      close(in)
		//      return nil
		//    case msg := <-out:
		//      stream.Send(&proto.ServerOutput{Msg: msg})
		//    }
		//  }
	}

	return nil
}
