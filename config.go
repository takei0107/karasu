package karasu

import (
	"io"

	"github.com/pelletier/go-toml/v2"
)

type config struct {
	Document document_config
	Server   server_config
}

type document_config struct {
	Root string
}

type server_config struct {
	Port int
}

func load_config(cfg_file io.Reader) (*config, error) {
	b, err := io.ReadAll(cfg_file)
	if err != nil {
		return nil, err
	}
	var conf config
	perr := toml.Unmarshal(b, &conf)
	if perr != nil {
		return nil, perr
	}
	return &conf, nil
}
