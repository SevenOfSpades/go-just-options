package options

import "fmt"

type registry struct {
	values   map[OptionKey]any
	typesMap map[OptionKey]string
}

func newRegistry() *registry {
	return &registry{values: make(map[OptionKey]any), typesMap: make(map[OptionKey]string)}
}

func (r *registry) set(key OptionKey, val any) error {
	if _, ok := r.values[key]; ok {
		return ErrDuplicatedKey
	}

	r.values[key] = val
	r.typesMap[key] = fmt.Sprintf("%T", val)

	return nil
}

func (r *registry) get(key OptionKey) (any, string, error) {
	if val, ok := r.values[key]; ok {
		return val, r.typesMap[key], nil
	}
	return nil, "", ErrNotFound
}
