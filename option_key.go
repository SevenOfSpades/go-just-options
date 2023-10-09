package options

// OptionKey is unique identifier of every option in current set.
type OptionKey string

func OptionKeyFromString[T ~string](val T) OptionKey {
	return OptionKey(val)
}

func (k OptionKey) String() string {
	return string(k)
}
