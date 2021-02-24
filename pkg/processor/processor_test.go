package processor_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ksahli/lyagushka/pkg/core"
	"github.com/ksahli/lyagushka/pkg/processor"
)

var base = `{{define "base"}}base {{block "main" .}}{{end}}{{end}}`

func TestProcessor(t *testing.T) {
	t.Run("process ressources", func(t *testing.T) {
		in, errs := make(chan core.Ressource), make(chan error)
		templates := core.Templates{"base": base}
		processor := processor.New(in, templates)

		go processor.Run(errs)

		want := setup(10)
		emit(in, 10)

		close(in)
		close(errs)

		got := make(map[string]core.Ressource)
		for ressource := range processor.Ressources() {
			got[ressource.Path] = ressource
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want %s, got %s", want, got)
		}

		for err := range errs {
			if err != nil {
				t.Fatal(err)
			}
		}
	})
}

func setup(number int) map[string]core.Ressource {
	rs := make(map[string]core.Ressource)
	for i := 0; i < number; i++ {
		path := fmt.Sprintf("path-%d", i)
		content := fmt.Sprintf("base content %d", i)
		rs[path] = core.Ressource{
			Path:    path,
			Content: []byte(content),
		}
	}
	return rs
}

func emit(in chan core.Ressource, number int) {
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("path-%d", i)
		content := fmt.Sprintf(`{{define "main"}}content %d{{end}}`, i)
		in <- core.Ressource{
			Path:    path,
			Content: []byte(content),
		}
	}
}
