package options

// Option should be provided by wrapper function as value setter.
// It will be called when passed to Options.Resolve.
type Option func(Options)
