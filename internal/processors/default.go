package processors

type Processor interface {
	Label() string
	Execute(interface{})
	Fn(func(interface{}))
}

type DefaultProcessor struct {
	Name    string
	Handler func(interface{})
}

func (p *DefaultProcessor) Label() string {
	return p.Name
}
func (p *DefaultProcessor) Fn(handler func(interface{})) {
	if handler != nil {
		p.Handler = handler
	}
}
func (p *DefaultProcessor) Execute(args interface{}) {
	p.Handler(args)
}
