// Package trit provides three-level logic with the states False, Nil and True.
//
// Trit (short for "trinary digit") is an information unit that can take three
// states, usually expressed as False, Nil, and True. Trit is a fundamental
// unit of trinary or ternary logic systems, including trinary computers and
// balanced ternary systems. This package provides basic logic operations
// including NOT, AND, OR, XOR, NAND, NOR, and XNOR.
//
// The three-level logic (trinary logic) has various applications in computer
// science, particularly in scenarios where a "maybe" or "unknown" state
// is beneficial, such as database systems and logic circuits.
//
// Truth Tables of Three-valued logic
// (T=True, N=Nil, F=False)
//
//	NA   - Not
//	MA   - Modus Ponens Absorption
//	LA   - Law of Absorption
//	IA   - Implication Absorption
//
//	AND  - Logical AND
//	OR   - Logical OR
//	XOR  - Exclusive OR
//
//	NAND - Logical not AND
//	NOR  - Logical not OR
//	NXOR - Logical not XOR
//
//	IMP  - Implication in Lukasevich's Logic
//	MIN  - Minimum
//	MAX  - Maximum
//
//	 A  | NA      A  | MA      A  | LA      A  | IA
//	----+----    ----+----    ----+----    ----+----
//	 F  |  T      F  |  F      F  |  F      F  |  F
//	 N  |  N      N  |  T      N  |  F      N  |  T
//	 T  |  F      T  |  T      T  |  T      T  |  F
//
//
//	 A | B | AND       A | B |  OR       A | B | XOR
//	---+---+------    ---+---+------    ---+---+------
//	 F | F |  F        F | F |  F        F | F |  F
//	 F | N |  F        F | N |  N        F | N |  N
//	 F | T |  F        F | T |  T        F | T |  T
//	 N | F |  F        N | F |  N        N | F |  N
//	 N | N |  N        N | N |  N        N | N |  N
//	 N | T |  N        N | T |  T        N | T |  N
//	 T | F |  F        T | F |  T        T | F |  T
//	 T | N |  N        T | N |  T        T | N |  N
//	 T | T |  T        T | T |  T        T | T |  F
//
//
//	 A | B | NAND      A | B | NOR       A | B | NXOR
//	---+---+------    ---+---+------    ---+---+------
//	 F | F |  T        F | F |  T        F | F |  T
//	 F | N |  T        F | N |  N        F | N |  N
//	 F | T |  T        F | T |  F        F | T |  F
//	 N | F |  T        N | F |  N        N | F |  N
//	 N | N |  N        N | N |  N        N | N |  N
//	 N | T |  N        N | T |  F        N | T |  N
//	 T | F |  T        T | F |  F        T | F |  F
//	 T | N |  N        T | N |  F        T | N |  N
//	 T | T |  F        T | T |  F        T | T |  T
//
//
//	 A | B | IMP       A | B | MIN       A | B | MAX
//	---+---+------    ---+---+------    ---+---+------
//	 F | F |  T        F | F |  F        F | F |  F
//	 F | N |  T        F | N |  F        F | N |  N
//	 F | T |  T        F | T |  F        F | T |  T
//	 N | F |  N        N | F |  F        N | F |  N
//	 N | N |  T        N | N |  N        N | N |  N
//	 N | T |  T        N | T |  N        N | T |  T
//	 T | F |  F        T | F |  F        T | F |  T
//	 T | N |  N        T | N |  N        T | N |  T
//	 T | T |  T        T | T |  T        T | T |  T
package trit

import "reflect"

// Trit represents a trinary digit, which can take on three distinct
// states: False, Nil, or True. This type is a fundamental unit of
// trinary or ternary logic systems, including trinary computers and
// balanced ternary systems.
type Trit int8

const (
	// False represents the Trit state equivalent to 'false' in binary
	// logic. It is represented by a negative value (-1) in the underlying
	// int8 type.
	//
	// *Any negative numbers are also considered False.
	False Trit = -1

	// Nil represents an unknown or indeterminate Trit state. It is
	// beneficial in scenarios where a "maybe" state is required.
	// Nil is represented by a zero value in the underlying int8 type.
	Nil Trit = 0

	// True represents the Trit state equivalent to 'true' in binary
	// logic. It is represented by a positive value (1) in the underlying
	// int8 type.
	//
	// *Any positive numbers are also considered True.
	True Trit = 1
)

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

// Def is a method that checks if the value of the Trit is Nil.
// If it is, it sets the Trit to the given Trit argument.
//
// Example usage:
//
//	t := trit.Nil
//	t.Def(trit.True)
//	fmt.Println(t.String()) // Output: True
func (t *Trit) Def(trit Trit) Trit {
	if t.Val() == Nil {
		*t = trit
	}

	return *t
}

// TrueIfNil is a method that checks if the value of the Trit is Nil.
// If it is, it sets the Trit to True.
// It then returns the updated Trit.
//
// Example usage:
//
//	t := trit.Nil
//	t.TrueIfNil()
//	fmt.Println(t.String()) // Output: True
func (t *Trit) TrueIfNil() Trit {
	if t.Val() == Nil {
		*t = True
	}

	return *t
}

// FalseIfNil is a method that checks if the value of the Trit is Nil.
// If it is, it sets the Trit to False.
// It then returns the updated Trit.
//
// Example usage:
//
//	t := trit.Nil
//	t.FalseIfNil()
//	fmt.Println(t.String()) // Output: False
func (t *Trit) FalseIfNil() Trit {
	if t.Val() == Nil {
		*t = False
	}

	return *t
}

// Clean is a method that checks if the value of the Trit is Nil.
// If it is, it resets the Trit to Nil.
// It then returns the updated Trit.
//
// Example usage:
//
//	t := trit.True
//	t.Clean()
//	fmt.Println(t.String()) // Output: Nil
func (t *Trit) Clean() Trit {
	if t.Val() == Nil {
		*t = Nil
	}

	return *t
}

// IsFalse returns true if the Trit value represents a False state, which
// is any negative number in the underlying int8 type.
//
// Example usage:
//
//	t := trit.Trit(-2)
//	fmt.Println(t.IsFalse()) // Output: true
func (t Trit) IsFalse() bool {
	if int8(t) < 0 {
		return true
	}

	return false
}

// IsNil returns true if the Trit value represents a Nil state, which
// is zero in the underlying int8 type.
//
// Example usage:
//
//	t := trit.Trit(0)
//	fmt.Println(t.IsNil()) // Output: true
func (t Trit) IsNil() bool {
	if int8(t) == 0 {
		return true
	}

	return false
}

// IsTrue returns true if the Trit value represents a True state, which
// is any positive number in the underlying int8 type.
//
// Example usage:
//
//	t := trit.Trit(2)
//	fmt.Println(t.IsTrue()) // Output: true
func (t Trit) IsTrue() bool {
	if int8(t) > 0 {
		return true
	}

	return false
}

// Set assigns a Trit value based on the given integer. Negative values
// are interpreted as False, zero as Nil, and positive values as True.
//
// Example usage:
//
//	var t trit.Trit
//	t.Set(-2)
//	fmt.Println(t.String()) // Output: False
func (t *Trit) Set(v int) Trit {
	switch {
	case v < 0:
		*t = False
	case v > 0:
		*t = True
	default:
		*t = Nil
	}

	return *t
}

// Val returns the normalized Trit value if a Trit was not properly
// initialized using the predefined constants (False, Nil, True).
// If a Trit was set using an int8 value other than -1, 0, or 1, Val
// maps it to the closest Trit value.
//
// Example usage:
//
//	t := trit.Trit(7)
//	fmt.Println(t.Val().String()) // Output: True
func (t Trit) Val() Trit {
	if t.IsFalse() {
		return False
	}

	if t.IsTrue() {
		return True
	}

	return Nil
}

// Norm normalizes the Trit value. If a Trit was set using an int8
// value other than -1, 0, or 1, Norm maps it to the closest Trit value.
//
// Example usage:
//
//	t := trit.Trit(7)
//	t.Norm()
//	fmt.Println(t.Int()) // Output: 1
func (t *Trit) Norm() Trit {
	*t = t.Val()
	return *t
}

// Int returns the underlying int8 value of a Trit in the form of an int.
// It first normalizes the Trit value and then returns it as an int.
//
// Example usage:
//
//	t := trit.Trit(7)
//	fmt.Println(t.Int()) // Output: 1
func (t Trit) Int() int {
	return int(t.Val())
}

// String returns a string representation of a Trit value.
// The possible return values are "False", "Nil", and "True".
//
// Example usage:
//
//	t := trit.True
//	fmt.Println(t.String()) // Output: True
func (t Trit) String() string {
	switch t.Val() {
	case False:
		return "False"
	case True:
		return "True"
	}

	return "Nil"
}

// Not performs a logical NOT operation on a Trit value and returns the result.
// This function applies the following rules based on the truth table for NOT:
//   - Not(False) => True
//   - Not(Nil)   => Nil
//   - Not(True)  => False
//
// Example usage:
//
//	t := trit.True
//	result := t.Not()
//	fmt.Println(result.String()) // Output: False
func (t Trit) Not() Trit {
	switch t.Val() {
	case False:
		return True
	case True:
		return False
	}

	return Nil
}

// Ma performs a logical MA (Modus Ponens Absorption) operation on a Trit
// value and returns the result. This function applies the following rules
// based on the truth table for MA:
//   - Ma(False) => False
//   - Ma(Nil)   => True
//   - Ma(True)  => True
//
// Example usage:
//
//	a := trit.True
//	result := a.Ma()
//	fmt.Println(result.String()) // Output: True
func (t Trit) Ma() Trit {
	if t.Val() == False {
		return False
	}

	return True
}

// La performs a logical LA (Law of Absorption) operation on a Trit value
// and returns the result. This function applies the following rules based
// on the truth table for LA:
//   - La(False) => False
//   - La(Nil)   => False
//   - La(True)  => True
//
// Example usage:
//
//	a := trit.True
//	result := a.La()
//	fmt.Println(result.String()) // Output: True
func (t Trit) La() Trit {
	if t.Val() == True {
		return True
	}

	return False
}

// Ia performs a logical IA (Implication Absorption) operation on a Trit
// values and returns the result. This function applies the following
// rules based on the truth table for IA:
//   - Ia(False) => False
//   - Ia(Nil)   => True
//   - Ia(True)  => False
//
// Example usage:
//
//	a := trit.True
//	result := a.Ia()
//	fmt.Println(result.String()) // Output: False
func (t Trit) Ia() Trit {
	if t.Val() == Nil {
		return True
	}

	return False
}

// And performs a logical AND operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for AND:
//   - And(False, False) => False
//   - And(False, Nil)   => False
//   - And(False, True)  => False
//   - And(Nil, False)   => False
//   - And(Nil, Nil)     => Nil
//   - And(Nil, True)    => Nil
//   - And(True, False)  => False
//   - And(True, Nil)    => Nil
//   - And(True, True)   => True
//
// Example usage:
//
//	a := trit.True
//	b := trit.Nil
//	result := a.And(b)
//	fmt.Println(result.String()) // Output: Nil
func (t Trit) And(trit Trit) Trit {
	if t.Val() == False || trit.Val() == False {
		return False
	}

	if t.Val() == Nil || trit.Val() == Nil {
		return Nil
	}

	return True
}

// Or performs a logical OR operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for OR:
//   - Or(False, False) => False
//   - Or(False, Nil)   => Nil
//   - Or(False, True)  => True
//   - Or(Nil, False)   => Nil
//   - Or(Nil, Nil)     => Nil
//   - Or(Nil, True)    => True
//   - Or(True, False)  => True
//   - Or(True, Nil)    => True
//   - Or(True, True)   => True
//
// Example usage:
//
//	a := trit.True
//	b := trit.False
//	result := a.Or(b)
//	fmt.Println(result.String()) // Output: True
func (t Trit) Or(trit Trit) Trit {
	if t.Val() == True || trit.Val() == True {
		return True
	}

	if t.Val() == Nil || trit.Val() == Nil {
		return Nil
	}

	return False
}

// Xor performs a logical XOR operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for XOR:
//   - Xor(False, False) => False
//   - Xor(False, Nil)   => Nil
//   - Xor(False, True)  => True
//   - Xor(Nil, False)   => Nil
//   - Xor(Nil, Nil)     => Nil
//   - Xor(Nil, True)    => Nil
//   - Xor(True, False)  => True
//   - Xor(True, Nil)    => Nil
//   - Xor(True, True)   => False
//
// Example usage:
//
//	a := trit.True
//	b := trit.False
//	result := a.Xor(b)
//	fmt.Println(result.String()) // Output: True
func (t Trit) Xor(trit Trit) Trit {
	// Check first, because Xor(Nil, Nil) should be Nil.
	if t.Val() == Nil || trit.Val() == Nil {
		return Nil
	}

	// Pay attention, Nil == Nil != False
	if t.Val() == trit.Val() {
		return False
	}

	return True
}

// Nand performs a logical NAND operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for NAND:
//   - Nand(False, False) => True
//   - Nand(False, Nil)   => True
//   - Nand(False, True)  => True
//   - Nand(Nil, False)   => True
//   - Nand(Nil, Nil)     => Nil
//   - Nand(Nil, True)    => Nil
//   - Nand(True, False)  => True
//   - Nand(True, Nil)    => Nil
//   - Nand(True, True)   => False
//
// Example usage:
//
//	a := trit.True
//	b := trit.Nil
//	result := a.Nand(b)
//	fmt.Println(result.String()) // Output: True
func (t Trit) Nand(trit Trit) Trit {
	return t.And(trit).Not()
}

// Nor performs a logical NOR operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for NOR:
//   - Nor(False, False) => True
//   - Nor(False, Nil)   => Nil
//   - Nor(False, True)  => False
//   - Nor(Nil, False)   => Nil
//   - Nor(Nil, Nil)     => Nil
//   - Nor(Nil, True)    => False
//   - Nor(True, False)  => False
//   - Nor(True, Nil)    => False
//   - Nor(True, True)   => False
//
// Example usage:
//
//	a := trit.True
//	b := trit.False
//	result := a.Nor(b)
//	fmt.Println(result.String()) // Output: False
func (t Trit) Nor(trit Trit) Trit {
	return t.Or(trit).Not()
}

// Nxor performs a logical XNOR operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for XNOR:
//   - Nxor(False, False) => True
//   - Nxor(False, Nil)   => Nil
//   - Nxor(False, True)  => False
//   - Nxor(Nil, False)   => Nil
//   - Nxor(Nil, Nil)     => Nil
//   - Nxor(Nil, True)    => Nil
//   - Nxor(True, False)  => False
//   - Nxor(True, Nil)    => Nil
//   - Nxor(True, True)  => True
//
// Example usage:
//
//	a := trit.True
//	b := trit.False
//	result := a.Nxor(b)
//	fmt.Println(result.String()) // Output: False
func (t Trit) Nxor(trit Trit) Trit {
	return t.Xor(trit).Not()
}

// Min performs a logical MIN operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for MIN:
//   - Min(False, False) => False
//   - Min(False, Nil)   => False
//   - Min(False, True)  => False
//   - Min(Nil, False)   => False
//   - Min(Nil, Nil)     => Nil
//   - Min(Nil, True)    => Nil
//   - Min(True, False)  => False
//   - Min(True, Nil)    => Nil
//   - Min(True, True)   => True
//
// Example usage:
//
//	a := trit.True
//	b := trit.False
//	result := a.Min(b)
//	fmt.Println(result.String()) // Output: False
func (t Trit) Min(trit Trit) Trit {
	return t.And(trit)
}

// Max performs a logical MAX operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for MAX:
//   - Max(False, False) => False
//   - Max(False, Nil)   => Nil
//   - Max(False, True)  => True
//   - Max(Nil, False)   => Nil
//   - Max(Nil, Nil)     => Nil
//   - Max(Nil, True)    => True
//   - Max(True, False)  => True
//   - Max(True, Nil)    => True
//   - Max(True, True)   => True
//
// Example usage:
//
//	a := trit.True
//	b := trit.False
//	result := a.Max(b)
//	fmt.Println(result.String()) // Output: True
func (t Trit) Max(trit Trit) Trit {
	return t.Or(trit)
}

// Imp performs a logical IMP operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for IMP:
//   - Imp(False, False) => True
//   - Imp(False, Nil)   => True
//   - Imp(False, True)  => True
//   - Imp(Nil, False)   => Nil
//   - Imp(Nil, Nil)     => True
//   - Imp(Nil, True)    => True
//   - Imp(True, False)  => False
//   - Imp(True, Nil)    => Nil
//   - Imp(True, True)   => True
//
// Example usage:
//
//	a := trit.True
//	b := trit.False
//	result := a.Imp(b)
//	fmt.Println(result.String()) // Output: False
func (t Trit) Imp(trit Trit) Trit {
	if t.Val() == Nil && trit.Val() == Nil {
		return True
	} else if t.Val() == False || trit.Val() == True {
		return True
	} else if t.Val() == Nil || trit.Val() == Nil {
		return Nil
	}

	return False
}
