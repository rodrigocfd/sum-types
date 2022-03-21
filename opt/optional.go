package opt

// A type which contains an optional value, which can be some or none.
type Optional[T any] struct {
	some bool
	val  T
}

// Creates a new Optional containing a value.
func Some[T any](value T) Optional[T] {
	return Optional[T]{
		some: true,
		val:  value,
	}
}

// Creates a new Optional containing no value.
func None[T any]() Optional[T] {
	return Optional[T]{}
}

// Tells whether there is a stored value.
func (o *Optional[T]) IsSome() bool {
	return o.some
}

// If there is value, return it and true; otherwise, returns a default value and
// false.
func (o *Optional[T]) Get() (T, bool) {
	if !o.some {
		var def T
		return def, false
	}
	return o.val, true
}
