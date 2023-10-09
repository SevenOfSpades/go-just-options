package options

import "errors"

var (
	ErrNotFound      = errors.New("option not found")
	ErrTypeMismatch  = errors.New("wrong type expected from option")
	ErrDuplicatedKey = errors.New("option has already been set")
)

type Options interface {
	Resolve(...Option) Options
	getRegistry() *registry
}
