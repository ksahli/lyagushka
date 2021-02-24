package reader_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ksahli/lyagushka/pkg/core"
	"github.com/ksahli/lyagushka/pkg/reader"
)

func TestRead(t *testing.T) {
	t.Run("read ressources from non exisiting path", func(t *testing.T) {
		errs := make(chan error)
		reader := reader.New("non/existing/path")

		go reader.Run(errs)

		err := <-errs
		if err == nil {
			t.Fatal("want an error, got nothing")
		}
	})

	t.Run("read ressources", func(t *testing.T) {
		path := t.TempDir()
		errs := make(chan error, 10)
		expected := setup(path, 10)
		write(t, path, expected)

		reader := reader.New(path)
		go reader.Run(errs)

		result := map[string]core.Ressource{}
		for ressource := range reader.Ressources() {
			result[ressource.Path] = ressource
		}

		if !reflect.DeepEqual(expected, result) {
			t.Fatalf("want %v, got %v", expected, result)
		}

		close(errs)
		for err := range errs {
			t.Fatal(err)
		}
	})
}

func setup(dir string, number int) map[string]core.Ressource {
	ressources := map[string]core.Ressource{}
	for i := 0; i <= 10; i++ {
		path := fmt.Sprintf("name-%d", i)
		content := fmt.Sprintf("content-%d", i)
		ressources[path] = core.Ressource{
			Path:    path,
			Content: []byte(content),
		}
	}
	return ressources
}

func write(t *testing.T, path string, rs map[string]core.Ressource) {
	for _, ressource := range rs {
		p := filepath.Join(path, ressource.Path)
		if err := ioutil.WriteFile(p, []byte(ressource.Content), 0664); err != nil {
			t.Fatal(err)
		}
	}
}
