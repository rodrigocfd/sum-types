package st3

// A sum type which can have 3 different types.
type SumType3[T0, T1, T2 any] struct {
	which uint8
	val0  T0
	val1  T1
	val2  T2
}

// Creates a new SumType3 containing a value of type T0.
func New0[T0, T1, T2 any](value T0) SumType3[T0, T1, T2] {
	return SumType3[T0, T1, T2]{
		which: 0,
		val0:  value,
	}
}

// Creates a new SumType3 containing a value of type T1.
func New1[T0, T1, T2 any](value T1) SumType3[T0, T1, T2] {
	return SumType3[T0, T1, T2]{
		which: 1,
		val1:  value,
	}
}

// Creates a new SumType3 containing a value of type T2.
func New2[T0, T1, T2 any](value T2) SumType3[T0, T1, T2] {
	return SumType3[T0, T1, T2]{
		which: 2,
		val2:  value,
	}
}

// Returns the index of the currently stored type.
func (st *SumType3[T0, T1, T2]) Which() uint8 {
	return st.which
}

// Tells whether the stored value is of type T0.
func (st *SumType3[T0, T1, T2]) Has0() bool {
	return st.which == 0
}

// Tells whether the stored value is of type T1.
func (st *SumType3[T0, T1, T2]) Has1() bool {
	return st.which == 1
}

// Tells whether the stored value is of type T2.
func (st *SumType3[T0, T1, T2]) Has2() bool {
	return st.which == 2
}

// If the stored value is of type T0, returns it and true; otherwise, returns a
// default value and false.
func (st *SumType3[T0, T1, T2]) Get0() (T0, bool) {
	if st.which != 0 {
		var def0 T0
		return def0, false
	}
	return st.val0, true
}

// If the stored value is of type T1, returns it and true; otherwise, returns a
// default value and false.
func (st *SumType3[T0, T1, T2]) Get1() (T1, bool) {
	if st.which != 1 {
		var def1 T1
		return def1, false
	}
	return st.val1, true
}

// If the stored value is of type T2, returns it and true; otherwise, returns a
// default value and false.
func (st *SumType3[T0, T1, T2]) Get2() (T2, bool) {
	if st.which != 2 {
		var def2 T2
		return def2, false
	}
	return st.val2, true
}

// Calls the function according to the stored value; the other functions won't
// be called. This is equivalent to an exhaustive pattern matching.
func (st *SumType3[T0, T1, T2]) Match(
	case0 func(value T0),
	case1 func(value T1),
	case2 func(value T2)) {

	switch st.which {
	case 0:
		case0(st.val0)
	case 1:
		case1(st.val1)
	case 2:
		case2(st.val2)
	}
}
