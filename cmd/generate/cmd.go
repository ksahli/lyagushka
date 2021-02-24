package generate

import (
	"path/filepath"

	"github.com/ksahli/lyagushka/pkg/processor"
	"github.com/ksahli/lyagushka/pkg/reader"
	"github.com/ksahli/lyagushka/pkg/templates"
	"github.com/ksahli/lyagushka/pkg/writer"
)

func Run(path string) error {
	errs := make(chan error, 10)

	tplDir := filepath.Join(path, "tpl")
	templates, err := templates.Load(tplDir)
	if err != nil {
		return err
	}

	srcDir := filepath.Join(path, "src")
	reader := reader.New(srcDir)

	processor := processor.New(reader.Ressources(), templates)

	dstDir := filepath.Join(path, "dst")
	writer := writer.New(dstDir, processor.Ressources())

	go writer.Run(errs)
	go processor.Run(errs)
	go reader.Run(errs)

	<-writer.Done()

	close(errs)
	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
