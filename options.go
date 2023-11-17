package options

import (
	"errors"
	"fmt"
)

// Resolve will create Options instance filled with values based on received Option modificators.
func Resolve(opts ...Option) Options {
	return newOptions().Resolve(opts...)
}

// Read will attempt to acquire value associated with provided key from received Options set.
// If operation fails due to value being missing it will respond with ErrNotFound.
// If operation fails due to mismatch type between provided value and stored value then it will respond with ErrTypeMismatch.
func Read[T any](options Options, key OptionKey) (T, error) {
	var t T
	val, typeName, err := options.getRegistry().get(key)
	if err != nil {
		return t, fmt.Errorf("cannot read option '%s' from provided option set: %w", key.String(), err)
	}
	expType := fmt.Sprintf("%T", t)
	if expType == "<nil>" {
		if _, ok := val.(T); !ok {
			if typeName == "<nil>" {
				return t, fmt.Errorf("option '%s' is expected to be %s but got nil: %w", key.String(), fmt.Sprintf("%T", (*T)(nil))[1:], ErrNilValue)
			}
			return t, fmt.Errorf("option '%s' is expected to implement %s but %s is not compatible with it: %w", key.String(), fmt.Sprintf("%T", (*T)(nil))[1:], typeName, ErrTypeMismatch)
		}
	} else if expType != typeName {
		return t, fmt.Errorf("option '%s' is expected to be %s but is %s: %w", key.String(), expType, typeName, ErrTypeMismatch)
	}
	return val.(T), nil
}

// ReadOrDefault acts exactly the same as Read with one difference being the option for providing default value if
// there is none set in received Options set under provided key.
// Attempt to acquire value can still fail with ErrTypeMismatch is types do not match.
func ReadOrDefault[T any](options Options, key OptionKey, defaultValue T) (T, error) {
	val, err := Read[T](options, key)
	if err != nil {
		if (any(defaultValue) == nil && errors.Is(err, ErrNilValue)) || errors.Is(err, ErrNotFound) {
			return defaultValue, nil
		}

		var t T
		return t, err
	}
	return val, nil
}

// ReadOrPanic acts exactly the same as Read, but any error occurring during attempt to acquire value
// will result in panic instead of returning an error. This also means that number of return parameters
// is reduced to only one.
func ReadOrPanic[T any](options Options, key OptionKey) T {
	val, err := Read[T](options, key)
	if err != nil {
		panic(err)
	}
	return val
}

// ReadOrDefaultOrPanic acts exactly the same as ReadOrDefault, but any error occurring during attempt to acquire value
// (except for ErrNotFound) will result in panic instead of returning an error. This also means that number of return parameters
// is reduced to only one.
func ReadOrDefaultOrPanic[T any](options Options, key OptionKey, defaultValue T) T {
	val, err := ReadOrDefault[T](options, key, defaultValue)
	if err != nil {
		panic(err)
	}
	return val
}

// Write will attempt to store supplied value and associate it with provided key in received Options set.
// If operation fails due to key being already present within received Options set it will respond with ErrDuplicatedKey.
func Write[T any](options Options, key OptionKey, val T) error {
	if err := options.getRegistry().set(key, val); err != nil {
		return fmt.Errorf("cannot write option '%s' to provided option set: %w", key.String(), err)
	}
	return nil
}

// WriteOrPanic acts exactly the same as Write, but any error occurring during attempt to store value
// will result in panic instead of returning an error. This also means there are no longer any return parameters
// needed for this operation.
func WriteOrPanic[T any](options Options, key OptionKey, val T) {
	if err := Write[T](options, key, val); err != nil {
		panic(err)
	}
}
