package opt

import "testing"

func TestIsSome(t *testing.T) {
	op := Some("hi")
	want := true
	got := op.IsSome()
	if want != got {
		t.Errorf("Expected %t, got %t.", want, got)
	}

	op = None[string]()
	want = false
	got = op.IsSome()
	if want != got {
		t.Errorf("Expected %t, got %t.", want, got)
	}
}

func TestGet(t *testing.T) {
	op := Some("hi")
	wantA, wantB := "hi", true
	gotA, gotB := op.Get()
	if wantA != gotA || wantB != gotB {
		t.Errorf("Expected %s/%t, got %s/%t.", wantA, wantB, gotA, gotB)
	}

	op = None[string]()
	wantA, wantB = "", false
	gotA, gotB = op.Get()
	if wantA != gotA || wantB != gotB {
		t.Errorf("Expected %s/%t, got %s/%t.", wantA, wantB, gotA, gotB)
	}
}
