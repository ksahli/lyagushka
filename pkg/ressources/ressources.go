package ressources

import (
	"io/ioutil"

	"github.com/ksahli/lyagushka/pkg/core"
)

// Reads a ressource from file system
func Read(path string) (core.Ressource, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return core.Ressource{}, err
	}
	ressource := core.Ressource{
		Path:    path,
		Content: content,
	}
	return ressource, nil
}

// Writes a ressource to file system
func Write(ressource core.Ressource) error {
	path, content := ressource.Path, ressource.Content
	return ioutil.WriteFile(path, content, 0664)
}
