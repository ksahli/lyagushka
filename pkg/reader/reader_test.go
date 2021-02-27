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
	t.Run("send en errors", func(t *testing.T) {
		errs := make(chan error)
		reader := reader.New("no/path")

		go reader.Run(errs)

		err := <-errs
		if err == nil {
			t.Fatal("want an error, got nothing")
		}
	})

	t.Run("send ressources", func(t *testing.T) {
		path := t.TempDir()

		want := make(map[string]core.Ressource)
		for i := 0; i < 10; i++ {
			name := fmt.Sprintf("name-%d", i)
			text := fmt.Sprintf("text-%d", i)

			fullpath := filepath.Join(path, name)
			content := []byte(text)

			if err := ioutil.WriteFile(fullpath, content, 0664); err != nil {
				t.Fatal(err)
			}

			want[name] = core.Ressource{
				Path:    name,
				Content: content,
			}
		}

		reader := reader.New(path)
		errs := make(chan error, 10)

		go reader.Run(errs)

		got := make(map[string]core.Ressource)
		for ressource := range reader.Ressources() {
			got[ressource.Path] = ressource
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want %v, got %v", want, got)
		}

		close(errs)
		for err := range errs {
			t.Fatal(err)
		}
	})
}
