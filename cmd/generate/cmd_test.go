package generate_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ksahli/lyagushka/cmd/generate"
)

func TestGenerate(t *testing.T) {
	path := t.TempDir()

	want := make(map[string][]byte)
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("name-%d", i)
		content := fmt.Sprintf("base content %d", i)
		want[name] = []byte(content)
	}
	setup(t, path, 10)

	err := generate.Run(path)
	if err != nil {
		t.Fatal(err)
	}

	got := make(map[string][]byte)
	dstPath := filepath.Join(path, "dst")
	filepath.Walk(dstPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			got[info.Name()] = content
		}
		return nil
	})

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func setup(t *testing.T, dir string, number int) {
	tplDir := filepath.Join(dir, "tpl")
	if err := os.Mkdir(tplDir, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	baseTpl := `{{define "base"}}base {{block "main" .}}{{end}}{{end}}`
	baseTplPath := filepath.Join(tplDir, "base.html")
	if err := ioutil.WriteFile(baseTplPath, []byte(baseTpl), 0664); err != nil {
		t.Fatal(err)
	}

	srcDir := filepath.Join(dir, "src")
	if err := os.Mkdir(srcDir, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("name-%d", i)
		content := fmt.Sprintf(`{{define "main"}}content %d{{end}}`, i)
		path := filepath.Join(srcDir, name)

		if err := ioutil.WriteFile(path, []byte(content), 0664); err != nil {
			t.Fatal(err)
		}
	}

	dstDir := filepath.Join(dir, "dst")
	if err := os.Mkdir(dstDir, os.ModePerm); err != nil {
		t.Fatal(err)
	}
}
