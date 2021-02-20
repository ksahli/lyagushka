package writer

import (
	"path/filepath"

	"github.com/ksahli/lyagushka/pkg/core"
	"github.com/ksahli/lyagushka/pkg/ressources"
)

type Writer struct {
	path string
	in   chan core.Ressource
	done chan bool
}

func New(path string, in chan core.Ressource) Writer {
	return Writer{
		path: path,
		in:   in,
		done: make(chan bool),
	}
}

func (w Writer) Run(errs chan error) {
	defer close(w.done)
	for ressource := range w.in {
		ressource.Path = filepath.Join(w.path, ressource.Path)
		if err := ressources.Write(ressource); err != nil {
			errs <- err
		}
	}
	w.done <- true
}

func (w Writer) Done() chan bool {
	return w.done
}
