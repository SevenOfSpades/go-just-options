package options

import "errors"

var (
	ErrNotFound      = errors.New("option not found")
	ErrTypeMismatch  = errors.New("wrong type expected from option")
	ErrNilValue      = errors.New("nil value")
	ErrDuplicatedKey = errors.New("option has already been set")
)

type Resolver interface {
	Resolve(Options) Resolver
	getRegistry() *registry
}
