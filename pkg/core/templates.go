package core

import (
	"bytes"
	"errors"
	"html/template"
)

type Templates map[string]string

func (t Templates) base() (*template.Template, error) {
	content, ok := t["base"]
	if !ok {
		return nil, errors.New("base template not defined")
	}
	return template.New("base").Parse(content)
}

func (t Templates) Execute(content string) ([]byte, error) {
	base, err := t.base()
	if err != nil {
		return nil, err
	}
	tpl, err := base.Parse(content)
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, nil); err != nil {
		return nil, err
	}
	return buffer.Bytes(), err
}
