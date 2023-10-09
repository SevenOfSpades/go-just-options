package options

type defaultOptions struct {
	registry *registry
}

func newOptions() *defaultOptions {
	return &defaultOptions{registry: newRegistry()}
}

func (o *defaultOptions) Resolve(setters ...Option) Options {
	for _, x := range setters {
		x(o)
	}
	return o
}

func (o *defaultOptions) getRegistry() *registry {
	return o.registry
}
