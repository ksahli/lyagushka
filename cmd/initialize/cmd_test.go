package initialize_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ksahli/lyagushka/cmd/initialize"
)

func TestRun(t *testing.T) {
	t.Run("initialize directories in non existing path", func(t *testing.T) {
		if err := initialize.Run("non/existing/path"); err == nil {
			t.Fatal("want an error, got nothing")
		}
	})

	t.Run("initialize directories", func(t *testing.T) {
		dir := t.TempDir()
		if err := initialize.Run(dir); err != nil {
			t.Fatal(err)
		}

		for _, name := range []string{"src", "dst", "tpl"} {
			info := read(t, dir, name)
			if !info.IsDir() {
				t.Fatalf("want %s directory, got nothing", name)
			}
		}

		homePath := filepath.Join(dir, "src", "home.html")
		if _, err := os.Open(homePath); err != nil {
			t.Fatal(err)
		}

		tplBasePath := filepath.Join(dir, "tpl", "base.html" )
		if _, err := os.Open(tplBasePath); err != nil {
			t.Fatal(err)
		}
	})
}

func read(t *testing.T, dir, name string) os.FileInfo {
	path := filepath.Join(dir, name)
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	info, err := file.Stat()
	if err != nil {
		t.Fatal(err)
	}
	return info
}
