package core

import "path/filepath"

type Ressource struct {
	Path    string
	Content []byte
}

func (r *Ressource) Rel(path string) error {
	relpath, err := filepath.Rel(path, r.Path)
	if err != nil {
		return err
	}
	r.Path = relpath
	return nil
}
