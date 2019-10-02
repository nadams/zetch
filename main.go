package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/alecthomas/kingpin.v2"

	"gitlab.node-3.net/nadams/zetch/client"
	"gitlab.node-3.net/nadams/zetch/daemon"
)

var (
	app         = kingpin.New("zetch", "A Zandronum server manager")
	zetchSocket = app.Flag("socket", "Path of socket").Default("/tmp/zetch").String()

	daemonMode = app.Command("daemon", "Run this as a daemon")

	list = app.Command("list", "List running servers")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case daemonMode.FullCommand():
		var socket string
		if zetchSocket == nil {
			socket = filepath.Join(os.TempDir(), fmt.Sprintf("zetch"))
		} else {
			socket = *zetchSocket
		}
		defer func() {
			os.RemoveAll(socket)
		}()

		d := daemon.New(socket)
		if err := d.Listen(); err != nil {
			log.Println(err)
			return
		}
	case list.FullCommand():
		if err := clientAction(func(c *client.Client) error {
			resp, err := c.Version()
			if err != nil {
				return err
			}

			spew.Dump(resp)

			return nil
		}); err != nil {
			log.Println(err)
		}
	}
}

func clientAction(action func(c *client.Client) error) error {
	c := client.New(*zetchSocket)
	if err := c.Open(); err != nil {
		return err
	}
	defer c.Close()

	return action(c)
}
