package gomod

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/mod/module"
)

// emptyGoModHash is the hash of a 1-file tree containing a 0-length go.mod.
// A bug caused us to write these into go.sum files for non-modules.
// We detect and remove them.
// https://github.com/golang/go/blob/go1.15.6/src/cmd/go/internal/modfetch/fetch.go#L422
const emptyGoModHash = "h1:G7mAYYxgmS0lVkHyy2hEOLQCFB0DlQFTMLWggykrydY="

type GoSum struct {
	Modules map[module.Version][]string
	Status  map[ModSum]ModSumStatus
}

func NewGoSum() *GoSum {
	return &GoSum{
		Modules: map[module.Version][]string{},
		Status:  map[ModSum]ModSumStatus{},
	}
}

// LoadGoSum loads a go.sum file.
func LoadGoSum(path string) (*GoSum, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadGoSum(f, path)
}

// ReadGoSum reads a go.sum file from an io.Reader. An optional
// filename argument can be provided for better error messages.
func ReadGoSum(r io.Reader, filename string) (*GoSum, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	lineno := 0
	gosum := NewGoSum()
	for len(data) > 0 {
		var line []byte
		lineno++
		i := bytes.IndexByte(data, '\n')
		if i < 0 {
			line, data = data, nil
		} else {
			line, data = data[:i], data[i+1:]
		}
		f := strings.Fields(string(line))
		if len(f) == 0 {
			// Skip blank lines.
			continue
		}
		if len(f) != 3 {
			return nil, fmt.Errorf("malformed go.sum:\n%s:%d: wrong number of fields %v", filename, lineno, len(f))
		}
		if f[2] == emptyGoModHash {
			// Old bug in Go modules so skip this line.
			continue
		}
		mod := module.Version{Path: f[0], Version: f[1]}
		gosum.Modules[mod] = append(gosum.Modules[mod], f[2])
	}

	return gosum, nil
}

type ModSum struct {
	Module module.Version
	Sum    string
}

type ModSumStatus struct {
	Used  bool
	Dirty bool
}
