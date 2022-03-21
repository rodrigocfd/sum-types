package st2

import "testing"

func TestWhich(t *testing.T) {
	st := New0[int, string](0x100)
	want := uint8(0)
	got := st.Which()
	if want != got {
		t.Errorf("Expected %d, got %d.", want, got)
	}

	st = New1[int]("hi")
	want = uint8(1)
	got = st.Which()
	if want != got {
		t.Errorf("Expected %d, got %d.", want, got)
	}
}

func TestHas(t *testing.T) {
	st := New0[int, string](0x100)
	want := true
	got := st.Has0()
	if want != got {
		t.Errorf("Expected %t, got %t.", want, got)
	}

	st = New1[int]("hi")
	want = true
	got = st.Has1()
	if want != got {
		t.Errorf("Expected %t, got %t.", want, got)
	}
}

func TestGet(t *testing.T) {
	st := New0[int, string](0x100)
	wantA, wantB := 0x100, true
	gotA, gotB := st.Get0()
	if wantA != gotA || wantB != gotB {
		t.Errorf("Expected %d/%t, got %d/%t.", wantA, wantB, gotA, gotB)
	}

	st = New1[int]("hi")
	wantC, wantD := "hi", true
	gotC, gotD := st.Get1()
	if wantC != gotC || wantD != gotD {
		t.Errorf("Expected %s/%t, got %s/%t.", wantC, wantD, gotC, gotD)
	}
}

func TestMatch(t *testing.T) {
	st := New0[int, string](0x100)
	st.Match(
		func(got int) {
			want := 0x100
			if want != got {
				t.Errorf("Expected %d, got %d.", want, got)
			}
		}, func(_ string) {
			t.Errorf("Expected not to be called.")
		})

	st = New1[int]("hi")
	st.Match(
		func(_ int) {
			t.Errorf("Expected not to be called.")
		}, func(got string) {
			want := "hi"
			if want != got {
				t.Errorf("Expected %s, got %s.", want, got)
			}
		})
}
