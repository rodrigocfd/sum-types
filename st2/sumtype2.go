package st2

// A sum type which can have 2 different types.
type SumType2[T0, T1 any] struct {
	which uint8
	val0  T0
	val1  T1
}

// Creates a new SumType2 containing a value of type T0.
func New0[T0, T1 any](value T0) SumType2[T0, T1] {
	return SumType2[T0, T1]{
		which: 0,
		val0:  value,
	}
}

// Creates a new SumType2 containing a value of type T1.
func New1[T0, T1 any](value T1) SumType2[T0, T1] {
	return SumType2[T0, T1]{
		which: 1,
		val1:  value,
	}
}

// Returns the index of the currently stored type.
func (st *SumType2[T0, T1]) Which() uint8 {
	return st.which
}

// Tells whether the stored value is of type T0.
func (st *SumType2[T0, T1]) Has0() bool {
	return st.which == 0
}

// Tells whether the stored value is of type T1.
func (st *SumType2[T0, T1]) Has1() bool {
	return st.which == 1
}

// If the stored value is of type T0, returns it and true; otherwise, returns a
// default value and false.
func (st *SumType2[T0, T1]) Get0() (T0, bool) {
	if st.which != 0 {
		var def0 T0
		return def0, false
	}
	return st.val0, true
}

// If the stored value is of type T1, returns it and true; otherwise, returns a
// default value and false.
func (st *SumType2[T0, T1]) Get1() (T1, bool) {
	if st.which != 1 {
		var def1 T1
		return def1, false
	}
	return st.val1, true
}

// Calls the function according to the stored value; the other functions won't
// be called. This is equivalent to an exhaustive pattern matching.
func (st *SumType2[T0, T1]) Match(
	case0 func(value T0),
	case1 func(value T1)) {

	switch st.which {
	case 0:
		case0(st.val0)
	case 1:
		case1(st.val1)
	}
}
