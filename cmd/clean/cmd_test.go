package clean_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ksahli/lyagushka/cmd/clean"
)

func TestRun(t *testing.T) {
	dir := t.TempDir()
	dstPath := filepath.Join(dir, "dst")
	if err := os.Mkdir(dstPath, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	if err := clean.Run(dir); err != nil {
		t.Fatal(err)
	}

	_, err := os.Open(dstPath)
	if err == nil {
		t.Fatal("expecting an error, got nothing")
	}
}
