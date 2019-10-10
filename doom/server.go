package doom

import (
	"fmt"
	"strconv"
)

type Mode string

const (
	DM  Mode = "dm"
	CTF      = "ctf"
	TDM      = "tdm"
)

type args []string

func (a args) Add(key, value string) args {
	a = append(a, key, value)

	return a
}

type Config struct {
	Name          string   `toml:"name"`
	Hostname      string   `toml:"hostname"`
	Mode          Mode     `toml:"mode"`
	Port          int16    `toml:"port"`
	WADs          []string `toml:"wads"`
	DMFlags       uint64   `toml:"dmflags"`
	DMFlags2      uint64   `toml:"dmflags2"`
	ZADmflags     uint64   `toml:"zadmflags"`
	Compatflags   uint64   `toml:"compatflags"`
	Compatflags2  uint64   `toml:"compatflags2"`
	ZACompatFlags uint64   `toml:"zacompatflags"`
	Disabled      bool     `toml:"disabled"`
}

func (c Config) Args() []string {
	args := make(args, 0, 20).
		Add("+sv_hostname", c.Hostname).
		Add("-port", strconv.Itoa(int(c.Port))).
		Add("+skill", "4").
		Add("+sv_maxplayers", "16").
		Add("+dmflags", fmt.Sprintf("%v", c.DMFlags)).
		Add("+dmflags2", fmt.Sprintf("%v", c.DMFlags2)).
		Add("+zadmflags", fmt.Sprintf("%v", c.ZADmflags)).
		Add("+compatflags", fmt.Sprintf("%v", c.Compatflags)).
		Add("+compatflags2", fmt.Sprintf("%v", c.Compatflags2)).
		Add("+zacompatflags", fmt.Sprintf("%v", c.ZADmflags)).
		Add("+sv_updatemaster", "false").
		Add(fmt.Sprintf("+%s", c.Mode), "1")

	for _, wad := range c.WADs {
		args = args.Add("-file", wad)
	}

	return args
}
