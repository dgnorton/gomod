package gomod

import (
	"fmt"
	"io"
	"os"

	"github.com/pelletier/go-toml"
)

type DepLock struct {
	Projects []*Project `toml:"projects"`
}

type Project struct {
	Digest    string   `toml:"digest"`
	Name      string   `toml:"name"`
	Packages  []string `toml:"packages"`
	PruneOpts string   `toml:"pruneopts"`
	Revision  string   `toml:"revision"`
	Version   string   `toml:"version",omitempty`
}

func NewDepLock() *DepLock {
	return &DepLock{}
}

func LoadDepLock(path string) (*DepLock, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadDepLock(f, path)
}

func ReadDepLock(r io.Reader, filename string) (*DepLock, error) {
	depLock := NewDepLock()
	if err := toml.NewDecoder(r).Decode(depLock); err != nil {
		return nil, fmt.Errorf("%s: %s", filename, err)
	}
	return depLock, nil
}
