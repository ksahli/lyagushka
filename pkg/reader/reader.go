package reader

import (
	"log"
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
	log.Printf("reading %s ...", r.path)
	if err := filepath.Walk(r.path, r.read); err != nil {
		errs <- err
	}
	log.Printf("reading complete")
}

func (r Reader) read(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	ressource, err := ressources.Read(path)
	if err != nil {
		return err
	}
	r.ressources <- ressource
	return nil
}
