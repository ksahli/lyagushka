package writer_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ksahli/lyagushka/pkg/core"
	"github.com/ksahli/lyagushka/pkg/ressources"
	"github.com/ksahli/lyagushka/pkg/writer"
)

func TestWrite(t *testing.T) {
	t.Run("test write to non existing path", func(t *testing.T) {
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

	t.Run("test write ressources", func(t *testing.T) {
		dir := t.TempDir()

		in, errs := make(chan core.Ressource, 10), make(chan error, 1)
		writer := writer.New(dir, in)

		go writer.Run(errs)
		want := setup(dir, 10)
		for _, ressource := range want {
			emit(t, dir, ressource, in)
		}

		close(in)
		<-writer.Done()

		got := read(t, dir)

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want %v \n got %v", want, got)
		}

		close(errs)
		err := <-errs
		if err != nil {
			t.Fatal(err)
		}
	})
}

func setup(dir string, number int) map[string]core.Ressource {
	ressources := map[string]core.Ressource{}
	for i := 0; i < 10; i++ {
		name, content := fmt.Sprintf("name-%d", i), fmt.Sprintf("content %d", i)
		path := filepath.Join(dir, name)
		ressources[name] = core.Ressource{
			Path:    path,
			Content: []byte(content),
		}
	}
	return ressources
}

func emit(t *testing.T, dir string, ressource core.Ressource, in chan core.Ressource) {
	path, err := filepath.Rel(dir, ressource.Path)
	if err != nil {
		t.Fatal(err)
	}
	ressource.Path = path
	in <- ressource
}

func read(t *testing.T, dir string) map[string]core.Ressource {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	rs := map[string]core.Ressource{}
	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		ressource, err := ressources.Read(path)
		if err != nil {
			t.Fatal(err)
		}
		rs[file.Name()] = ressource
	}
	return rs
}
