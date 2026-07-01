package trit

import (
	"iter"
	"testing"
)

// trackedSeq returns an iter.Seq over vals together with a pointer to a counter
// that records how many elements were actually pulled. This lets the tests
// prove the aggregates short-circuit rather than always draining the sequence.
func trackedSeq(vals []Trit) (iter.Seq[Trit], *int) {
	pulled := 0
	seq := func(yield func(Trit) bool) {
		for _, v := range vals {
			pulled++
			if !yield(v) {
				return
			}
		}
	}
	return seq, &pulled
}

// TestSeqShortCircuit is the key reliability test: each Seq aggregate must stop
// pulling from the sequence at the first decisive element, never consuming the
// tail. The counter makes that observable.
func TestSeqShortCircuit(t *testing.T) {
	cases := []struct {
		name       string
		fn         func(iter.Seq[Trit]) Trit
		vals       []Trit
		want       Trit
		wantPulled int
	}{
		{
			name:       "AllSeq stops at first non-True",
			fn:         AllSeq[Trit],
			vals:       []Trit{True, True, False, True, True},
			want:       False,
			wantPulled: 3,
		},
		{
			name:       "AnySeq stops at first True",
			fn:         AnySeq[Trit],
			vals:       []Trit{False, False, True, False, False},
			want:       True,
			wantPulled: 3,
		},
		{
			name:       "NoneSeq stops at first True",
			fn:         NoneSeq[Trit],
			vals:       []Trit{False, True, False, False},
			want:       False,
			wantPulled: 2,
		},
		{
			name:       "KnownSeq stops at first Unknown",
			fn:         KnownSeq[Trit],
			vals:       []Trit{True, False, Unknown, True},
			want:       False,
			wantPulled: 3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			seq, pulled := trackedSeq(c.vals)
			if got := c.fn(seq); got != c.want {
				t.Errorf("result = %s, want %s", got, c.want)
			}
			if *pulled != c.wantPulled {
				t.Errorf("pulled %d elements, want %d (short-circuit broken)",
					*pulled, c.wantPulled)
			}
		})
	}
}

// TestSeqFullDrain checks the no-short-circuit path: when no decisive element
// exists, the whole sequence is consumed and the vacuous result returned.
func TestSeqFullDrain(t *testing.T) {
	vals := []Trit{True, True, True}
	for _, tc := range []struct {
		name string
		fn   func(iter.Seq[Trit]) Trit
		want Trit
	}{
		{"AllSeq", AllSeq[Trit], True},
		{"KnownSeq", KnownSeq[Trit], True},
	} {
		seq, pulled := trackedSeq(vals)
		if got := tc.fn(seq); got != tc.want {
			t.Errorf("%s = %s, want %s", tc.name, got, tc.want)
		}
		if *pulled != len(vals) {
			t.Errorf("%s pulled %d, want full %d", tc.name, *pulled, len(vals))
		}
	}
}

// TestSeqEmpty pins the empty-sequence convention to match the variadic
// aggregates exactly (formal-logic vacuous truth).
func TestSeqEmpty(t *testing.T) {
	empty := func(yield func(Trit) bool) {}

	if AllSeq(empty) != All[Trit]() || AllSeq(empty) != True {
		t.Errorf("AllSeq(empty) must equal All() = True")
	}
	if AnySeq(empty) != Any[Trit]() || AnySeq(empty) != False {
		t.Errorf("AnySeq(empty) must equal Any() = False")
	}
	if NoneSeq(empty) != None[Trit]() || NoneSeq(empty) != True {
		t.Errorf("NoneSeq(empty) must equal None() = True")
	}
	if KnownSeq(empty) != Known[Trit]() || KnownSeq(empty) != True {
		t.Errorf("KnownSeq(empty) must equal Known() = True")
	}
}

// TestSeqParityWithVariadic asserts the Seq forms agree with their variadic
// siblings over a variety of inputs, so the two surfaces never diverge.
func TestSeqParityWithVariadic(t *testing.T) {
	inputs := [][]Trit{
		{},
		{True},
		{Unknown},
		{False},
		{True, True, True},
		{True, False, True},
		{True, Unknown, False},
		{False, False, False},
		{Unknown, Unknown},
	}
	for _, in := range inputs {
		full := func(yield func(Trit) bool) {
			for _, v := range in {
				if !yield(v) {
					return
				}
			}
		}
		if AllSeq(full) != All(in...) {
			t.Errorf("AllSeq != All for %v", in)
		}
		if AnySeq(full) != Any(in...) {
			t.Errorf("AnySeq != Any for %v", in)
		}
		if NoneSeq(full) != None(in...) {
			t.Errorf("NoneSeq != None for %v", in)
		}
		if KnownSeq(full) != Known(in...) {
			t.Errorf("KnownSeq != Known for %v", in)
		}
	}
}

// TestSeqGenericType makes sure the Seq aggregates work over any Logicable
// element type, not only Trit.
func TestSeqGenericType(t *testing.T) {
	ints := func(yield func(int) bool) {
		for _, v := range []int{1, 2, 3} {
			if !yield(v) {
				return
			}
		}
	}
	if AllSeq(ints) != True {
		t.Errorf("AllSeq over positive ints must be True")
	}
	if KnownSeq(ints) != True {
		t.Errorf("KnownSeq over non-zero ints must be True")
	}

	withZero := func(yield func(int) bool) {
		for _, v := range []int{1, 0, 3} {
			if !yield(v) {
				return
			}
		}
	}
	if KnownSeq(withZero) != False {
		t.Errorf("KnownSeq with a zero (Unknown) must be False")
	}
}
