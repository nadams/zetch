package main

import (
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/davecgh/go-spew/spew"
	"github.com/nadams/zetch/client"
	"github.com/nadams/zetch/daemon"
)

var (
	app    = kingpin.New("zetch", "A Zandronum server manager")
	socket = app.Flag("socket", "Path of socket").Default("/tmp/zetch").String()

	daemonMode = app.Command("daemon", "Run this as a daemon")

	list = app.Command("list", "List running servers")
)

func main() {
	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch parsed {
	case daemonMode.FullCommand():
		if err := daemon.New(*socket).Listen(); err != nil {
			log.Println(err)
			return
		}
	default:
		c := client.New(*socket)
		if err := c.Open(); err != nil {
			log.Println(err)
			return
		}
		defer c.Close()

		switch parsed {
		case list.FullCommand():
			spew.Dump(c.List())
		}
	}
}
