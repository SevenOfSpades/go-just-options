package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	t.Run("it should read value from options", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedValue := 42
		optionKey := OptionKeyFromString("test-option")
		opt := New().Resolve(setOptionForTest(optionKey, expectedValue))

		// WHEN
		result, err := Read[int](opt, optionKey)

		// THEN
		require.NoError(t, err)
		assert.Equal(t, expectedValue, result)
	})
	t.Run("it should return an error when option is not present in option set", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		optionKey := OptionKeyFromString("test-option")
		opt := New().Resolve()

		// WHEN
		result, err := Read[int](opt, optionKey)

		// THEN
		require.ErrorIs(t, err, ErrNotFound)
		assert.Equal(t, 0, result)
	})
}

func TestReadOrPanic(t *testing.T) {
	t.Run("it should panic when there is an error during read operation", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		optionKey := OptionKeyFromString("test-option")
		opt := New().Resolve()

		// WHEN-THEN
		assert.PanicsWithError(t, "cannot read option 'test-option' from provided option set: option not found", func() {
			_ = ReadOrPanic[int](opt, optionKey)
		})
	})
}

func TestReadOrDefaultOrPanic(t *testing.T) {
	t.Run("it should return default value for non-existing option in set", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedDefaultValue := 100
		optionKey := OptionKeyFromString("test-option")
		opt := New().Resolve()

		// WHEN
		result := ReadOrDefaultOrPanic[int](opt, optionKey, expectedDefaultValue)

		// THEN
		assert.Equal(t, expectedDefaultValue, result)
	})
	t.Run("it should panic on different errors than not found", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		optionKey := OptionKeyFromString("test-option")
		opt := New().Resolve(setOptionForTest(optionKey, "100"))

		// WHEN-THEN
		assert.PanicsWithError(t, "option 'test-option' is expected to be int but is string: wrong type expected from option", func() {
			_ = ReadOrDefaultOrPanic[int](opt, optionKey, 200)
		})
	})
}

func setOptionForTest[T any](key OptionKey, value T) Option {
	return func(options Options) {
		WriteOrPanic[T](options, key, value)
	}
}
