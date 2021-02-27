package ressources

import (
	"os"

	"github.com/ksahli/lyagushka/pkg/core"
)

type ReadFunc func(string) ([]byte, error)

type Reader struct {
	ReadFunc
}

func (r Reader) Read(path string) (core.Ressource, error) {
	content, err := r.ReadFunc(path)
	if err != nil {
		return core.Ressource{}, err
	}
	ressource := core.Ressource{
		Path:    path,
		Content: content,
	}
	return ressource, nil
}

type WriteFunc func(string, []byte, os.FileMode) error

type Writer struct {
	WriteFunc
}

func (w Writer) Write(ressource core.Ressource) error {
	path, content := ressource.Path, ressource.Content
	return w.WriteFunc(path, content, 0664)
}
