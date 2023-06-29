package trit

import (
	"reflect"

	"github.com/goloop/g"
)

// Logicable is a special data type from which to determine the state of trit.
type Logicable interface {
	bool | g.Numerable | Trit
}

// The logicToTrit function converts any logic type to Trit.
func logicToTrit[T Logicable](v T) Trit {
	switch any(v).(type) {
	case bool:
		if any(v).(bool) {
			return True
		}
		return False
	case int, int8, int16, int32, int64:
		switch reflect.TypeOf(v).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value := reflect.ValueOf(v).Int()
			if value > 0 {
				return True
			} else if value < 0 {
				return False
			}

			return Unknown
		}
	case uint, uint8, uint16, uint32, uint64:
		switch reflect.TypeOf(v).Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64:
			value := reflect.ValueOf(v).Uint()
			if value > 0 {
				return True
			}

			// Can't be less than 0
			return Unknown
		}
	case float32, float64:
		switch reflect.TypeOf(v).Kind() {
		case reflect.Float32, reflect.Float64:
			value := reflect.ValueOf(v).Float()
			if value > 0 {
				return True
			} else if value < 0 {
				return False
			}

			return Unknown
		}
	}

	return any(v).(Trit)
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

	trit := logicToTrit(v)
	*t = trit
	return *t
}

// IsFalse checks if the trit-object is False.
//
// See Trit.IsFalse() for more information.
func IsFalse[T Logicable](t T) bool {
	trit := logicToTrit(t)
	return trit.IsFalse()
}

// IsUnknown checks if the trit-object is Unknown.
//
// See Trit.IsUnknown() for more information.
func IsUnknown[T Logicable](t T) bool {
	trit := logicToTrit(t)
	return trit.IsUnknown()
}

// IsTrue checks if the trit-object is True.
//
// See Trit.IsTrue() for more information.
func IsTrue[T Logicable](t T) bool {
	trit := logicToTrit(t)
	return trit.IsTrue()
}

// Set sets the value of the trit-object.
//
// See Trit.Set() for more information.
func Set[T Logicable](t *Trit, v T) Trit {
	trit := logicToTrit(v)
	*t = trit
	return *t
}

// Convert converts the any Logicable type to Trit.
//
// Example usage:
//
//	t := trit.Convert(true)
//	fmt.Println(t.String()) // Output: True
func Convert[T Logicable](v T) Trit {
	trit := logicToTrit(v)
	return trit
}

// All returns True if all the trit-objects are True.
//
// Example usage:
//
//	t := trit.All(trit.True, trit.True, trit.True)
//	fmt.Println(t.String()) // Output: True
func All[T Logicable](t ...T) Trit {
	for _, v := range t {
		trit := logicToTrit(v)
		if trit.IsFalse() {
			return False
		}
	}

	return True
}

// Any returns True if any of the trit-objects are True.
//
// Example usage:
//
//	t := trit.Any(trit.True, trit.False, trit.False)
//	fmt.Println(t.String()) // Output: True
func Any[T Logicable](t ...T) Trit {
	for _, v := range t {
		trit := logicToTrit(v)
		if trit.IsTrue() {
			return True
		}
	}

	return False
}

// None returns True if none of the trit-objects are True.
//
// Example usage:
//
//	t := trit.None(trit.False, trit.False, trit.False)
//	fmt.Println(t.String()) // Output: True
func None[T Logicable](t ...T) Trit {
	for _, v := range t {
		trit := logicToTrit(v)
		if trit.IsTrue() {
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
	trit := logicToTrit(t)
	return trit.Not()
}

// Ma performs a logical MA (Modus Ponens Absorption) operation
// on a Trit-Like values and returns the result as Trit.
//
// See Trit.Ma() for more information.
func Ma[T Logicable](t T) Trit {
	trit := logicToTrit(t)
	return trit.Ma()
}

// La performs a logical LA (Law of Absorption) operation on a Trit-Like
// value and returns the result as Trit.
//
// See Trit.La() for more information.
func La[T Logicable](t T) Trit {
	trit := logicToTrit(t)
	return trit.La()
}

// Ia performs a logical IA (Idempotent Absorption) operation on a Trit-Like
// value and returns the result as Trit.
//
// See Trit.Ia() for more information.
func Ia[T Logicable](t T) Trit {
	trit := logicToTrit(t)
	return trit.Ia()
}

// And performs a logical AND operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.And() for more information.
func And[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.And(tb)
}

// Or performs a logical OR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Or() for more information.
func Or[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Or(tb)
}

// Xor performs a logical XOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Xor() for more information.
func Xor[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Xor(tb)
}

// Nand performs a logical NAND operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nand() for more information.
func Nand[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nand(tb)
}

// Nor performs a logical NOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nor() for more information.
func Nor[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nor(tb)
}

// Nxor performs a logical NXOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nxor() for more information.
func Nxor[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nxor(tb)
}

// Min performs a logical MIN operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Min() for more information.
func Min[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Min(tb)
}

// Max performs a logical MAX operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Max() for more information.
func Max[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Max(tb)
}
