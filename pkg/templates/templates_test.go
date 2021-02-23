package templates_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/ksahli/lyagushka/pkg/templates"
)

func TestLoad(t *testing.T) {
	t.Run("load templates from non existing path", func(t *testing.T) {
		if _, err := templates.Load("non/existing/path"); err == nil {
			t.Fatal("want an error, got nothing")
		}
	})

	t.Run("load templates without an undefined bqse template", func(t *testing.T) {
		path := t.TempDir()
		if _, err := templates.Load(path); err == nil {
			t.Fatal("want an error, got noting")
		}
	})

	t.Run("load templates", func(t *testing.T) {
		path := t.TempDir()
		want := "content"
		setup(t, path, want)

		templates, err := templates.Load(path)
		if err != nil {
			t.Fatal(err)
		}

		if got := templates["base"]; want != got {
			t.Fatalf("want %s, got %s", want, got)
		}
	})
}

func setup(t *testing.T, dir, content string) {
	path := filepath.Join(dir, "base.html")
	b := []byte(content)
	if err := ioutil.WriteFile(path, b, 0664); err != nil {
		t.Fatal(err)
	}
}
