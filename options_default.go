package options

type defaultResolver struct {
	registry *registry
}

func newResolver() *defaultResolver {
	return &defaultResolver{registry: newRegistry()}
}

func (o *defaultResolver) Resolve(setters Options) Resolver {
	for _, x := range setters {
		x(o)
	}
	return o
}

func (o *defaultResolver) getRegistry() *registry {
	return o.registry
}
