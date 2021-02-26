package clean

import (
	"os"
	"path/filepath"
)

func Run(path string) error {
	dst := filepath.Join(path, "dst")
	if err := os.RemoveAll(dst); err != nil {
		return err
	}
	return nil
}
