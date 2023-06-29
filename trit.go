// Package trit provides three-level logic with the states False,
// Unknown and True.
//
// Trit (short for "trinary digit") is an information unit that can take three
// states, usually expressed as False, Unknown, and True. Trit is a fundamental
// unit of trinary or ternary logic systems, including trinary computers and
// balanced ternary systems. This package provides basic logic operations
// including NOT, AND, OR, XOR, NAND, NOR, and XNOR.
//
// The three-level logic (trinary logic) has various applications in computer
// science, particularly in scenarios where a "maybe" or "unknown" state
// is beneficial, such as database systems and logic circuits.
//
// Truth Tables of Three-valued logic
// (T=True, U=Unknown, F=False)
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
//	 U  |  U      U  |  T      U  |  F      U  |  T
//	 T  |  F      T  |  T      T  |  T      T  |  F
//
//
//	 A | B | AND       A | B |  OR       A | B | XOR
//	---+---+------    ---+---+------    ---+---+------
//	 F | F |  F        F | F |  F        F | F |  F
//	 F | U |  F        F | U |  U        F | U |  U
//	 F | T |  F        F | T |  T        F | T |  T
//	 U | F |  F        U | F |  U        U | F |  U
//	 U | U |  U        U | U |  U        U | U |  U
//	 U | T |  U        U | T |  T        U | T |  U
//	 T | F |  F        T | F |  T        T | F |  T
//	 T | U |  U        T | U |  T        T | U |  U
//	 T | T |  T        T | T |  T        T | T |  F
//
//
//	 A | B | NAND      A | B | NOR       A | B | NXOR
//	---+---+------    ---+---+------    ---+---+------
//	 F | F |  T        F | F |  T        F | F |  T
//	 F | U |  T        F | U |  U        F | U |  U
//	 F | T |  T        F | T |  F        F | T |  F
//	 U | F |  T        U | F |  U        U | F |  U
//	 U | U |  U        U | U |  U        U | U |  U
//	 U | T |  U        U | T |  F        U | T |  U
//	 T | F |  T        T | F |  F        T | F |  F
//	 T | U |  U        T | U |  F        T | U |  U
//	 T | T |  F        T | T |  F        T | T |  T
//
//
//	 A | B | IMP       A | B | MIN       A | B | MAX
//	---+---+------    ---+---+------    ---+---+------
//	 F | F |  T        F | F |  F        F | F |  F
//	 F | U |  T        F | U |  F        F | U |  U
//	 F | T |  T        F | T |  F        F | T |  T
//	 U | F |  U        U | F |  F        U | F |  U
//	 U | U |  T        U | U |  U        U | U |  U
//	 U | T |  T        U | T |  U        U | T |  T
//	 T | F |  F        T | F |  F        T | F |  T
//	 T | U |  U        T | U |  U        T | U |  T
//	 T | T |  T        T | T |  T        T | T |  T
package trit

// Trit represents a trinary digit, which can take on three distinct
// states: False, Unknown, or True. This type is a fundamental unit of
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

	// Unknown represents an unknown or indeterminate Trit state. It is
	// beneficial in scenarios where a "maybe" state is required.
	// Unknown is represented by a zero value in the underlying int8 type.
	Unknown Trit = 0

	// True represents the Trit state equivalent to 'true' in binary
	// logic. It is represented by a positive value (1) in the underlying
	// int8 type.
	//
	// *Any positive numbers are also considered True.
	True Trit = 1
)

// Default is a method that checks if the value of the Trit is Unknown.
// If it is, it sets the Trit to the given Trit argument.
//
// Example usage:
//
//	t := trit.Unknown
//	t.Default(trit.True)
//	fmt.Println(t.String()) // Output: True
func (t *Trit) Default(trit Trit) Trit {
	if t.Val() == Unknown {
		*t = trit
	}

	return *t
}

// TrueIfUnknown is a method that checks if the value of the Trit is Unknown.
// If it is, it sets the Trit to True.
// It then returns the updated Trit.
//
// Example usage:
//
//	t := trit.Unknown
//	t.TrueIfUnknown()
//	fmt.Println(t.String()) // Output: True
func (t *Trit) TrueIfUnknown() Trit {
	if t.Val() == Unknown {
		*t = True
	}

	return *t
}

// FalseIfUnknown is a method that checks if the value of the Trit is Unknown.
// If it is, it sets the Trit to False.
// It then returns the updated Trit.
//
// Example usage:
//
//	t := trit.Unknown
//	t.FalseIfUnknown()
//	fmt.Println(t.String()) // Output: False
func (t *Trit) FalseIfUnknown() Trit {
	if t.Val() == Unknown {
		*t = False
	}

	return *t
}

// Clean is a method that checks if the value of the Trit is Unknown.
// If it is, it resets the Trit to Unknown.
// It then returns the updated Trit.
//
// Example usage:
//
//	t := trit.True
//	t.Clean()
//	fmt.Println(t.String()) // Output: Unknown
func (t *Trit) Clean() Trit {
	if t.Val() == Unknown {
		*t = Unknown
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

// IsUnknown returns true if the Trit value represents a Unknown state, which
// is zero in the underlying int8 type.
//
// Example usage:
//
//	t := trit.Trit(0)
//	fmt.Println(t.IsUnknown()) // Output: true
func (t Trit) IsUnknown() bool {
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
// are interpreted as False, zero as Unknown, and positive values as True.
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
		*t = Unknown
	}

	return *t
}

// Val returns the normalized Trit value if a Trit was not properly
// initialized using the predefined constants (False, Unknown, True).
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

	return Unknown
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
// The possible return values are "False", "Unknown", and "True".
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

	return "Unknown"
}

// Not performs a logical NOT operation on a Trit value and returns the result.
// This function applies the following rules based on the truth table for NOT:
//   - Not(False) => True
//   - Not(Unknown)   => Unknown
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

	return Unknown
}

// Ma performs a logical MA (Modus Ponens Absorption) operation on a Trit
// value and returns the result. This function applies the following rules
// based on the truth table for MA:
//   - Ma(False) => False
//   - Ma(Unknown)   => True
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
//   - La(Unknown)   => False
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
//   - Ia(Unknown)   => True
//   - Ia(True)  => False
//
// Example usage:
//
//	a := trit.True
//	result := a.Ia()
//	fmt.Println(result.String()) // Output: False
func (t Trit) Ia() Trit {
	if t.Val() == Unknown {
		return True
	}

	return False
}

// And performs a logical AND operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for AND:
//   - And(False, False) => False
//   - And(False, Unknown)   => False
//   - And(False, True)  => False
//   - And(Unknown, False)   => False
//   - And(Unknown, Unknown)     => Unknown
//   - And(Unknown, True)    => Unknown
//   - And(True, False)  => False
//   - And(True, Unknown)    => Unknown
//   - And(True, True)   => True
//
// Example usage:
//
//	a := trit.True
//	b := trit.Unknown
//	result := a.And(b)
//	fmt.Println(result.String()) // Output: Unknown
func (t Trit) And(trit Trit) Trit {
	if t.Val() == False || trit.Val() == False {
		return False
	}

	if t.Val() == Unknown || trit.Val() == Unknown {
		return Unknown
	}

	return True
}

// Or performs a logical OR operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for OR:
//   - Or(False, False) => False
//   - Or(False, Unknown)   => Unknown
//   - Or(False, True)  => True
//   - Or(Unknown, False)   => Unknown
//   - Or(Unknown, Unknown)     => Unknown
//   - Or(Unknown, True)    => True
//   - Or(True, False)  => True
//   - Or(True, Unknown)    => True
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

	if t.Val() == Unknown || trit.Val() == Unknown {
		return Unknown
	}

	return False
}

// Xor performs a logical XOR operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for XOR:
//   - Xor(False, False) => False
//   - Xor(False, Unknown)   => Unknown
//   - Xor(False, True)  => True
//   - Xor(Unknown, False)   => Unknown
//   - Xor(Unknown, Unknown)     => Unknown
//   - Xor(Unknown, True)    => Unknown
//   - Xor(True, False)  => True
//   - Xor(True, Unknown)    => Unknown
//   - Xor(True, True)   => False
//
// Example usage:
//
//	a := trit.True
//	b := trit.False
//	result := a.Xor(b)
//	fmt.Println(result.String()) // Output: True
func (t Trit) Xor(trit Trit) Trit {
	// Check first, because Xor(Unknown, Unknown) should be Unknown.
	if t.Val() == Unknown || trit.Val() == Unknown {
		return Unknown
	}

	// Pay attention, Unknown == Unknown != False
	if t.Val() == trit.Val() {
		return False
	}

	return True
}

// Nand performs a logical NAND operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for NAND:
//   - Nand(False, False) => True
//   - Nand(False, Unknown)   => True
//   - Nand(False, True)  => True
//   - Nand(Unknown, False)   => True
//   - Nand(Unknown, Unknown)     => Unknown
//   - Nand(Unknown, True)    => Unknown
//   - Nand(True, False)  => True
//   - Nand(True, Unknown)    => Unknown
//   - Nand(True, True)   => False
//
// Example usage:
//
//	a := trit.True
//	b := trit.Unknown
//	result := a.Nand(b)
//	fmt.Println(result.String()) // Output: True
func (t Trit) Nand(trit Trit) Trit {
	return t.And(trit).Not()
}

// Nor performs a logical NOR operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for NOR:
//   - Nor(False, False) => True
//   - Nor(False, Unknown)   => Unknown
//   - Nor(False, True)  => False
//   - Nor(Unknown, False)   => Unknown
//   - Nor(Unknown, Unknown)     => Unknown
//   - Nor(Unknown, True)    => False
//   - Nor(True, False)  => False
//   - Nor(True, Unknown)    => False
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
//   - Nxor(False, Unknown)   => Unknown
//   - Nxor(False, True)  => False
//   - Nxor(Unknown, False)   => Unknown
//   - Nxor(Unknown, Unknown)     => Unknown
//   - Nxor(Unknown, True)    => Unknown
//   - Nxor(True, False)  => False
//   - Nxor(True, Unknown)    => Unknown
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
//   - Min(False, Unknown)   => False
//   - Min(False, True)  => False
//   - Min(Unknown, False)   => False
//   - Min(Unknown, Unknown)     => Unknown
//   - Min(Unknown, True)    => Unknown
//   - Min(True, False)  => False
//   - Min(True, Unknown)    => Unknown
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
//   - Max(False, Unknown)   => Unknown
//   - Max(False, True)  => True
//   - Max(Unknown, False)   => Unknown
//   - Max(Unknown, Unknown)     => Unknown
//   - Max(Unknown, True)    => True
//   - Max(True, False)  => True
//   - Max(True, Unknown)    => True
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
//   - Imp(False, Unknown)   => True
//   - Imp(False, True)  => True
//   - Imp(Unknown, False)   => Unknown
//   - Imp(Unknown, Unknown)     => True
//   - Imp(Unknown, True)    => True
//   - Imp(True, False)  => False
//   - Imp(True, Unknown)    => Unknown
//   - Imp(True, True)   => True
//
// Example usage:
//
//	a := trit.True
//	b := trit.False
//	result := a.Imp(b)
//	fmt.Println(result.String()) // Output: False
func (t Trit) Imp(trit Trit) Trit {
	if t.Val() == Unknown && trit.Val() == Unknown {
		return True
	} else if t.Val() == False || trit.Val() == True {
		return True
	} else if t.Val() == Unknown || trit.Val() == Unknown {
		return Unknown
	}

	return False
}
