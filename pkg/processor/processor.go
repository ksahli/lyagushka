package processor

import (
	"github.com/ksahli/lyagushka/pkg/core"
)

type Transformer interface {
	Transform(core.Ressource) (core.Ressource, error)
}

type Processor struct {
	in        chan core.Ressource
	out       chan core.Ressource
	templates core.Templates
}

func New(in chan core.Ressource, templates core.Templates) Processor {
	return Processor{
		in:        in,
		out:       make(chan core.Ressource, 10),
		templates: templates,
	}
}

func (p Processor) Ressources() chan core.Ressource {
	return p.out
}

func (p Processor) Run(errs chan error) {
	defer close(p.out)
	for ressource := range p.in {
		content, err := p.templates.Execute(ressource.Content)
		if err != nil {
			errs <- err
		} else {
			p.out <- core.Ressource{
				Path:    ressource.Path,
				Content: content,
			}
		}
	}
}
