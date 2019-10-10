package daemon

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"

	"github.com/nadams/zetch/doom"
	"github.com/nadams/zetch/proto"
)

type Config struct {
	WADDir          string `toml:"wad_dir"`
	ServerBinary    string `toml:"server_binary"`
	ServerConfigDir string `toml:"server_config_dir"`
}

var defaultConfig = Config{
	WADDir:          "wads",
	ServerBinary:    "zandronum-server",
	ServerConfigDir: "conf.d",
}

type Daemon struct {
	m          sync.Mutex
	socket     string
	configPath *string
	config     *Config
	server     *grpc.Server
	instances  []*doom.Instance
}

func New(socket string, config *string) *Daemon {
	return &Daemon{
		socket:     socket,
		configPath: config,
	}
}

func (d *Daemon) Listen() error {
	c, err := d.loadConfig()
	if err != nil {
		return err
	}

	d.config = &c

	defer func() {
		log.Println("removing socket")
		os.RemoveAll(d.socket)
	}()

	if err := os.MkdirAll(filepath.Dir(d.socket), 0755); err != nil {
		return err
	}

	l, err := net.Listen("unix", d.socket)
	if err != nil {
		return err
	}
	defer l.Close()

	d.server = grpc.NewServer()
	proto.RegisterDaemonServer(d.server, NewServer(d))

	log.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-quit

		log.Println("quitting...")

		d.server.Stop()
	}()

	go func() {
		if err := d.launchServer(); err != nil {
			spew.Dump(err)
		}
	}()

	return d.server.Serve(l)
}

func (d *Daemon) launchServer() error {
	d.m.Lock()
	defer d.m.Unlock()

	confs, err := ioutil.ReadDir("conf.d")
	if err != nil {
		return err
	}

	for _, info := range confs {
		if strings.HasSuffix(info.Name(), ".toml") {
			var conf doom.Config
			if _, err := toml.DecodeFile(filepath.Join("conf.d", info.Name()), &conf); err != nil {
				log.Println(err)
				continue
			}

			d.instances = append(d.instances, doom.NewInstance(conf))
		}
	}

	for _, instance := range d.instances {
		go func(i *doom.Instance) {
			if err := i.Start(); err != nil {
				log.Println(err)
				return
			}
		}(instance)
	}

	return nil
}

func (d *Daemon) loadConfig() (Config, error) {
	candidates := []string{"config.toml"}
	if d.configPath != nil {
		candidates = append(candidates, *d.configPath)
	}

	for _, c := range candidates {
		fi, err := os.Stat(c)
		if os.IsNotExist(err) {
			continue
		}

		c := defaultConfig

		if _, err := toml.DecodeFile(fi.Name(), &c); err != nil {
			return Config{}, err
		}

		return c, nil
	}

	return defaultConfig, nil
}
