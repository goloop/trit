package trit

import (
	"encoding/json"
	"testing"
)

// The truth tables have only three states, so "property" tests can be
// exhaustive rather than randomized: every law below is checked over the full
// domain, which is a complete proof for these operators.

// TestPropNotInvolution verifies Not is an involution: Not(Not(x)) == x.
func TestPropNotInvolution(t *testing.T) {
	for _, x := range canonical {
		if got := x.Not().Not(); got != x {
			t.Errorf("Not(Not(%s)) = %s", x, got)
		}
	}
}

// TestPropNegatedCompositions checks that each "N-" operator is exactly the
// negation of its base operator. This is the guarantee that keeps the derived
// tables from silently drifting.
func TestPropNegatedCompositions(t *testing.T) {
	pairs := []struct {
		name      string
		neg, base func(Trit, Trit) Trit
	}{
		{"Nand", Trit.Nand, Trit.And},
		{"Nor", Trit.Nor, Trit.Or},
		{"Nxor", Trit.Nxor, Trit.Xor},
		{"Nimp", Trit.Nimp, Trit.Imp},
		{"Neq", Trit.Neq, Trit.Eq},
	}
	for _, p := range pairs {
		for _, a := range canonical {
			for _, b := range canonical {
				if p.neg(a, b) != p.base(a, b).Not() {
					t.Errorf("%s(%s,%s) != Not(base)", p.name, a, b)
				}
			}
		}
	}
}

// TestPropDeMorgan verifies De Morgan's laws hold in this (Kleene) three-valued
// logic, where And = min, Or = max and Not = arithmetic negation.
func TestPropDeMorgan(t *testing.T) {
	for _, a := range canonical {
		for _, b := range canonical {
			// Not(a AND b) == (Not a) OR (Not b)
			if a.And(b).Not() != a.Not().Or(b.Not()) {
				t.Errorf("De Morgan AND failed for %s,%s", a, b)
			}
			// Not(a OR b) == (Not a) AND (Not b)
			if a.Or(b).Not() != a.Not().And(b.Not()) {
				t.Errorf("De Morgan OR failed for %s,%s", a, b)
			}
		}
	}
}

// TestPropCommutativity verifies the symmetric operators are commutative and,
// as a guard, that implication is *not* (asymmetry is a real property here).
func TestPropCommutativity(t *testing.T) {
	comm := []struct {
		name string
		fn   func(Trit, Trit) Trit
	}{
		{"And", Trit.And}, {"Or", Trit.Or}, {"Xor", Trit.Xor},
		{"Nand", Trit.Nand}, {"Nor", Trit.Nor}, {"Nxor", Trit.Nxor},
		{"Eq", Trit.Eq}, {"Neq", Trit.Neq},
		{"Min", Trit.Min}, {"Max", Trit.Max},
	}
	for _, op := range comm {
		for _, a := range canonical {
			for _, b := range canonical {
				if op.fn(a, b) != op.fn(b, a) {
					t.Errorf("%s not commutative at %s,%s", op.name, a, b)
				}
			}
		}
	}

	// Sanity: Imp must have at least one asymmetric pair, otherwise the test
	// above would be meaningless for it.
	if True.Imp(False) == False.Imp(True) {
		t.Errorf("Imp unexpectedly symmetric; test guard is ineffective")
	}
}

// TestPropIdempotence verifies x AND x == x and x OR x == x (normalized).
func TestPropIdempotence(t *testing.T) {
	for _, x := range canonical {
		if x.And(x) != x || x.Or(x) != x {
			t.Errorf("idempotence failed for %s: And=%s Or=%s",
				x, x.And(x), x.Or(x))
		}
	}
}

// TestPropAssociativity verifies And/Or/Min/Max are associative over the full
// domain (a key law for using them as fold operators).
func TestPropAssociativity(t *testing.T) {
	ops := []struct {
		name string
		fn   func(Trit, Trit) Trit
	}{
		{"And", Trit.And}, {"Or", Trit.Or},
	}
	for _, op := range ops {
		for _, a := range canonical {
			for _, b := range canonical {
				for _, c := range canonical {
					l := op.fn(op.fn(a, b), c)
					r := op.fn(a, op.fn(b, c))
					if l != r {
						t.Errorf("%s not associative at %s,%s,%s", op.name, a, b, c)
					}
				}
			}
		}
	}
}

// TestPropSerializationRoundTrip checks JSON and text encodings are lossless
// for the canonical states.
func TestPropSerializationRoundTrip(t *testing.T) {
	for _, x := range canonical {
		// JSON.
		js, err := json.Marshal(x)
		if err != nil {
			t.Fatalf("Marshal(%s): %v", x, err)
		}
		var jb Trit
		if err := json.Unmarshal(js, &jb); err != nil || jb != x {
			t.Errorf("JSON round-trip %s -> %s -> %s (%v)", x, js, jb, err)
		}

		// Text.
		txt, err := x.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText(%s): %v", x, err)
		}
		var tb Trit
		if err := tb.UnmarshalText(txt); err != nil || tb != x {
			t.Errorf("text round-trip %s -> %s -> %s (%v)", x, txt, tb, err)
		}
	}
}

// TestPropCompareTotalOrder verifies Compare induces a consistent total order:
// exactly one of <, ==, > holds and transitivity is respected.
func TestPropCompareTotalOrder(t *testing.T) {
	all := []Trit{False, Unknown, True}
	for _, a := range all {
		for _, b := range all {
			c := a.Compare(b)
			if c < -1 || c > 1 {
				t.Errorf("Compare(%s,%s) out of range: %d", a, b, c)
			}
			if (c == 0) != (a == b) {
				t.Errorf("Compare(%s,%s)==0 disagrees with equality", a, b)
			}
			for _, d := range all {
				// Transitivity: a<=b<=d implies a<=d.
				if a.Compare(b) <= 0 && b.Compare(d) <= 0 && a.Compare(d) > 0 {
					t.Errorf("transitivity broken: %s,%s,%s", a, b, d)
				}
			}
		}
	}
}
