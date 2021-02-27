package writer_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ksahli/lyagushka/pkg/core"
	"github.com/ksahli/lyagushka/pkg/ressources"
	"github.com/ksahli/lyagushka/pkg/writer"
)

func TestWrite(t *testing.T) {
	t.Run("send an error", func(t *testing.T) {
		in, errs := make(chan core.Ressource, 10), make(chan error, 1)
		writer := writer.New("non/exisiting/path", in)

		go writer.Run(errs)

		close(in)
		<-writer.Done()

		close(errs)
		for err := range errs {
			if err == nil {
				t.Fatal("want an error, got nothing")
			}
		}
	})

	t.Run("write ressources", func(t *testing.T) {
		path := t.TempDir()

		in, errs := make(chan core.Ressource, 10), make(chan error, 1)
		writer := writer.New(path, in)

		go writer.Run(errs)

		want := make(map[string]core.Ressource)
		for i := 0; i < 10; i++ {
			name := fmt.Sprintf("name-%d", i)
			text := fmt.Sprintf("text-%d", i)

			fullpath := filepath.Join(path, name)
			content := []byte(text)

			in <- core.Ressource{
				Path:    name,
				Content: content,
			}

			want[name] = core.Ressource{
				Path:    fullpath,
				Content: content,
			}
		}

		close(in)
		<-writer.Done()

		got := make(map[string]core.Ressource)
		filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
			if err != nil {
				t.Fatal(err)
			}
			if file.IsDir() {
				return nil
			}
			reader := ressources.Reader{
				ReadFunc: ioutil.ReadFile,
			}
			ressource, err := reader.Read(path)
			if err != nil {
				t.Fatal(err)
			}
			got[file.Name()] = ressource
			return nil
		})

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want %v \n got %v", want, got)
		}

		close(errs)
		if err := <-errs; err != nil {
			t.Fatal(err)
		}
	})
}
