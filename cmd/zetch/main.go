package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
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

	list = app.Command("list", "List running servers")
)

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
		defer c.Close()

		switch parsed {
		case list.FullCommand():
			spew.Dump(c.List())
		}
	}
}
