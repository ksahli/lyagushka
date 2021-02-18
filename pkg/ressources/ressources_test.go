package ressources_test

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ksahli/lyagushka/pkg/core"
	"github.com/ksahli/lyagushka/pkg/ressources"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	t.Run("read from non existing path", func(t *testing.T) {
		if _, err := ressources.Read("non_existing/ressource"); err == nil {
			t.Fatal("want an error, got nothing")
		}
	})

	t.Run("read ressource", func(t *testing.T) {
		want := setup(dir, "content")
		save(t, want.Path, want.Content)

		got, err := ressources.Read(want.Path)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want %v, got %v", want, got)
		}
	})
}

func TestWrite(t *testing.T) {
	dir := t.TempDir()

	t.Run("write to non existing path", func(t *testing.T) {
		ressource := setup("non_existing/directory", "content")
		if err := ressources.Write(ressource); err == nil {
			t.Fatal("want an error, got nothing")
		}
	})

	t.Run("write ressource", func(t *testing.T) {
		want := setup(dir, "content")

		if err := ressources.Write(want); err != nil {
			t.Fatal(err)
		}

		got := load(t, want.Path)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("want %v, got %v", want, got)
		}
	})
}

func setup(dir, content string) core.Ressource {
	path := filepath.Join(dir, "ressource")
	return core.Ressource{
		Path:    path,
		Content: []byte(content),
	}
}

func save(t *testing.T, path string, content []byte) {
	if err := ioutil.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
}

func load(t *testing.T, path string) core.Ressource {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return core.Ressource{
		Path:    path,
		Content: content,
	}
}
