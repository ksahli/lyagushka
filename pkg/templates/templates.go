package templates

import (
	"io/ioutil"
	"path/filepath"

	"github.com/ksahli/lyagushka/pkg/core"
)

func Load(path string) (core.Templates, error) {
	templates := make(core.Templates)
	base, err := load(path, "base.html")
	if err != nil {
		return templates, err
	}
	templates["base"] = base
	return templates, nil
}

func load(path, name string) (string, error) {
	p := filepath.Join(path, name)
	content, err := ioutil.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
