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
// Truth Tables (T=True, N=Nil, F=False)
//
//	   NOT            AND             OR            XOR
//	 A | B      A | B | C      A | B | C      A | B | C
//	-------    -----------    -----------    -----------
//	 T | F      T | T | T      T | T | T      T | T | F
//	 N | N      T | N | N      T | N | T      T | N | T
//	 F | T      T | F | F      T | F | T      T | F | T
//	            N | T | N      N | T | T      N | T | T
//	            N | N | N      N | N | N      N | N | F
//	            N | F | F      N | F | F      N | F | F
//	            F | T | F      F | T | T      F | T | T
//	            F | N | F      F | N | N      F | N | T
//	            F | F | F      F | F | F      F | F | F
//
//
//	      NAND            NOR           XNOR
//	 A | B | C      A | B | C      A | B | C
//	-----------    -----------    -----------
//	 T | T | F      T | T | F      T | T | T
//	 T | N | T      T | N | F      T | N | F
//	 T | F | T      T | F | F      T | F | F
//	 N | T | T      N | T | F      N | T | F
//	 N | N | T      N | N | T      N | N | T
//	 N | F | T      N | F | F      N | F | T
//	 F | T | T      F | T | F      F | T | F
//	 F | N | T      F | N | T      F | N | T
//	 F | F | T      F | F | T      F | F | T
package trit

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

// DefTrue is a method that checks if the value of the Trit is Nil.
// If it is, it sets the Trit to True.
// It then returns the updated Trit.
//
// Example usage:
//
//	t := trit.Nil
//	t.DefTrue()
//	fmt.Println(t.String()) // Output: True
func (t *Trit) DefTrue() Trit {
	if t.Val() == Nil {
		*t = True
	}
	return *t
}

// DefFalse is a method that checks if the value of the Trit is Nil.
// If it is, it sets the Trit to False.
// It then returns the updated Trit.
//
// Example usage:
//
//	t := trit.Nil
//	t.DefFalse()
//	fmt.Println(t.String()) // Output: False
func (t *Trit) DefFalse() Trit {
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
	if t < 0 {
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
	if t == 0 {
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
	if t > 0 {
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
	if t < 0 {
		return False
	}

	if t > 0 {
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
//	fmt.Println(t.String()) // Output: True
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
//   - Not(True) => False
//   - Not(Nil) => Nil
//   - Not(False) => True
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

// And performs a logical AND operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for AND:
//   - And(True, True) => True
//   - And(True, Nil) => Nil
//   - And(True, False) => False
//   - And(Nil, True) => Nil
//   - And(Nil, Nil) => Nil
//   - And(Nil, False) => False
//   - And(False, True) => False
//   - And(False, Nil) => False
//   - And(False, False) => False
//
// Example usage:
//
//	t1 := trit.True
//	t2 := trit.Nil
//	result := t1.And(t2)
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
//   - Or(True, True) => True
//   - Or(True, Nil) => True
//   - Or(True, False) => True
//   - Or(Nil, True) => True
//   - Or(Nil, Nil) => Nil
//   - Or(Nil, False) => Nil
//   - Or(False, True) => True
//   - Or(False, Nil) => Nil
//   - Or(False, False) => False
//
// Example usage:
//
//	t1 := trit.True
//	t2 := trit.False
//	result := t1.Or(t2)
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
//   - Xor(True, True) => False
//   - Xor(True, Nil) => Nil
//   - Xor(True, False) => True
//   - Xor(Nil, True) => Nil
//   - Xor(Nil, Nil) => False
//   - Xor(Nil, False) => Nil
//   - Xor(False, True) => True
//   - Xor(False, Nil) => Nil
//   - Xor(False, False) => False
//
// Example usage:
//
//	t1 := trit.True
//	t2 := trit.False
//	result := t1.Xor(t2)
//	fmt.Println(result.String()) // Output: True
func (t Trit) Xor(trit Trit) Trit {
	if t.Val() == trit.Val() {
		return False
	}

	if t.Val() == Nil || trit.Val() == Nil {
		return Nil
	}

	return True
}

// Nand performs a logical NAND operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for NAND:
//   - Nand(True, True) => False
//   - Nand(True, Nil) => True
//   - Nand(True, False) => True
//   - Nand(Nil, True) => True
//   - Nand(Nil, Nil) => True
//   - Nand(Nil, False) => True
//   - Nand(False, True) => True
//   - Nand(False, Nil) => True
//   - Nand(False, False) => True
//
// Example usage:
//
//	t1 := trit.True
//	t2 := trit.Nil
//	result := t1.Nand(t2)
//	fmt.Println(result.String()) // Output: True
func (t Trit) Nand(trit Trit) Trit {
	if t.Val() == False || trit.Val() == False {
		return True
	}
	if t.Val() == True && trit.Val() == True {
		return False
	}

	return True
}

// Nor performs a logical NOR operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for NOR:
//   - Nor(True, True) => False
//   - Nor(True, Nil) => False
//   - Nor(True, False) => False
//   - Nor(Nil, True) => False
//   - Nor(Nil, Nil) => True
//   - Nor(Nil, False) => True
//   - Nor(False, True) => False
//   - Nor(False, Nil) => True
//   - Nor(False, False) => True
//
// Example usage:
//
//	t1 := trit.True
//	t2 := trit.False
//	result := t1.Nor(t2)
//	fmt.Println(result.String()) // Output: False
func (t Trit) Nor(trit Trit) Trit {
	// t.Or(trit).Not()
	if t.Val() == True || trit.Val() == True {
		return False
	} else if t.Val() == Nil || trit.Val() == Nil {
		return Nil
	}

	return True
}

// Xnor performs a logical XNOR operation between two Trit values and returns
// the result. This function applies the following rules based on the truth
// table for XNOR:
//   - Xnor(True, True) => True
//   - Xnor(True, Nil) => Nil
//   - Xnor(True, False) => False
//   - Xnor(Nil, True) => Nil
//   - Xnor(Nil, Nil) => True
//   - Xnor(Nil, False) => Nil
//   - Xnor(False, True) => False
//   - Xnor(False, Nil) => Nil
//   - Xnor(False, False) => True
//
// Example usage:
//
//	t1 := trit.True
//	t2 := trit.False
//	result := t1.Xnor(t2)
//	fmt.Println(result.String()) // Output: False
func (t Trit) Xnor(trit Trit) Trit {
	// t.Xor(trit).Not()
	if t.Val() == trit.Val() {
		return True
	} else if t.Val() == Nil || trit.Val() == Nil {
		return Nil
	}

	return False
}
