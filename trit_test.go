package trit

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math"
	"testing"
)

// canonical is the set of the three canonical Trit states.
var canonical = []Trit{False, Unknown, True}

// TestConstants pins the numeric identity of the states. The whole package
// leans on Unknown being the int8 zero value; if that ever changes, the
// zero-value guarantee and the JSON/SQL NULL mapping silently break.
func TestConstants(t *testing.T) {
	if int8(False) != -1 || int8(Unknown) != 0 || int8(True) != 1 {
		t.Fatalf("constant drift: F=%d U=%d T=%d",
			int8(False), int8(Unknown), int8(True))
	}

	var zero Trit
	if zero != Unknown {
		t.Errorf("zero value must be Unknown, got %s", zero)
	}
}

// TestPredicatesTolerant checks that the predicates treat *any* negative int8
// as False and *any* positive int8 as True, including the extreme boundaries
// of the underlying type. This is the "tolerant" contract of the type.
func TestPredicatesTolerant(t *testing.T) {
	tests := []struct {
		in                     int8
		isFalse, isUnk, isTrue bool
	}{
		{-128, true, false, false}, // int8 minimum
		{-1, true, false, false},
		{0, false, true, false},
		{1, false, false, true},
		{127, false, false, true}, // int8 maximum
	}

	for _, tt := range tests {
		v := Trit(tt.in)
		if v.IsFalse() != tt.isFalse ||
			v.IsUnknown() != tt.isUnk ||
			v.IsTrue() != tt.isTrue {
			t.Errorf("Trit(%d): got F=%v U=%v T=%v",
				tt.in, v.IsFalse(), v.IsUnknown(), v.IsTrue())
		}

		// Exactly one predicate must hold at all times.
		n := 0
		for _, b := range []bool{v.IsFalse(), v.IsUnknown(), v.IsTrue()} {
			if b {
				n++
			}
		}
		if n != 1 {
			t.Errorf("Trit(%d): %d predicates true, want exactly 1", tt.in, n)
		}
	}
}

// TestValNormInt verifies that Val/Norm/Int all collapse non-canonical values
// to the canonical {-1,0,1} triple and stay mutually consistent.
func TestValNormInt(t *testing.T) {
	for raw := -128; raw <= 127; raw++ {
		v := Trit(int8(raw))

		var want Trit
		switch {
		case raw < 0:
			want = False
		case raw > 0:
			want = True
		default:
			want = Unknown
		}

		if got := v.Val(); got != want {
			t.Errorf("Trit(%d).Val() = %s, want %s", raw, got, want)
		}

		// Norm mutates in place and must agree with Val.
		n := v
		if got := n.Norm(); got != want || n != want {
			t.Errorf("Trit(%d).Norm() = %s (self=%s), want %s",
				raw, got, n, want)
		}

		if got := v.Int(); got != int(want) {
			t.Errorf("Trit(%d).Int() = %d, want %d", raw, got, int(want))
		}
	}
}

// TestString covers canonical and non-canonical values.
func TestString(t *testing.T) {
	cases := map[Trit]string{
		False: "False", Unknown: "Unknown", True: "True",
		Trit(-42): "False", Trit(42): "True",
	}
	for in, want := range cases {
		if got := in.String(); got != want {
			t.Errorf("Trit(%d).String() = %q, want %q", int8(in), got, want)
		}
	}
}

// TestMutators exercises the pointer-receiver helpers, including the corrected
// Clean semantics (must reset ANY state to Unknown).
func TestMutators(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		u := Unknown
		if got := u.Default(True); got != True || u != True {
			t.Errorf("Default on Unknown: got %s self=%s", got, u)
		}
		f := False
		if got := f.Default(True); got != False || f != False {
			t.Errorf("Default must not override a defined value, got %s", got)
		}
	})

	t.Run("TrueIfUnknown", func(t *testing.T) {
		u := Unknown
		u.TrueIfUnknown()
		if u != True {
			t.Errorf("TrueIfUnknown: got %s", u)
		}
		f := False
		f.TrueIfUnknown()
		if f != False {
			t.Errorf("TrueIfUnknown must not touch False, got %s", f)
		}
	})

	t.Run("FalseIfUnknown", func(t *testing.T) {
		u := Unknown
		u.FalseIfUnknown()
		if u != False {
			t.Errorf("FalseIfUnknown: got %s", u)
		}
		tr := True
		tr.FalseIfUnknown()
		if tr != True {
			t.Errorf("FalseIfUnknown must not touch True, got %s", tr)
		}
	})

	t.Run("Clean resets every state", func(t *testing.T) {
		for _, start := range []Trit{False, Unknown, True, Trit(99)} {
			v := start
			if got := v.Clean(); got != Unknown || v != Unknown {
				t.Errorf("Clean from %s: got %s self=%s, want Unknown",
					start, got, v)
			}
		}
	})

	t.Run("Set by sign", func(t *testing.T) {
		var v Trit
		if v.Set(-7); v != False {
			t.Errorf("Set(-7) = %s", v)
		}
		if v.Set(7); v != True {
			t.Errorf("Set(7) = %s", v)
		}
		if v.Set(0); v != Unknown {
			t.Errorf("Set(0) = %s", v)
		}
	})
}

// TestToBool covers the safe bool conversion and its sentinel error.
func TestToBool(t *testing.T) {
	if b, err := True.ToBool(); err != nil || b != true {
		t.Errorf("True.ToBool() = %v, %v", b, err)
	}
	if b, err := False.ToBool(); err != nil || b != false {
		t.Errorf("False.ToBool() = %v, %v", b, err)
	}

	b, err := Unknown.ToBool()
	if b != false || !errors.Is(err, ErrUnknownValue) {
		t.Errorf("Unknown.ToBool() = %v, %v; want false, ErrUnknownValue", b, err)
	}

	if True.CanBeBool() == false || Unknown.CanBeBool() == true {
		t.Errorf("CanBeBool mismatch")
	}
}

// TestJSONRoundTrip verifies the canonical mapping and lossless round-trip.
func TestJSONRoundTrip(t *testing.T) {
	want := map[Trit]string{
		False: "false", Unknown: "null", True: "true",
	}
	for v, js := range want {
		got, err := json.Marshal(v)
		if err != nil || string(got) != js {
			t.Errorf("Marshal(%s) = %q, %v; want %q", v, got, err, js)
		}

		var back Trit
		if err := json.Unmarshal([]byte(js), &back); err != nil {
			t.Fatalf("Unmarshal(%q): %v", js, err)
		}
		if back != v {
			t.Errorf("round-trip %s -> %q -> %s", v, js, back)
		}
	}
}

// TestUnmarshalJSONTolerant checks the new number-accepting behaviour and that
// junk is rejected instead of silently mapped.
func TestUnmarshalJSONTolerant(t *testing.T) {
	ok := map[string]Trit{
		"null": Unknown, "true": True, "false": False,
		"1": True, "42": True, "-1": False, "-7": False, "0": Unknown,
		"1.5": True, "-0.0": Unknown, " true ": True,
	}
	for in, want := range ok {
		var v Trit
		if err := json.Unmarshal([]byte(in), &v); err != nil {
			t.Errorf("Unmarshal(%q) unexpected error: %v", in, err)
			continue
		}
		if v != want {
			t.Errorf("Unmarshal(%q) = %s, want %s", in, v, want)
		}
	}

	for _, bad := range []string{`"true"`, `"maybe"`, `[]`, `{}`, `abc`} {
		var v Trit
		if err := json.Unmarshal([]byte(bad), &v); err == nil {
			t.Errorf("Unmarshal(%q) = %s, want error", bad, v)
		}
	}
}

// TestTextMarshaling checks MarshalText/UnmarshalText and their round-trip.
func TestTextMarshaling(t *testing.T) {
	for _, v := range canonical {
		b, err := v.MarshalText()
		if err != nil || string(b) != v.String() {
			t.Errorf("MarshalText(%s) = %q, %v", v, b, err)
		}

		var back Trit
		if err := back.UnmarshalText(b); err != nil || back != v {
			t.Errorf("UnmarshalText(%q) = %s, %v", b, back, err)
		}
	}

	var v Trit
	if err := v.UnmarshalText([]byte("garbage")); !errors.Is(err, ErrInvalidTrit) {
		t.Errorf("UnmarshalText(garbage) err = %v, want ErrInvalidTrit", err)
	}
}

// TestParseTrit covers the recognized tokens (case/space-insensitive) and the
// error path.
func TestParseTrit(t *testing.T) {
	trues := []string{"true", "T", "Yes", " y ", "ON", "1"}
	falses := []string{"false", "F", "no", "N", "off", "-1"}
	unknowns := []string{"unknown", "U", "maybe", "NULL", "nil", "none", "0", "", "  "}

	check := func(inputs []string, want Trit) {
		for _, in := range inputs {
			got, err := ParseTrit(in)
			if err != nil || got != want {
				t.Errorf("ParseTrit(%q) = %s, %v; want %s", in, got, err, want)
			}
		}
	}
	check(trues, True)
	check(falses, False)
	check(unknowns, Unknown)

	for _, bad := range []string{"tru", "2", "-2", "yeah", "?"} {
		if _, err := ParseTrit(bad); !errors.Is(err, ErrInvalidTrit) {
			t.Errorf("ParseTrit(%q) err = %v, want ErrInvalidTrit", bad, err)
		}
	}
}

// TestCompare verifies the total order False < Unknown < True, its consistency
// with normalization, and the anti-symmetry / reflexivity properties.
func TestCompare(t *testing.T) {
	order := []Trit{False, Unknown, True}
	for i, a := range order {
		for j, b := range order {
			want := 0
			switch {
			case i < j:
				want = -1
			case i > j:
				want = 1
			}
			if got := a.Compare(b); got != want {
				t.Errorf("%s.Compare(%s) = %d, want %d", a, b, got, want)
			}
			// Anti-symmetry: sign must invert when arguments swap.
			if a.Compare(b) != -b.Compare(a) {
				t.Errorf("anti-symmetry broken for %s,%s", a, b)
			}
		}
	}

	// Non-canonical inputs compare by their normalized value.
	if Trit(50).Compare(True) != 0 || Trit(-50).Compare(False) != 0 {
		t.Errorf("Compare must normalize non-canonical operands")
	}
}

// TestValuer checks the driver.Valuer mapping, in particular Unknown -> NULL.
func TestValuer(t *testing.T) {
	cases := []struct {
		in   Trit
		want driver.Value
	}{
		{True, true},
		{False, false},
		{Unknown, nil},
		{Trit(9), true},   // non-canonical positive
		{Trit(-9), false}, // non-canonical negative
	}
	for _, c := range cases {
		got, err := c.in.Value()
		if err != nil {
			t.Errorf("%s.Value() error: %v", c.in, err)
		}
		if got != c.want {
			t.Errorf("%s.Value() = %#v, want %#v", c.in, got, c.want)
		}
	}
}

// TestScanner exercises every accepted driver value kind plus the error path,
// and asserts the Value/Scan round-trip for canonical states.
func TestScanner(t *testing.T) {
	cases := []struct {
		src  any
		want Trit
	}{
		{nil, Unknown},
		{true, True},
		{false, False},
		{int64(5), True},
		{int64(-5), False},
		{int64(0), Unknown},
		{float64(1.2), True},
		{float64(-1.2), False},
		{float64(0), Unknown},
		{[]byte("true"), True},
		{[]byte("null"), Unknown},
		{"false", False},
		{"maybe", Unknown},
	}
	for _, c := range cases {
		var v Trit
		if err := v.Scan(c.src); err != nil {
			t.Errorf("Scan(%#v) error: %v", c.src, err)
			continue
		}
		if v != c.want {
			t.Errorf("Scan(%#v) = %s, want %s", c.src, v, c.want)
		}
	}

	// Unsupported type and unparseable string must error.
	var v Trit
	if err := v.Scan(struct{}{}); !errors.Is(err, ErrInvalidTrit) {
		t.Errorf("Scan(struct{}) err = %v, want ErrInvalidTrit", err)
	}
	if err := v.Scan("nonsense"); !errors.Is(err, ErrInvalidTrit) {
		t.Errorf("Scan(nonsense) err = %v, want ErrInvalidTrit", err)
	}
	if err := v.Scan([]byte("bogus")); !errors.Is(err, ErrInvalidTrit) {
		t.Errorf("Scan([]byte bogus) err = %v, want ErrInvalidTrit", err)
	}

	// Round-trip through the driver interfaces.
	for _, start := range canonical {
		dv, err := start.Value()
		if err != nil {
			t.Fatalf("Value(%s): %v", start, err)
		}
		var back Trit
		if err := back.Scan(dv); err != nil {
			t.Fatalf("Scan(%#v): %v", dv, err)
		}
		if back != start {
			t.Errorf("driver round-trip %s -> %#v -> %s", start, dv, back)
		}
	}
}

// TestUnaryMethods pins the four unary truth tables directly on the methods
// (tables_test.go covers them through the exported helpers as well).
func TestUnaryMethods(t *testing.T) {
	type row struct{ not, ma, la, ia Trit }
	want := map[Trit]row{
		False:   {not: True, ma: False, la: False, ia: False},
		Unknown: {not: Unknown, ma: True, la: False, ia: True},
		True:    {not: False, ma: True, la: True, ia: False},
	}
	for _, v := range canonical {
		w := want[v]
		if v.Not() != w.not || v.Ma() != w.ma || v.La() != w.la || v.Ia() != w.ia {
			t.Errorf("unary(%s): Not=%s Ma=%s La=%s Ia=%s",
				v, v.Not(), v.Ma(), v.La(), v.Ia())
		}
	}
}

// TestFromFloatSpecials documents how the float mapping treats non-finite and
// signed-zero inputs (used by the JSON/SQL number paths).
func TestFromFloatSpecials(t *testing.T) {
	cases := map[float64]Trit{
		math.Inf(1):          True,
		math.Inf(-1):         False,
		math.NaN():           Unknown, // NaN is neither >0 nor <0
		math.Copysign(0, -1): Unknown,
	}
	for in, want := range cases {
		if got := fromFloat(in); got != want {
			t.Errorf("fromFloat(%v) = %s, want %s", in, got, want)
		}
	}
}
