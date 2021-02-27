package ressources_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ksahli/lyagushka/pkg/core"
	"github.com/ksahli/lyagushka/pkg/ressources"
)

var (
	path    = "ressource"
	content = []byte("content")
)

func TestRead(t *testing.T) {
	t.Run("returns an error", func(t *testing.T) {
		readFunc := func(path string) ([]byte, error) {
			return nil, errors.New("fail")
		}

		reader := ressources.Reader{
			ReadFunc: readFunc,
		}

		if _, err := reader.Read("something"); err == nil {
			t.Fatal("want an error, got nothing")
		}
	})

	t.Run("reads a ressource", func(t *testing.T) {
		want := core.Ressource{
			Path:    path,
			Content: content,
		}

		readFunc := func(path string) ([]byte, error) {
			return content, nil
		}

		reader := ressources.Reader{
			ReadFunc: readFunc,
		}

		got, err := reader.Read(path)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want %v, got %v", want, got)
		}
	})

	t.Run("reads a ressource from filesystem", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, path)
		want := core.Ressource{
			Path:    path,
			Content: content,
		}

		if err := ioutil.WriteFile(path, content, 0664); err != nil {
			t.Fatal(err)
		}

		reader := ressources.Reader{
			ReadFunc: ioutil.ReadFile,
		}

		got, err := reader.Read(path)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want %v, got %v", want, got)
		}
	})
}

func TestWrite(t *testing.T) {
	t.Run("returns an error", func(t *testing.T) {
		writerFunc := func(string, []byte, os.FileMode) error {
			return errors.New("fail")
		}

		writer := ressources.Writer{
			WriteFunc: writerFunc,
		}

		ressource := core.Ressource{
			Path:    path,
			Content: content,
		}

		if err := writer.Write(ressource); err == nil {
			t.Fatal("want an error, got nothing")
		}
	})

	t.Run("writes a ressource", func(t *testing.T) {
		want := core.Ressource{
			Path:    path,
			Content: content,
		}

		writerFunc := func(string, []byte, os.FileMode) error {
			return nil
		}

		writer := ressources.Writer{
			WriteFunc: writerFunc,
		}

		if err := writer.Write(want); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("writes a ressource to filesystem", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, path)
		want := core.Ressource{
			Path:    path,
			Content: content,
		}

		writer := ressources.Writer{
			WriteFunc: ioutil.WriteFile,
		}

		if err := writer.Write(want); err != nil {
			t.Fatal(err)
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		got := core.Ressource{
			Path:    path,
			Content: content,
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want %v, got %v", want, got)
		}
	})
}
