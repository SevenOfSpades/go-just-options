package options

import (
	"errors"
	"fmt"
	"testing"
)

type (
	testInterface interface {
		doStuff() bool
	}
	testInterfaceImplementation struct {
	}
	testWithoutInterfaceImplementation struct {
	}
)

func (t *testInterfaceImplementation) doStuff() bool {
	return true
}

func TestRead(t *testing.T) {
	t.Run("it should read value from options", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedValue := 42
		optionKey := OptionKeyFromString("test-option")
		opt := Resolve(Options{setOptionForTest(optionKey, expectedValue)})

		// WHEN
		result, err := Read[int](opt, optionKey)

		// THEN
		if err != nil {
			t.Error(fmt.Sprintf("not expected to get an error { - %s - }", err.Error()))
			return
		}
		if expectedValue != result {
			t.Error(fmt.Sprintf("value %v[expected] is not equal to %v[actual]", expectedValue, result))
		}
	})
	t.Run("it should return an error when option is not present in option set", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		optionKey := OptionKeyFromString("test-option")
		opt := Resolve(Options{})

		// WHEN
		result, err := Read[int](opt, optionKey)

		// THEN
		if !errors.Is(err, ErrNotFound) {
			t.Error(fmt.Sprintf("no error or error is not options.ErrNotFound [%v]", err))
			return
		}
		if result != 0 {
			t.Error(fmt.Sprintf("result is expected to be equal to 0 [%v]", result))
		}
	})
	t.Run("it should read value for struct implementing interface", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedValue := &testInterfaceImplementation{}
		optionKey := OptionKeyFromString("test-with-interface")
		opt := Resolve(Options{setOptionForTest(optionKey, expectedValue)})

		// WHEN
		result, err := Read[testInterface](opt, optionKey)

		// THEN
		if err != nil {
			t.Error(fmt.Sprintf("not expected to get an error { - %s - }", err.Error()))
			return
		}
		if _, ok := result.(*testInterfaceImplementation); !ok {
			t.Error("result is not compatible with testInterface")
		}
	})
	t.Run("it should fail with proper error when interface is not compatible with value", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedValue := &testWithoutInterfaceImplementation{}
		optionKey := OptionKeyFromString("test-with-interface")
		opt := Resolve(Options{setOptionForTest(optionKey, expectedValue)})

		// WHEN
		_, err := Read[testInterface](opt, optionKey)

		// THEN
		if err == nil {
			t.Error("an error is expected")
			return
		}
		if err.Error() != "option 'test-with-interface' is expected to implement options.testInterface but *options.testWithoutInterfaceImplementation is not compatible with it: wrong type expected from option" {
			t.Error("error message is incorrect")
			return
		}
	})
	t.Run("it should not allow for interface to be nil", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		optionKey := OptionKeyFromString("test-with-nil-interface")
		opt := Resolve(Options{setOptionForInterface(optionKey, nil)})

		// WHEN
		_, err := Read[testInterface](opt, optionKey)

		// THEN
		if err == nil {
			t.Error("an error is expected")
			return
		}
		if err.Error() != "option 'test-with-nil-interface' is expected to be options.testInterface but got nil: nil value" {
			t.Error("error message is incorrect")
			return
		}
	})
}

func TestReadOrDefault(t *testing.T) {
	t.Run("it should allow for interface to be nil when nil is specified as default value", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		optionKey := OptionKeyFromString("test-with-nil-interface")
		opt := Resolve(Options{setOptionForInterface(optionKey, nil)})

		// WHEN
		val, err := ReadOrDefault[testInterface](opt, optionKey, nil)

		// THEN
		if err != nil {
			t.Error(fmt.Sprintf("not expected to get an error { - %s - }", err.Error()))
			return
		}
		if val != nil {
			t.Error(fmt.Sprintf("value should be nil but is %T", val))
			return
		}
	})
}

func TestReadOrPanic(t *testing.T) {
	t.Run("it should panic when there is an error during read operation", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		optionKey := OptionKeyFromString("test-option")
		opt := Resolve(Options{})

		// WHEN-THEN
		func() {
			defer func() {
				r := recover()
				if r == nil {
					t.Error("operation is expected to trigger panic")
					return
				}
				err, ok := r.(error)
				if !ok {
					t.Error("recover value is not an error")
					return
				}
				if err.Error() != "cannot read option 'test-option' from provided option set: option not found" {
					t.Error(fmt.Sprintf("recover error contains unexpected value { - %s - }", err.Error()))
				}
			}()

			_ = ReadOrPanic[int](opt, optionKey)
		}()
	})
}

func TestReadOrDefaultOrPanic(t *testing.T) {
	t.Run("it should return default value for non-existing option in set", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedDefaultValue := 100
		optionKey := OptionKeyFromString("test-option")
		opt := Resolve(Options{})

		// WHEN
		result := ReadOrDefaultOrPanic[int](opt, optionKey, expectedDefaultValue)

		// THEN
		if expectedDefaultValue != result {
			t.Error(fmt.Sprintf("value %v[expected] is not equal to %v[actual]", expectedDefaultValue, result))
		}
	})
	t.Run("it should panic on different errors than not found", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		optionKey := OptionKeyFromString("test-option")
		opt := Resolve(Options{setOptionForTest(optionKey, "100")})

		// WHEN-THEN
		func() {
			defer func() {
				r := recover()
				if r == nil {
					t.Error("operation is expected to trigger panic")
					return
				}
				err, ok := r.(error)
				if !ok {
					t.Error("recover value is not an error")
					return
				}
				if err.Error() != "option 'test-option' is expected to be int but is string: wrong type expected from option" {
					t.Error(fmt.Sprintf("recover error contains unexpected value { - %s - }", err.Error()))
				}
			}()

			_ = ReadOrDefaultOrPanic[int](opt, optionKey, 200)
		}()
	})
}

func setOptionForInterface(key OptionKey, val testInterface) Option {
	return func(r Resolver) {
		WriteOrPanic[testInterface](r, key, val)
	}
}

func setOptionForTest[T any](key OptionKey, value T) Option {
	return func(r Resolver) {
		WriteOrPanic[T](r, key, value)
	}
}
