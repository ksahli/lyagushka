package core_test

import (
	"bytes"
	"testing"

	"github.com/ksahli/lyagushka/pkg/core"
)

func TestExecute(t *testing.T) {
	t.Run("execute without base template", func(t *testing.T) {
		templates := core.Templates{}
		content := []byte(`{{define "main"}}content{{end}}`)
		_, err := templates.Execute(content)
		if err == nil {
			t.Fatal("want an error, got nothing")
		}
	})

	t.Run("execute template", func(t *testing.T) {
		want := []byte("base content")
		templates := core.Templates{
			"base": `{{define "base"}}base {{ block "main" .}}{{end}}{{ end }}`,
		}

		content := []byte(`{{define "main"}}content{{end}}`)
		got, err := templates.Execute(content)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(want, got) {
			t.Fatalf("want %s, got %s", want, got)
		}
	})
}
