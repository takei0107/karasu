package karasu

import (
	"io"

	"github.com/pelletier/go-toml/v2"
)

var (
	port_min = 0
	port_max = 49151
)

type config struct {
	Document documentConfig
	Server   serverConfig
}

type documentConfig struct {
	Root     string
	Location []locationConfig
}

type locationConfig struct {
	Pattern string
	Path    string
}

type serverConfig struct {
	Port int
}

func load_config(cfg_file io.Reader) (*config, error) {
	b, err := io.ReadAll(cfg_file)
	if err != nil {
		return nil, err
	}
	var conf config
	// parse toml
	err = toml.Unmarshal(b, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
