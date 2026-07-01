package trit

import (
	"iter"
	"math/rand/v2"
)

// This line asserts at compile time that the type *Trit
// implements the Tritter interface.
var _ Tritter = (*Trit)(nil)

// Logicable is a special data type from which to determine the state of Trit
// in the context of three-valued logic.
//
// Signed and floating-point values map by sign: any negative value is False,
// any positive value is True, and zero is Unknown. Unsigned values are True
// when non-zero and Unknown when zero (they can never be negative). A bool
// maps true to True and false to False.
type Logicable interface {
	bool | Trit |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Tritter is the behavioural contract of a three-valued digit. The concrete
// *Trit type satisfies it (see the compile-time assertion above).
type Tritter interface {
	IsTrue() bool
	IsFalse() bool
	IsUnknown() bool
	Int() int
	String() string
}

// fromInt maps a signed integer onto a Trit by its sign.
func fromInt(v int64) Trit {
	switch {
	case v > 0:
		return True
	case v < 0:
		return False
	default:
		return Unknown
	}
}

// fromUint maps an unsigned integer onto a Trit. Unsigned values are never
// negative, so the result is either True (non-zero) or Unknown (zero).
func fromUint(v uint64) Trit {
	if v > 0 {
		return True
	}
	return Unknown
}

// fromFloat maps a floating-point value onto a Trit by its sign. Note that
// negative zero has the same sign bit semantics as zero here and yields
// Unknown, while NaN (which is neither > 0 nor < 0) also yields Unknown.
func fromFloat(v float64) Trit {
	switch {
	case v > 0:
		return True
	case v < 0:
		return False
	default:
		return Unknown
	}
}

// logicToTrit converts any Logicable value to a normalized Trit. It uses a
// plain type switch (no reflection): the switch already narrows the dynamic
// type, so the mapping is a direct, allocation-free branch.
func logicToTrit[T Logicable](v T) Trit {
	switch x := any(v).(type) {
	case bool:
		if x {
			return True
		}
		return False
	case Trit:
		return x.Val()
	case int:
		return fromInt(int64(x))
	case int8:
		return fromInt(int64(x))
	case int16:
		return fromInt(int64(x))
	case int32:
		return fromInt(int64(x))
	case int64:
		return fromInt(x)
	case uint:
		return fromUint(uint64(x))
	case uint8:
		return fromUint(uint64(x))
	case uint16:
		return fromUint(uint64(x))
	case uint32:
		return fromUint(uint64(x))
	case uint64:
		return fromUint(x)
	case float32:
		return fromFloat(float64(x))
	case float64:
		return fromFloat(x)
	}

	// Unreachable for the Logicable type set, but keeps the function total.
	return Unknown
}

// Default sets the default value for the trit-object
// if this one has a Unknown state.
//
// Example usage:
//
//	t := trit.Unknown
//	trit.Default(&t, trit.True)
//	fmt.Println(t.String()) // Output: True
func Default[T Logicable](t *Trit, v T) Trit {
	// If the trit is not Unknown, return the trit.
	if t.Val() != Unknown {
		return *t
	}

	*t = logicToTrit(v)
	return *t
}

// IsFalse checks if the trit-object is False.
//
// See Trit.IsFalse() for more information.
func IsFalse[T Logicable](t T) bool {
	return logicToTrit(t).IsFalse()
}

// IsUnknown checks if the trit-object is Unknown.
//
// See Trit.IsUnknown() for more information.
func IsUnknown[T Logicable](t T) bool {
	return logicToTrit(t).IsUnknown()
}

// IsTrue checks if the trit-object is True.
//
// See Trit.IsTrue() for more information.
func IsTrue[T Logicable](t T) bool {
	return logicToTrit(t).IsTrue()
}

// Set sets the value of the trit-object.
//
// See Trit.Set() for more information.
func Set[T Logicable](t *Trit, v T) Trit {
	*t = logicToTrit(v)
	return *t
}

// Convert converts the any Logicable types to Trit.
//
// Example usage:
//
//	tuf := trit.Convert(true, 0, -1)
//	fmt.Println(tuf[0].String()) // Output: True
//	fmt.Println(tuf[1].String()) // Output: Unknown
//	fmt.Println(tuf[2].String()) // Output: False
func Convert[T Logicable](v ...T) []Trit {
	trits := make([]Trit, len(v))
	for i, value := range v {
		trits[i] = logicToTrit(value)
	}

	return trits
}

// Define converts the any Logicable type to Trit.
//
// Example usage:
//
//	t := trit.Define(true)
//	fmt.Println(t.String()) // Output: True
func Define[T Logicable](v T) Trit {
	return logicToTrit(v)
}

// All returns True if every value is True, and False as soon as any value is
// False or Unknown.
//
// It follows the vacuous-truth convention of universal quantification: with
// no arguments the predicate holds trivially, so All() returns True.
//
// Example usage:
//
//	t := trit.All(trit.True, trit.True, trit.True)
//	fmt.Println(t.String()) // Output: True
func All[T Logicable](t ...T) Trit {
	for _, v := range t {
		trit := logicToTrit(v)
		if trit.IsFalse() || trit.IsUnknown() {
			return False
		}
	}

	return True
}

// Any returns True as soon as any value is True, and False otherwise.
//
// It follows the convention of existential quantification: with no arguments
// there is nothing that can be True, so Any() returns False.
//
// Example usage:
//
//	t := trit.Any(trit.True, trit.False, trit.False)
//	fmt.Println(t.String()) // Output: True
func Any[T Logicable](t ...T) Trit {
	for _, v := range t {
		if logicToTrit(v).IsTrue() {
			return True
		}
	}

	return False
}

// None returns True if none of the values are True, and False as soon as any
// value is True.
//
// It is the negation of Any: None() returns True for an empty input.
//
// Example usage:
//
//	t := trit.None(trit.False, trit.False, trit.False)
//	fmt.Println(t.String()) // Output: True
func None[T Logicable](t ...T) Trit {
	for _, v := range t {
		if logicToTrit(v).IsTrue() {
			return False
		}
	}

	return True
}

// Not performs a logical NOT operation on a Trit-Like value and
// returns the result as Trit.
//
// See Trit.Not() for more information.
func Not[T Logicable](t T) Trit {
	return logicToTrit(t).Not()
}

// Ma performs a logical MA (Modus Ponens Absorption) operation
// on a Trit-Like values and returns the result as Trit.
//
// See Trit.Ma() for more information.
func Ma[T Logicable](t T) Trit {
	return logicToTrit(t).Ma()
}

// La performs a logical LA (Law of Absorption) operation on a Trit-Like
// value and returns the result as Trit.
//
// See Trit.La() for more information.
func La[T Logicable](t T) Trit {
	return logicToTrit(t).La()
}

// Ia performs a logical IA (Implication Absorption) operation on a Trit-Like
// value and returns the result as Trit.
//
// See Trit.Ia() for more information.
func Ia[T Logicable](t T) Trit {
	return logicToTrit(t).Ia()
}

// And performs a logical AND operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.And() for more information.
func And[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).And(logicToTrit(b))
}

// Or performs a logical OR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Or() for more information.
func Or[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Or(logicToTrit(b))
}

// Xor performs a logical XOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Xor() for more information.
func Xor[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Xor(logicToTrit(b))
}

// Nand performs a logical NAND operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nand() for more information.
func Nand[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Nand(logicToTrit(b))
}

// Nor performs a logical NOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nor() for more information.
func Nor[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Nor(logicToTrit(b))
}

// Nxor performs a logical NXOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nxor() for more information.
func Nxor[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Nxor(logicToTrit(b))
}

// Min performs a logical MIN operation between two Trit-Like values
// and returns the result as Trit. MIN is an alias of AND.
//
// See Trit.Min() for more information.
func Min[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Min(logicToTrit(b))
}

// Max performs a logical MAX operation between two Trit-Like values
// and returns the result as Trit. MAX is an alias of OR.
//
// See Trit.Max() for more information.
func Max[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Max(logicToTrit(b))
}

// Imp performs a logical IMP operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Imp() for more information.
func Imp[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Imp(logicToTrit(b))
}

// Nimp performs a logical NIMP operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nimp() for more information.
func Nimp[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Nimp(logicToTrit(b))
}

// Eq performs a logical EQ operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Eq() for more information.
func Eq[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Eq(logicToTrit(b))
}

// Neq performs a logical NEQ operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Neq() for more information.
func Neq[T, U Logicable](a T, b U) Trit {
	return logicToTrit(a).Neq(logicToTrit(b))
}

// Known returns True if every value is definite (True or False), and False as
// soon as any value is Unknown.
//
// It follows the vacuous-truth convention: with no arguments there is no
// Unknown value, so Known() returns True.
//
// Example usage:
//
//	a := trit.Known(trit.True, trit.False, trit.Unknown)
//	fmt.Println(a.String()) // Output: False
//
//	b := trit.Known(trit.True, trit.True, trit.False)
//	fmt.Println(b.String()) // Output: True
func Known[T Logicable](ts ...T) Trit {
	for _, t := range ts {
		if logicToTrit(t).IsUnknown() {
			return False
		}
	}

	return True
}

// Random returns a random Trit value.
//
// The optional argument sets the probability, in percent, of the Unknown
// outcome; the remaining probability is split evenly between True and False.
// Multiple arguments are summed and the total is clamped to the range [0, 100].
// With no argument the probability of Unknown is 33%.
//
// Example usage:
//
//	a := trit.Random()
//	fmt.Println(a.String()) // Output: True, False or Unknown
//
//	b := trit.Random(0)
//	fmt.Println(b.String()) // Output: True or False (never Unknown)
//
//	c := trit.Random(50)
//	fmt.Println(c.String()) // Output: Unknown with a probability of 50%
func Random(up ...uint8) Trit {
	p := 33
	if len(up) > 0 {
		p = 0
		for _, v := range up {
			p += int(v)
		}
	}

	if p > 100 {
		p = 100
	}

	// value is drawn uniformly from [0, 100). The first p units go to Unknown;
	// the remaining (100-p) units are split evenly: the lower half to True and
	// the upper half to False. This keeps True and False symmetric for any p.
	value := rand.IntN(100)
	if value < p {
		return Unknown
	}

	if value < p+(100-p)/2 {
		return True
	}

	return False
}

// Consensus returns True if all input trits are True, False if all are False,
// and Unknown otherwise (including when any trit is Unknown).
//
// With no arguments there is no shared position to agree on, so Consensus()
// returns Unknown.
//
// Example usage:
//
//	t1, t2, t3 := trit.True, trit.True, trit.Unknown
//	result := trit.Consensus(t1, t2, t3)
//	// result will be Unknown, as not all trits are the same
func Consensus[T Logicable](trits ...T) Trit {
	if len(trits) == 0 {
		return Unknown
	}

	countT := 0
	countF := 0
	for _, x := range trits {
		switch logicToTrit(x) {
		case True:
			countT++
		case False:
			countF++
		default:
			return Unknown
		}
	}

	if countT == len(trits) {
		return True
	} else if countF == len(trits) {
		return False
	}

	return Unknown
}

// Majority returns True if more than half of the input trits are True, False
// if more than half are False, and Unknown otherwise.
//
// With no arguments there is no majority, so Majority() returns Unknown.
//
// Example usage:
//
//	t1, t2, t3, t4 := trit.True, trit.True, trit.False, trit.Unknown
//	result := trit.Majority(t1, t2, t3, t4)
//	// result will be True, as more than half of the trits are True
func Majority[T Logicable](trits ...T) Trit {
	countT := 0
	countF := 0
	for _, x := range trits {
		switch logicToTrit(x) {
		case True:
			countT++
		case False:
			countF++
		}
	}

	if countT > len(trits)/2 {
		return True
	} else if countF > len(trits)/2 {
		return False
	}

	return Unknown
}

// AllSeq is the iterator form of All. It returns True if every value produced
// by seq is True, and False as soon as a False or Unknown value appears. It
// stops pulling from seq at the first decisive value, so it can short-circuit
// over lazy or infinite sequences. An empty sequence yields True (vacuous
// truth), matching All.
//
// Example usage:
//
//	t := trit.AllSeq(slices.Values([]trit.Trit{trit.True, trit.True}))
//	fmt.Println(t.String()) // Output: True
func AllSeq[T Logicable](seq iter.Seq[T]) Trit {
	for v := range seq {
		trit := logicToTrit(v)
		if trit.IsFalse() || trit.IsUnknown() {
			return False
		}
	}

	return True
}

// AnySeq is the iterator form of Any. It returns True as soon as any value
// produced by seq is True, and False otherwise. It stops pulling from seq at
// the first True. An empty sequence yields False, matching Any.
func AnySeq[T Logicable](seq iter.Seq[T]) Trit {
	for v := range seq {
		if logicToTrit(v).IsTrue() {
			return True
		}
	}

	return False
}

// NoneSeq is the iterator form of None. It returns True if none of the values
// produced by seq are True, and False as soon as a True appears. An empty
// sequence yields True, matching None.
func NoneSeq[T Logicable](seq iter.Seq[T]) Trit {
	for v := range seq {
		if logicToTrit(v).IsTrue() {
			return False
		}
	}

	return True
}

// KnownSeq is the iterator form of Known. It returns True if every value
// produced by seq is definite (True or False), and False as soon as an Unknown
// appears. It stops pulling from seq at the first Unknown. An empty sequence
// yields True (nothing is unknown), matching Known.
func KnownSeq[T Logicable](seq iter.Seq[T]) Trit {
	for v := range seq {
		if logicToTrit(v).IsUnknown() {
			return False
		}
	}

	return True
}
