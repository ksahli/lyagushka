package reader

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ksahli/lyagushka/pkg/core"
	"github.com/ksahli/lyagushka/pkg/ressources"
)

type Reader struct {
	path       string
	ressources chan core.Ressource
}

func New(path string) Reader {
	return Reader{
		path:       path,
		ressources: make(chan core.Ressource, 10),
	}
}

func (r Reader) Ressources() chan core.Ressource {
	return r.ressources
}

func (r Reader) Run(errs chan error) {
	defer close(r.ressources)

	if err := filepath.Walk(r.path, r.walk); err != nil {
		errs <- fmt.Errorf("reader error: %w", err)
	}
}

func (r Reader) walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	} else {
		ressource, err := r.read(path)
		if err != nil {
			return err
		}
		r.ressources <- ressource
		return nil
	}
}

func (r Reader) read(path string) (core.Ressource, error) {
	reader := ressources.Reader{
		ReadFunc: ioutil.ReadFile,
	}

	ressource, err := reader.Read(path)
	if err != nil {
		return core.Ressource{}, err
	}

	if err := ressource.Rel(r.path); err != nil {
		return core.Ressource{}, err
	}

	return ressource, nil
}
