package options

type (
	// Options is collection of Option value setters.
	Options []Option

	// Option should be provided by wrapper function as value setter.
	// It will be called when passed to Resolver.Resolve.
	Option func(Resolver)
)
