package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/shibukawa/configdir"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/nadams/zetch/client"
	"github.com/nadams/zetch/daemon"
)

var (
	app    = kingpin.New("zetch", "A Zandronum server manager")
	socket = app.Flag("socket", "Path of socket").Default(filepath.Join(configDir(), "zetch.socket")).String()

	daemonMode   = app.Command("daemon", "Run this as a daemon")
	daemonConfig = daemonMode.Flag("config-dir", "Configuration directory").Default(configDir()).String()

	list         = app.Command("list", "List running servers")
	attach       = app.Command("attach", "Attach to server output")
	attachName   = attach.Arg("name", "Name of server").Required().String()
	stop         = app.Command("stop", "Stop servers")
	stopNames    = stop.Arg("name", "Server name").Required().Strings()
	start        = app.Command("start", "Start servers")
	startNames   = start.Arg("name", "Server name").Required().Strings()
	restart      = app.Command("restart", "Start servers")
	restartNames = restart.Arg("name", "Server name").Required().Strings()
)

type Out interface {
	Out(io.Writer)
}

func configDir() string {
	folders := configdir.New("zetch", "zetch").QueryFolders(configdir.Global)
	if len(folders) > 0 {
		return folders[0].Path
	}

	return ""
}

func main() {
	app.HelpFlag.Short('h')
	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch parsed {
	case daemonMode.FullCommand():
		if err := daemon.New(*socket, nil).Listen(); err != nil {
			log.Println(err)
			return
		}
	default:
		c := client.New(*socket)
		if err := c.Open(); err != nil {
			log.Println(err)
			return
		}

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		go func() {
			<-quit

			log.Println("quitting...")

			c.Close()
		}()

		defer c.Close()

		var out Out
		var err error

		switch parsed {
		case list.FullCommand():
			out, err = c.List()
		case attach.FullCommand():
			if err := c.Attach(*attachName, os.Stdin, os.Stdout); err != nil {
				log.Println(err)
			}
			return
		case stop.FullCommand():
			if err := c.Stop(*stopNames...); err != nil {
				log.Println(err)
			}
			return
		case start.FullCommand():
			if err := c.Start(*startNames...); err != nil {
				log.Println(err)
			}
			return
		case restart.FullCommand():
			if err := c.Restart(*restartNames...); err != nil {
				log.Println(err)
			}
			return
		}

		if err != nil {
			log.Println(err)
			return
		}

		out.Out(os.Stdout)
	}
}
