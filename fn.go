package trit

import "reflect"

// Logic is a special data type from which to determine the state of trit.
type Logic interface {
	bool | int | int8 | int16 | int32 | int64 | Trit
}

// The logicToTrit function converts any logic type to Trit object.
func logicToTrit[T Logic](v T) Trit {
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
			intValue := reflect.ValueOf(v).Int()
			if intValue > 0 {
				return True
			} else if intValue < 0 {
				return False
			}

			return Nil
		}
	case Trit:
		return any(v).(Trit)
	}

	// The event will never fire because the function cannot work with
	// types other than Logic. Therefore, this block cannot be tested.
	return Nil
}

// Default sets the default value for the trit-object
// if this one has a Nil state.
//
// Example usage:
//
//	t := trit.Nil
//	trit.Default(&t, trit.True)
//	fmt.Println(t.String()) // Output: True
func Default[T Logic](t *Trit, v T) Trit {
	// If the trit is not Nil, return the trit.
	if t.Val() != Nil {
		return *t
	}

	trit := logicToTrit(v)
	*t = trit
	return *t
}

// Not performs a logical NOT operation on a Trit-Like value and
// returns the result as Trit.
//
// See Trit.Not() for more information.
func Not[T Logic](t T) Trit {
	trit := logicToTrit(t)
	return trit.Not()
}

// Ma performs a logical MA (Modus Ponens Absorption) operation
// on a Trit-Like values and returns the result as Trit.
//
// See Trit.Ma() for more information.
func Ma[T Logic](t T) Trit {
	trit := logicToTrit(t)
	return trit.Ma()
}

// La performs a logical LA (Law of Absorption) operation on a Trit-Like
// value and returns the result as Trit.
//
// See Trit.La() for more information.
func La[T Logic](t T) Trit {
	trit := logicToTrit(t)
	return trit.La()
}

// Ia performs a logical IA (Idempotent Absorption) operation on a Trit-Like
// value and returns the result as Trit.
//
// See Trit.Ia() for more information.
func Ia[T Logic](t T) Trit {
	trit := logicToTrit(t)
	return trit.Ia()
}

// And performs a logical AND operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.And() for more information.
func And[T, U Logic](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.And(tb)
}

// Or performs a logical OR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Or() for more information.
func Or[T, U Logic](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Or(tb)
}

// Xor performs a logical XOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Xor() for more information.
func Xor[T, U Logic](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Xor(tb)
}

// Nand performs a logical NAND operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nand() for more information.
func Nand[T, U Logic](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nand(tb)
}

// Nor performs a logical NOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nor() for more information.
func Nor[T, U Logic](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nor(tb)
}

// Nxor performs a logical NXOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nxor() for more information.
func Nxor[T, U Logic](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nxor(tb)
}

// Min performs a logical MIN operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Min() for more information.
func Min[T, U Logic](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Min(tb)
}

// Max performs a logical MAX operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Max() for more information.
func Max[T, U Logic](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Max(tb)
}
