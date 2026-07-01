package trit

import (
	"math"
	"testing"
)

// TestLogicToTritAllTypes drives the reflection-free converter through every
// concrete type in the Logicable set, including the boundary magnitudes of
// each. The invariant: the result is always one of the three canonical states
// and matches the sign of the input.
func TestLogicToTritAllTypes(t *testing.T) {
	// bool
	if Define(true) != True || Define(false) != False {
		t.Errorf("bool mapping broken")
	}

	// Trit passthrough (normalized).
	if Define(Trit(7)) != True || Define(Trit(-7)) != False ||
		Define(Trit(0)) != Unknown {
		t.Errorf("Trit passthrough must normalize")
	}

	// Signed integers: negative/zero/positive incl. type extremes.
	checkSigned := func(name string, neg, zero, pos Trit) {
		if neg != False || zero != Unknown || pos != True {
			t.Errorf("%s: neg=%s zero=%s pos=%s", name, neg, zero, pos)
		}
	}
	checkSigned("int",
		Define(math.MinInt), Define(0), Define(math.MaxInt))
	checkSigned("int8",
		Define(int8(math.MinInt8)), Define(int8(0)), Define(int8(math.MaxInt8)))
	checkSigned("int16",
		Define(int16(math.MinInt16)), Define(int16(0)), Define(int16(math.MaxInt16)))
	checkSigned("int32",
		Define(int32(math.MinInt32)), Define(int32(0)), Define(int32(math.MaxInt32)))
	checkSigned("int64",
		Define(int64(math.MinInt64)), Define(int64(0)), Define(int64(math.MaxInt64)))

	// Unsigned integers: only zero -> Unknown, anything else -> True.
	checkUnsigned := func(name string, zero, pos Trit) {
		if zero != Unknown || pos != True {
			t.Errorf("%s: zero=%s pos=%s", name, zero, pos)
		}
	}
	checkUnsigned("uint", Define(uint(0)), Define(uint(math.MaxUint)))
	checkUnsigned("uint8", Define(uint8(0)), Define(uint8(math.MaxUint8)))
	checkUnsigned("uint16", Define(uint16(0)), Define(uint16(math.MaxUint16)))
	checkUnsigned("uint32", Define(uint32(0)), Define(uint32(math.MaxUint32)))
	checkUnsigned("uint64", Define(uint64(0)), Define(uint64(math.MaxUint64)))

	// Floats.
	checkSigned("float32",
		Define(float32(-1)), Define(float32(0)), Define(float32(1)))
	checkSigned("float64",
		Define(-math.MaxFloat64), Define(0.0), Define(math.MaxFloat64))
}

// TestConvert checks the batch converter preserves order and normalizes.
func TestConvert(t *testing.T) {
	got := Convert(1, 0, -1, 5, -9)
	want := []Trit{True, Unknown, False, True, False}
	if len(got) != len(want) {
		t.Fatalf("Convert length = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Convert[%d] = %s, want %s", i, got[i], want[i])
		}
	}

	if len(Convert[int]()) != 0 {
		t.Errorf("Convert() on empty must return empty slice")
	}
}

// TestFreeFunctionPredicates ties the generic predicate/setter free functions
// to their expected results across a couple of Logicable types.
func TestFreeFunctionPredicates(t *testing.T) {
	if !IsTrue(1) || !IsTrue(true) || IsTrue(0) {
		t.Errorf("IsTrue mismatch")
	}
	if !IsFalse(-1) || IsFalse(0) || IsFalse(true) {
		t.Errorf("IsFalse mismatch")
	}
	if !IsUnknown(0) || IsUnknown(1) {
		t.Errorf("IsUnknown mismatch")
	}

	var v Trit
	if Set(&v, -3.5); v != False {
		t.Errorf("Set(-3.5) = %s", v)
	}

	u := Unknown
	if Default(&u, uint(4)); u != True {
		t.Errorf("Default with uint(4) = %s", u)
	}
	d := False
	if Default(&d, true); d != False {
		t.Errorf("Default must not override a defined value, got %s", d)
	}
}

// TestFreeFunctionsMatchMethods asserts that every generic operation free
// function agrees with the corresponding method for all 3x3 (or 3) inputs.
// This guarantees the two API surfaces never drift apart.
func TestFreeFunctionsMatchMethods(t *testing.T) {
	unary := []struct {
		name string
		fn   func(Trit) Trit
		m    func(Trit) Trit
	}{
		{"Not", func(x Trit) Trit { return Not(x) }, Trit.Not},
		{"Ma", func(x Trit) Trit { return Ma(x) }, Trit.Ma},
		{"La", func(x Trit) Trit { return La(x) }, Trit.La},
		{"Ia", func(x Trit) Trit { return Ia(x) }, Trit.Ia},
	}
	for _, op := range unary {
		for _, a := range canonical {
			if op.fn(a) != op.m(a) {
				t.Errorf("%s(%s): free=%s method=%s",
					op.name, a, op.fn(a), op.m(a))
			}
		}
	}

	binary := []struct {
		name string
		fn   func(Trit, Trit) Trit
		m    func(Trit, Trit) Trit
	}{
		{"And", func(a, b Trit) Trit { return And(a, b) }, Trit.And},
		{"Or", func(a, b Trit) Trit { return Or(a, b) }, Trit.Or},
		{"Xor", func(a, b Trit) Trit { return Xor(a, b) }, Trit.Xor},
		{"Nand", func(a, b Trit) Trit { return Nand(a, b) }, Trit.Nand},
		{"Nor", func(a, b Trit) Trit { return Nor(a, b) }, Trit.Nor},
		{"Nxor", func(a, b Trit) Trit { return Nxor(a, b) }, Trit.Nxor},
		{"Imp", func(a, b Trit) Trit { return Imp(a, b) }, Trit.Imp},
		{"Nimp", func(a, b Trit) Trit { return Nimp(a, b) }, Trit.Nimp},
		{"Eq", func(a, b Trit) Trit { return Eq(a, b) }, Trit.Eq},
		{"Neq", func(a, b Trit) Trit { return Neq(a, b) }, Trit.Neq},
		{"Min", func(a, b Trit) Trit { return Min(a, b) }, Trit.Min},
		{"Max", func(a, b Trit) Trit { return Max(a, b) }, Trit.Max},
	}
	for _, op := range binary {
		for _, a := range canonical {
			for _, b := range canonical {
				if op.fn(a, b) != op.m(a, b) {
					t.Errorf("%s(%s,%s): free=%s method=%s",
						op.name, a, b, op.fn(a, b), op.m(a, b))
				}
			}
		}
	}
}

// TestMinMaxAliases documents that Min/Max are exact aliases of And/Or.
func TestMinMaxAliases(t *testing.T) {
	for _, a := range canonical {
		for _, b := range canonical {
			if a.Min(b) != a.And(b) {
				t.Errorf("Min(%s,%s) != And", a, b)
			}
			if a.Max(b) != a.Or(b) {
				t.Errorf("Max(%s,%s) != Or", a, b)
			}
		}
	}
}

// TestAggregatesBasic covers the "happy path" of the slice aggregates.
func TestAggregatesBasic(t *testing.T) {
	if All(True, True, True) != True {
		t.Errorf("All(T,T,T) != True")
	}
	if All(True, Unknown, True) != False {
		t.Errorf("All with Unknown must be False")
	}
	if All(True, False) != False {
		t.Errorf("All with False must be False")
	}

	if Any(False, False, True) != True {
		t.Errorf("Any with a True must be True")
	}
	if Any(False, Unknown, False) != False {
		t.Errorf("Any without True must be False")
	}

	if None(False, False) != True {
		t.Errorf("None without True must be True")
	}
	if None(False, True) != False {
		t.Errorf("None with True must be False")
	}

	if Known(True, False, True) != True {
		t.Errorf("Known of definite values must be True")
	}
	if Known(True, Unknown) != False {
		t.Errorf("Known with Unknown must be False")
	}
}

// TestAggregatesEmpty pins the agreed formal-logic empty-input convention. This
// is the invariant that the audit flagged as inconsistent; keep it explicit.
func TestAggregatesEmpty(t *testing.T) {
	cases := []struct {
		name string
		got  Trit
		want Trit
	}{
		{"All", All[Trit](), True},     // vacuous truth of forall
		{"Any", Any[Trit](), False},    // nothing can be true
		{"None", None[Trit](), True},   // nothing is true
		{"Known", Known[Trit](), True}, // nothing is unknown
		{"Consensus", Consensus[Trit](), Unknown},
		{"Majority", Majority[Trit](), Unknown},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s() on empty = %s, want %s", c.name, c.got, c.want)
		}
	}
}

// TestAggregatesMixedTypes makes sure the aggregates work over the generic
// Logicable set, not just Trit.
func TestAggregatesMixedTypes(t *testing.T) {
	if All(1, 2, 3) != True {
		t.Errorf("All over positive ints must be True")
	}
	if Any(0, 0, 5) != True {
		t.Errorf("Any over ints with a positive must be True")
	}
	if Known(1, -1, 0) != False {
		t.Errorf("Known with a zero (Unknown) must be False")
	}
	if Consensus(true, true, true) != True {
		t.Errorf("Consensus of all-true bools must be True")
	}
	if Consensus(1.0, -1.0) != Unknown {
		t.Errorf("Consensus of mixed signs must be Unknown")
	}
}

// TestConsensus covers agreement, disagreement and the Unknown-poison rule.
func TestConsensus(t *testing.T) {
	if Consensus(True, True, True) != True {
		t.Errorf("all True -> True")
	}
	if Consensus(False, False) != False {
		t.Errorf("all False -> False")
	}
	if Consensus(True, False) != Unknown {
		t.Errorf("mixed -> Unknown")
	}
	// A single Unknown poisons the result even if the rest agree.
	if Consensus(True, True, Unknown) != Unknown {
		t.Errorf("any Unknown -> Unknown")
	}
}

// TestMajority checks the strict-majority rule and its ties.
func TestMajority(t *testing.T) {
	if Majority(True, True, False) != True {
		t.Errorf("2/3 True -> True")
	}
	if Majority(False, False, True) != False {
		t.Errorf("2/3 False -> False")
	}
	// Exact tie is not a strict majority.
	if Majority(True, False) != Unknown {
		t.Errorf("tie -> Unknown")
	}
	// Unknown votes dilute the majority.
	if Majority(True, Unknown, Unknown) != Unknown {
		t.Errorf("1/3 True -> Unknown")
	}
	if Majority(True, True, False, Unknown) != Unknown {
		t.Errorf("2/4 True is not a strict majority (>2 needed) -> Unknown")
	}
}

// TestRandomReachability is the semantic test the audit called for: over many
// draws every reachable state must actually occur, and the empirical
// distribution must match the requested probability within tolerance. The old
// implementation silently never produced True.
func TestRandomReachability(t *testing.T) {
	const n = 200000

	count := func(up ...uint8) map[Trit]int {
		m := map[Trit]int{}
		for i := 0; i < n; i++ {
			m[Random(up...)]++
		}
		return m
	}

	// approxEq checks that got/n is within tol of want (a fraction).
	approxEq := func(got int, want, tol float64) bool {
		f := float64(got) / float64(n)
		return math.Abs(f-want) <= tol
	}

	t.Run("default p=33", func(t *testing.T) {
		m := count()
		if m[True] == 0 {
			t.Fatalf("True is unreachable by default: %v", m)
		}
		if m[False] == 0 || m[Unknown] == 0 {
			t.Fatalf("some state unreachable: %v", m)
		}
		// ~33% Unknown, remaining split evenly (~33.5% each).
		if !approxEq(m[Unknown], 0.33, 0.03) {
			t.Errorf("Unknown share off: %v", m)
		}
		// True and False must be roughly symmetric.
		if !approxEq(m[True], 0.335, 0.03) || !approxEq(m[False], 0.335, 0.03) {
			t.Errorf("True/False not symmetric: %v", m)
		}
	})

	t.Run("p=50 stays symmetric", func(t *testing.T) {
		m := count(50)
		if m[True] == 0 || m[False] == 0 {
			t.Fatalf("True/False unreachable at p=50: %v", m)
		}
		if !approxEq(m[Unknown], 0.50, 0.03) {
			t.Errorf("Unknown share off at p=50: %v", m)
		}
		if !approxEq(m[True], 0.25, 0.03) || !approxEq(m[False], 0.25, 0.03) {
			t.Errorf("True/False not ~25%% each at p=50: %v", m)
		}
	})

	t.Run("p=0 never Unknown", func(t *testing.T) {
		m := count(0)
		if m[Unknown] != 0 {
			t.Errorf("p=0 produced Unknown: %v", m)
		}
		if m[True] == 0 || m[False] == 0 {
			t.Errorf("p=0 must yield both True and False: %v", m)
		}
	})

	t.Run("p=100 always Unknown", func(t *testing.T) {
		m := count(100)
		if m[True] != 0 || m[False] != 0 {
			t.Errorf("p=100 must be Unknown only: %v", m)
		}
	})

	t.Run("summed and clamped", func(t *testing.T) {
		// 60+60 sums to 120, clamped to 100 -> always Unknown.
		m := count(60, 60)
		if m[True] != 0 || m[False] != 0 {
			t.Errorf("summed probability must clamp to 100: %v", m)
		}
	})
}

// TestTritterInterface confirms *Trit satisfies Tritter (also asserted at
// compile time) and that the interface methods behave.
func TestTritterInterface(t *testing.T) {
	var tr Tritter = new(Trit)
	if tr.IsUnknown() != true || tr.Int() != 0 || tr.String() != "Unknown" {
		t.Errorf("zero *Trit via Tritter misbehaves")
	}
}
