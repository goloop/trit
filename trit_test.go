package trit

import "testing"

// TestMethodDefault tests the Default method.
func TestMethodDefault(t *testing.T) {
	t.Run("Def should update Unknown to True", func(t *testing.T) {
		t1 := Unknown
		result := t1.Default(True)
		if result != True {
			t.Errorf("Did not update Unknown to True")
		}
	})

	t.Run("Def should update Unknown to False", func(t *testing.T) {
		t1 := Unknown
		result := t1.Default(False)
		if result != False {
			t.Errorf("Def did not update Unknown to False")
		}
	})

	t.Run("Def should not update non-Unknown Trit", func(t *testing.T) {
		t1 := True
		result := t1.Default(False)
		if result != True {
			t.Errorf("Def updated non-Unknown Trit")
		}
	})
}

// TestTrueIfUnknown tests the TrueIfUnknown method.
func TestTrueIfUnknown(t *testing.T) {
	t.Run("Should set Unknown to True", func(t *testing.T) {
		tr := Unknown
		tr.TrueIfUnknown()
		if tr != True {
			t.Errorf("TrueIfUnknown did not set Unknown to True")
		}
	})

	t.Run("Should not change True to False", func(t *testing.T) {
		tr := True
		tr.TrueIfUnknown()
		if tr != True {
			t.Errorf("TrueIfUnknown changed True to False")
		}
	})

	t.Run("Should not change False to True", func(t *testing.T) {
		tr := False
		tr.TrueIfUnknown()
		if tr != False {
			t.Errorf("TrueIfUnknown changed False to True")
		}
	})
}

// TestFalseIfUnknown tests the FalseIfUnknown method.
func TestFalseIfUnknown(t *testing.T) {
	t.Run("Should set Unknown to False", func(t *testing.T) {
		tr := Unknown
		tr.FalseIfUnknown()
		if tr != False {
			t.Errorf("FalseIfUnknown did not set Unknown to False")
		}
	})

	t.Run("Should not change False to True", func(t *testing.T) {
		tr := False
		tr.FalseIfUnknown()
		if tr != False {
			t.Errorf("FalseIfUnknown changed False to True")
		}
	})

	t.Run("Should not change True to False", func(t *testing.T) {
		tr := True
		tr.FalseIfUnknown()
		if tr != True {
			t.Errorf("FalseIfUnknown changed True to False")
		}
	})
}

// TestClean tests the Clean method.
func TestClean(t *testing.T) {
	t.Run("Clean should set Unknown to Unknown", func(t *testing.T) {
		tr := Unknown
		tr.Clean()
		if tr != Unknown {
			t.Errorf("Clean did not set Unknown to Unknown")
		}
	})

	t.Run("Clean should not change False to Unknown", func(t *testing.T) {
		tr := False
		tr.Clean()
		if tr != False {
			t.Errorf("Clean changed False to Unknown")
		}
	})

	t.Run("Clean should not change True to Unknown", func(t *testing.T) {
		tr := True
		tr.Clean()
		if tr != True {
			t.Errorf("Clean changed True to Unknown")
		}
	})
}

// TestMethodIsFalse tests the IsFalse method.
func TestMethodIsFalse(t *testing.T) {
	t.Run("IsFalse should return true for False", func(t *testing.T) {
		tr := False
		if !tr.IsFalse() {
			t.Errorf("IsFalse returned false for False")
		}
	})

	t.Run("IsFalse should return false for Unknown", func(t *testing.T) {
		tr := Unknown
		if tr.IsFalse() {
			t.Errorf("IsFalse returned true for Unknown")
		}
	})

	t.Run("IsFalse should return false for True", func(t *testing.T) {
		tr := True
		if tr.IsFalse() {
			t.Errorf("IsFalse returned true for True")
		}
	})
}

// TestMethodIsUnknown tests the IsUnknown method.
func TestMethodIsUnknown(t *testing.T) {
	t.Run("IsUnknown should return false for False", func(t *testing.T) {
		tr := False
		if tr.IsUnknown() {
			t.Errorf("IsUnknown returned true for False")
		}
	})

	t.Run("IsUnknown should return true for Unknown", func(t *testing.T) {
		tr := Unknown
		if !tr.IsUnknown() {
			t.Errorf("IsUnknown returned false for Unknown")
		}
	})

	t.Run("IsUnknown should return false for True", func(t *testing.T) {
		tr := True
		if tr.IsUnknown() {
			t.Errorf("IsUnknown returned true for True")
		}
	})
}

// TestMethodIsTrue tests the IsTrue method.
func TestMethodIsTrue(t *testing.T) {
	t.Run("IsTrue should return false for False", func(t *testing.T) {
		tr := False
		if tr.IsTrue() {
			t.Errorf("IsTrue returned true for False")
		}
	})

	t.Run("IsTrue should return false for Unknown", func(t *testing.T) {
		tr := Unknown
		if tr.IsTrue() {
			t.Errorf("IsTrue returned true for Unknown")
		}
	})

	t.Run("IsTrue should return true for True", func(t *testing.T) {
		tr := True
		if !tr.IsTrue() {
			t.Errorf("IsTrue returned false for True")
		}
	})
}

// TestMethodSet tests the Set method.
func TestMethodSet(t *testing.T) {
	t.Run("Set value to False for negative integer", func(t *testing.T) {
		tr := Trit(0)
		tr.Set(-2)
		if tr != False {
			t.Errorf("Set did not set value to False for negative integer")
		}
	})

	t.Run("Set should set value to Unknown for zero", func(t *testing.T) {
		tr := Trit(1)
		tr.Set(0)
		if tr != Unknown {
			t.Errorf("Set did not set value to Unknown for zero")
		}
	})

	t.Run("Set value to True for positive integer", func(t *testing.T) {
		tr := Trit(0)
		tr.Set(2)
		if tr != True {
			t.Errorf("Set did not set value to True for positive integer")
		}
	})
}

// TestVal tests the Val method.
func TestVal(t *testing.T) {
	t.Run("Val should return False for negative Trit", func(t *testing.T) {
		tr := Trit(-2)
		if tr.Val() != False {
			t.Errorf("Val did not return False for negative Trit")
		}
	})

	t.Run("Val should return Unknown for zero Trit", func(t *testing.T) {
		tr := Trit(0)
		if tr.Val() != Unknown {
			t.Errorf("Val did not return Unknown for zero Trit")
		}
	})

	t.Run("Val should return True for positive Trit", func(t *testing.T) {
		tr := Trit(2)
		if tr.Val() != True {
			t.Errorf("Val did not return True for positive Trit")
		}
	})
}

// TestNorm tests the Norm method.
func TestNorm(t *testing.T) {
	t.Run("Norm should normalize False to False", func(t *testing.T) {
		tr := False
		tr.Norm()
		if tr != False {
			t.Errorf("Norm did not normalize False to False")
		}
	})

	t.Run("Norm should normalize Unknown to Unknown", func(t *testing.T) {
		tr := Unknown
		tr.Norm()
		if tr != Unknown {
			t.Errorf("Norm did not normalize Unknown to Unknown")
		}
	})

	t.Run("Norm should normalize True to True", func(t *testing.T) {
		tr := True
		tr.Norm()
		if tr != True {
			t.Errorf("Norm did not normalize True to True")
		}
	})
}

// TestInt tests the Int method.
func TestInt(t *testing.T) {
	t.Run("Int should return -1 for False", func(t *testing.T) {
		tr := False
		if tr.Int() != -1 {
			t.Errorf("Int did not return -1 for False")
		}
	})

	t.Run("Int should return 0 for Unknown", func(t *testing.T) {
		tr := Unknown
		if tr.Int() != 0 {
			t.Errorf("Int did not return 0 for Unknown")
		}
	})

	t.Run("Int should return 1 for True", func(t *testing.T) {
		tr := True
		if tr.Int() != 1 {
			t.Errorf("Int did not return 1 for True")
		}
	})
}

// TestString tests the String method.
func TestString(t *testing.T) {
	t.Run("String should return 'False' for False", func(t *testing.T) {
		tr := False
		if tr.String() != "False" {
			t.Errorf("String did not return 'False' for False")
		}
	})

	t.Run("String should return 'Unknown' for Unknown", func(t *testing.T) {
		tr := Unknown
		if tr.String() != "Unknown" {
			t.Errorf("String did not return 'Unknown' for Unknown")
		}
	})

	t.Run("String should return 'True' for True", func(t *testing.T) {
		tr := True
		if tr.String() != "True" {
			t.Errorf("String did not return 'True' for True")
		}
	})
}

// TestMethodNot tests the Not method.
func TestMethodNot(t *testing.T) {
	tests := []struct {
		name string
		in   Trit
		out  Trit
	}{
		{"Not should return True for False", False, True},
		{"Not should return Unknown for Unknown", Unknown, Unknown},
		{"Not should return False for True", True, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.in.Not()
			if result != test.out {
				t.Errorf("Not did not return %v for %v", test.out, test.in)
			}
		})
	}
}

// TestMethodMa tests the Ma method.
func TestMethodMa(t *testing.T) {
	tests := []struct {
		name string
		in   Trit
		out  Trit
	}{
		{"Ma should return False for False", False, False},
		{"Ma should return True for Unknown", Unknown, True},
		{"Ma should return True for True", True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.in.Ma()
			if result != test.out {
				t.Errorf("Ma did not return %v for %v", test.out, test.in)
			}
		})
	}
}

// TestMethodLa tests the La method.
func TestMethodLa(t *testing.T) {
	tests := []struct {
		name string
		in   Trit
		out  Trit
	}{
		{"La should return False for False", False, False},
		{"La should return False for Unknown", Unknown, False},
		{"La should return True for True", True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.in.La()
			if result != test.out {
				t.Errorf("La did not return %v for %v", test.out, test.in)
			}
		})
	}
}

// TestMethodIa tests the Ia method.
func TestMethodIa(t *testing.T) {
	tests := []struct {
		name string
		in   Trit
		out  Trit
	}{
		{"Ia should return False for False", False, False},
		{"Ia should return True for Unknown", Unknown, True},
		{"Ia should return False for True", True, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.in.Ia()
			if result != test.out {
				t.Errorf("Ia did not return %v for %v", test.out, test.in)
			}
		})
	}
}

// TestMethodAnd tests the And method.
func TestMethodAnd(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"And should return False for (False, False)",
			False, False, False,
		},
		{
			"And should return False for (False, Unknown)",
			False, Unknown, False,
		},
		{
			"And should return False for (False, True)",
			False, True, False,
		},
		{
			"And should return False for (Unknown, False)",
			Unknown, False, False,
		},
		{
			"And should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"And should return Unknown for (Unknown, True)",
			Unknown, True, Unknown,
		},
		{
			"And should return False for (True, False)",
			True, False, False,
		},
		{
			"And should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{
			"And should return True for (True, True)",
			True, True, True,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.And(test.b)
			if result != test.out {
				t.Errorf("And did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodOr tests the Or method.
func TestMethodOr(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Or should return False for (False, False)",
			False, False, False,
		},
		{
			"Or should return Unknown for (False, Unknown)",
			False, Unknown, Unknown,
		},
		{
			"Or should return True for (False, True)",
			False, True, True,
		},
		{
			"Or should return Unknown for (Unknown, False)",
			Unknown, False, Unknown,
		},
		{
			"Or should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"Or should return True for (Unknown, True)",
			Unknown, True, True,
		},
		{
			"Or should return True for (True, False)",
			True, False, True,
		},
		{
			"Or should return True for (True, Unknown)",
			True, Unknown, True,
		},
		{
			"Or should return True for (True, True)",
			True, True, True,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Or(test.b)
			if result != test.out {
				t.Errorf("Or did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodXor tests the Xor method.
func TestMethodXor(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Xor should return False for (False, False)",
			False, False, False,
		},
		{
			"Xor should return Unknown for (False, Unknown)",
			False, Unknown, Unknown,
		},
		{
			"Xor should return True for (False, True)",
			False, True, True,
		},
		{
			"Xor should return Unknown for (Unknown, False)",
			Unknown, False, Unknown,
		},
		{
			"Xor should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"Xor should return Unknown for (Unknown, True)",
			Unknown, True, Unknown,
		},
		{
			"Xor should return True for (True, False)",
			True, False, True,
		},
		{
			"Xor should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{
			"Xor should return False for (True, True)",
			True, True, False,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Xor(test.b)
			if result != test.out {
				t.Errorf("Xor did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodNand tests the Nand method.
func TestMethodNand(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Nand should return True for (False, False)",
			False, False, True,
		},
		{
			"Nand should return True for (False, Unknown)",
			False, Unknown, True,
		},
		{
			"Nand should return True for (False, True)",
			False, True, True,
		},
		{
			"Nand should return True for (Unknown, False)",
			Unknown, False, True,
		},
		{
			"Nand should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"Nand should return Unknown for (Unknown, True)",
			Unknown, True, Unknown,
		},
		{
			"Nand should return True for (True, False)",
			True, False, True,
		},
		{
			"Nand should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{
			"Nand should return False for (True, True)",
			True, True, False,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Nand(test.b)
			if result != test.out {
				t.Errorf("Nand did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodNor tests the Nor method.
func TestMethodNor(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Nor should return True for (False, False)",
			False, False, True,
		},
		{
			"Nor should return Unknown for (False, Unknown)",
			False, Unknown, Unknown,
		},
		{
			"Nor should return False for (False, True)",
			False, True, False,
		},
		{
			"Nor should return Unknown for (Unknown, False)",
			Unknown, False, Unknown,
		},
		{
			"Nor should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"Nor should return False for (Unknown, True)",
			Unknown, True, False,
		},
		{
			"Nor should return False for (True, False)",
			True, False, False,
		},
		{
			"Nor should return False for (True, Unknown)",
			True, Unknown, False,
		},
		{
			"Nor should return False for (True, True)",
			True, True, False,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Nor(test.b)
			if result != test.out {
				t.Errorf("Nor did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodNxor tests the Nxor method.
func TestMethodNxor(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Nxor should return True for (False, False)",
			False, False, True,
		},
		{
			"Nxor should return Unknown for (False, Unknown)",
			False, Unknown, Unknown,
		},
		{
			"Nxor should return False for (False, True)",
			False, True, False,
		},
		{
			"Nxor should return Unknown for (Unknown, False)",
			Unknown, False, Unknown,
		},
		{
			"Nxor should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"Nxor should return Unknown for (Unknown, True)",
			Unknown, True, Unknown,
		},
		{
			"Nxor should return False for (True, False)",
			True, False, False,
		},
		{
			"Nxor should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{
			"Nxor should return True for (True, True)",
			True, True, True,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Nxor(test.b)
			if result != test.out {
				t.Errorf("Nxor did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodMin tests the Min method.
func TestMethodMin(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Min should return False for (False, False)",
			False, False, False,
		},
		{
			"Min should return False for (False, Unknown)",
			False, Unknown, False,
		},
		{
			"Min should return False for (False, True)",
			False, True, False,
		},
		{
			"Min should return False for (Unknown, False)",
			Unknown, False, False,
		},
		{
			"Min should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"Min should return Unknown for (Unknown, True)",
			Unknown, True, Unknown,
		},
		{
			"Min should return False for (True, False)",
			True, False, False,
		},
		{
			"Min should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{
			"Min should return True for (True, True)",
			True, True, True,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Min(test.b)
			if result != test.out {
				t.Errorf("Min did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodMax tests the Max method.
func TestMethodMax(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Max should return False for (False, False)",
			False, False, False,
		},
		{
			"Max should return Unknown for (False, Unknown)",
			False, Unknown, Unknown,
		},
		{
			"Max should return True for (False, True)",
			False, True, True,
		},
		{
			"Max should return Unknown for (Unknown, False)",
			Unknown, False, Unknown,
		},
		{
			"Max should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"Max should return True for (Unknown, True)",
			Unknown, True, True,
		},
		{
			"Max should return True for (True, False)",
			True, False, True,
		},
		{
			"Max should return True for (True, Unknown)",
			True, Unknown, True,
		},
		{
			"Max should return True for (True, True)",
			True, True, True,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Max(test.b)
			if result != test.out {
				t.Errorf("Max did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodImp tests the Imp method.
func TestMethodImp(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Imp should return True for (False, False)",
			False, False, True,
		},
		{
			"Imp should return True for (False, Unknown)",
			False, Unknown, True,
		},
		{
			"Imp should return True for (False, True)",
			False, True, True,
		},
		{
			"Imp should return Unknown for (Unknown, False)",
			Unknown, False, Unknown,
		},
		{
			"Imp should return True for (Unknown, Unknown)",
			Unknown, Unknown, True,
		},
		{
			"Imp should return True for (Unknown, True)",
			Unknown, True, True,
		},
		{
			"Imp should return False for (True, False)",
			True, False, False,
		},
		{
			"Imp should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{
			"Imp should return True for (True, True)",
			True, True, True,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Imp(test.b)
			if result != test.out {
				t.Errorf("Imp did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodNimp tests the Nimp method.
func TestMethodNimp(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Nimp should return False for (False, False)",
			False, False, False,
		},
		{
			"Nimp should return False for (False, Unknown)",
			False, Unknown, False,
		},
		{
			"Nimp should return False for (False, True)",
			False, True, False,
		},
		{
			"Nimp should return Unknown for (Unknown, False)",
			Unknown, False, Unknown,
		},
		{
			"Nimp should return False for (Unknown, Unknown)",
			Unknown, Unknown, False,
		},
		{
			"Nimp should return False for (Unknown, True)",
			Unknown, True, False,
		},
		{
			"Nimp should return True for (True, False)",
			True, False, True,
		},
		{
			"Nimp should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{
			"Nimp should return False for (True, True)",
			True, True, False,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Nimp(test.b)
			if result != test.out {
				t.Errorf("Nimp did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodEq tests the Eq method.
func TestMethodEq(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Eq should return True for (False, False)",
			False, False, True,
		},
		{
			"Eq should return Unknown for (False, Unknown)",
			False, Unknown, Unknown,
		},
		{
			"Eq should return False for (False, True)",
			False, True, False,
		},
		{
			"Eq should return Unknown for (Unknown, False)",
			Unknown, False, Unknown,
		},
		{
			"Eq should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"Eq should return Unknown for (Unknown, True)",
			Unknown, True, Unknown,
		},
		{
			"Eq should return False for (True, False)",
			True, False, False,
		},
		{
			"Eq should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{
			"Eq should return True for (True, True)",
			True, True, True,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Eq(test.b)
			if result != test.out {
				t.Errorf("Eq did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMethodNeq tests the Neq method.
func TestMethodNeq(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{
			"Neq should return False for (False, False)",
			False, False, False,
		},
		{
			"Neq should return Unknown for (False, Unknown)",
			False, Unknown, Unknown,
		},
		{
			"Neq should return True for (False, True)",
			False, True, True,
		},
		{
			"Neq should return Unknown for (Unknown, False)",
			Unknown, False, Unknown,
		},
		{
			"Neq should return Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{
			"Neq should return Unknown for (Unknown, True)",
			Unknown, True, Unknown,
		},
		{
			"Neq should return True for (True, False)",
			True, False, True,
		},
		{
			"Neq should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{
			"Neq should return False for (True, True)",
			True, True, False,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Neq(test.b)
			if result != test.out {
				t.Errorf("Neq did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}
