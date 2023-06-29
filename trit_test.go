package trit

import "testing"

// TestMethodDefault tests the Default method.
func TestMethodDefault(t *testing.T) {
	t.Run("Def should update Nil to True", func(t *testing.T) {
		t1 := Nil
		result := t1.Default(True)
		if result != True {
			t.Errorf("Def did not update Nil to True")
		}
	})

	t.Run("Def should update Nil to False", func(t *testing.T) {
		t1 := Nil
		result := t1.Default(False)
		if result != False {
			t.Errorf("Def did not update Nil to False")
		}
	})

	t.Run("Def should not update non-Nil Trit", func(t *testing.T) {
		t1 := True
		result := t1.Default(False)
		if result != True {
			t.Errorf("Def updated non-Nil Trit")
		}
	})
}

// TestTrueIfNil tests the TrueIfNil method.
func TestTrueIfNil(t *testing.T) {
	t.Run("Should set Nil to True", func(t *testing.T) {
		tr := Nil
		tr.TrueIfNil()
		if tr != True {
			t.Errorf("TrueIfNil did not set Nil to True")
		}
	})

	t.Run("Should not change True to False", func(t *testing.T) {
		tr := True
		tr.TrueIfNil()
		if tr != True {
			t.Errorf("TrueIfNil changed True to False")
		}
	})

	t.Run("Should not change False to True", func(t *testing.T) {
		tr := False
		tr.TrueIfNil()
		if tr != False {
			t.Errorf("TrueIfNil changed False to True")
		}
	})
}

// TestFalseIfNil tests the FalseIfNil method.
func TestFalseIfNil(t *testing.T) {
	t.Run("Should set Nil to False", func(t *testing.T) {
		tr := Nil
		tr.FalseIfNil()
		if tr != False {
			t.Errorf("FalseIfNil did not set Nil to False")
		}
	})

	t.Run("Should not change False to True", func(t *testing.T) {
		tr := False
		tr.FalseIfNil()
		if tr != False {
			t.Errorf("FalseIfNil changed False to True")
		}
	})

	t.Run("Should not change True to False", func(t *testing.T) {
		tr := True
		tr.FalseIfNil()
		if tr != True {
			t.Errorf("FalseIfNil changed True to False")
		}
	})
}

// TestClean tests the Clean method.
func TestClean(t *testing.T) {
	t.Run("Clean should set Nil to Nil", func(t *testing.T) {
		tr := Nil
		tr.Clean()
		if tr != Nil {
			t.Errorf("Clean did not set Nil to Nil")
		}
	})

	t.Run("Clean should not change False to Nil", func(t *testing.T) {
		tr := False
		tr.Clean()
		if tr != False {
			t.Errorf("Clean changed False to Nil")
		}
	})

	t.Run("Clean should not change True to Nil", func(t *testing.T) {
		tr := True
		tr.Clean()
		if tr != True {
			t.Errorf("Clean changed True to Nil")
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

	t.Run("IsFalse should return false for Nil", func(t *testing.T) {
		tr := Nil
		if tr.IsFalse() {
			t.Errorf("IsFalse returned true for Nil")
		}
	})

	t.Run("IsFalse should return false for True", func(t *testing.T) {
		tr := True
		if tr.IsFalse() {
			t.Errorf("IsFalse returned true for True")
		}
	})
}

// TestMethodIsNil tests the IsNil method.
func TestMethodIsNil(t *testing.T) {
	t.Run("IsNil should return false for False", func(t *testing.T) {
		tr := False
		if tr.IsNil() {
			t.Errorf("IsNil returned true for False")
		}
	})

	t.Run("IsNil should return true for Nil", func(t *testing.T) {
		tr := Nil
		if !tr.IsNil() {
			t.Errorf("IsNil returned false for Nil")
		}
	})

	t.Run("IsNil should return false for True", func(t *testing.T) {
		tr := True
		if tr.IsNil() {
			t.Errorf("IsNil returned true for True")
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

	t.Run("IsTrue should return false for Nil", func(t *testing.T) {
		tr := Nil
		if tr.IsTrue() {
			t.Errorf("IsTrue returned true for Nil")
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

	t.Run("Set should set value to Nil for zero", func(t *testing.T) {
		tr := Trit(1)
		tr.Set(0)
		if tr != Nil {
			t.Errorf("Set did not set value to Nil for zero")
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

	t.Run("Val should return Nil for zero Trit", func(t *testing.T) {
		tr := Trit(0)
		if tr.Val() != Nil {
			t.Errorf("Val did not return Nil for zero Trit")
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

	t.Run("Norm should normalize Nil to Nil", func(t *testing.T) {
		tr := Nil
		tr.Norm()
		if tr != Nil {
			t.Errorf("Norm did not normalize Nil to Nil")
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

	t.Run("Int should return 0 for Nil", func(t *testing.T) {
		tr := Nil
		if tr.Int() != 0 {
			t.Errorf("Int did not return 0 for Nil")
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

	t.Run("String should return 'Nil' for Nil", func(t *testing.T) {
		tr := Nil
		if tr.String() != "Nil" {
			t.Errorf("String did not return 'Nil' for Nil")
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
		{"Not should return Nil for Nil", Nil, Nil},
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
		{"Ma should return True for Nil", Nil, True},
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
		{"La should return False for Nil", Nil, False},
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
		{"Ia should return True for Nil", Nil, True},
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
		{"And should return False for (False, False)", False, False, False},
		{"And should return False for (False, Nil)", False, Nil, False},
		{"And should return False for (False, True)", False, True, False},
		{"And should return False for (Nil, False)", Nil, False, False},
		{"And should return Nil for (Nil, Nil)", Nil, Nil, Nil},
		{"And should return Nil for (Nil, True)", Nil, True, Nil},
		{"And should return False for (True, False)", True, False, False},
		{"And should return Nil for (True, Nil)", True, Nil, Nil},
		{"And should return True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.And(test.b)
			if result != test.out {
				t.Errorf("And did not return %v for (%v, %v)", test.out, test.a, test.b)
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
		{"Or should return False for (False, False)", False, False, False},
		{"Or should return Nil for (False, Nil)", False, Nil, Nil},
		{"Or should return True for (False, True)", False, True, True},
		{"Or should return Nil for (Nil, False)", Nil, False, Nil},
		{"Or should return Nil for (Nil, Nil)", Nil, Nil, Nil},
		{"Or should return True for (Nil, True)", Nil, True, True},
		{"Or should return True for (True, False)", True, False, True},
		{"Or should return True for (True, Nil)", True, Nil, True},
		{"Or should return True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Or(test.b)
			if result != test.out {
				t.Errorf("Or did not return %v for (%v, %v)", test.out, test.a, test.b)
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
		{"Xor should return False for (False, False)", False, False, False},
		{"Xor should return Nil for (False, Nil)", False, Nil, Nil},
		{"Xor should return True for (False, True)", False, True, True},
		{"Xor should return Nil for (Nil, False)", Nil, False, Nil},
		{"Xor should return Nil for (Nil, Nil)", Nil, Nil, Nil},
		{"Xor should return Nil for (Nil, True)", Nil, True, Nil},
		{"Xor should return True for (True, False)", True, False, True},
		{"Xor should return Nil for (True, Nil)", True, Nil, Nil},
		{"Xor should return False for (True, True)", True, True, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Xor(test.b)
			if result != test.out {
				t.Errorf("Xor did not return %v for (%v, %v)", test.out, test.a, test.b)
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
		{"Nand should return True for (False, False)", False, False, True},
		{"Nand should return True for (False, Nil)", False, Nil, True},
		{"Nand should return True for (False, True)", False, True, True},
		{"Nand should return True for (Nil, False)", Nil, False, True},
		{"Nand should return Nil for (Nil, Nil)", Nil, Nil, Nil},
		{"Nand should return Nil for (Nil, True)", Nil, True, Nil},
		{"Nand should return True for (True, False)", True, False, True},
		{"Nand should return Nil for (True, Nil)", True, Nil, Nil},
		{"Nand should return False for (True, True)", True, True, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Nand(test.b)
			if result != test.out {
				t.Errorf("Nand did not return %v for (%v, %v)", test.out, test.a, test.b)
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
		{"Nor should return True for (False, False)", False, False, True},
		{"Nor should return Nil for (False, Nil)", False, Nil, Nil},
		{"Nor should return False for (False, True)", False, True, False},
		{"Nor should return Nil for (Nil, False)", Nil, False, Nil},
		{"Nor should return Nil for (Nil, Nil)", Nil, Nil, Nil},
		{"Nor should return False for (Nil, True)", Nil, True, False},
		{"Nor should return False for (True, False)", True, False, False},
		{"Nor should return False for (True, Nil)", True, Nil, False},
		{"Nor should return False for (True, True)", True, True, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Nor(test.b)
			if result != test.out {
				t.Errorf("Nor did not return %v for (%v, %v)", test.out, test.a, test.b)
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
		{"Nxor should return True for (False, False)", False, False, True},
		{"Nxor should return Nil for (False, Nil)", False, Nil, Nil},
		{"Nxor should return False for (False, True)", False, True, False},
		{"Nxor should return Nil for (Nil, False)", Nil, False, Nil},
		{"Nxor should return Nil for (Nil, Nil)", Nil, Nil, Nil},
		{"Nxor should return Nil for (Nil, True)", Nil, True, Nil},
		{"Nxor should return False for (True, False)", True, False, False},
		{"Nxor should return Nil for (True, Nil)", True, Nil, Nil},
		{"Nxor should return True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Nxor(test.b)
			if result != test.out {
				t.Errorf("Nxor did not return %v for (%v, %v)", test.out, test.a, test.b)
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
		{"Min should return False for (False, False)", False, False, False},
		{"Min should return False for (False, Nil)", False, Nil, False},
		{"Min should return False for (False, True)", False, True, False},
		{"Min should return False for (Nil, False)", Nil, False, False},
		{"Min should return Nil for (Nil, Nil)", Nil, Nil, Nil},
		{"Min should return Nil for (Nil, True)", Nil, True, Nil},
		{"Min should return False for (True, False)", True, False, False},
		{"Min should return Nil for (True, Nil)", True, Nil, Nil},
		{"Min should return True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Min(test.b)
			if result != test.out {
				t.Errorf("Min did not return %v for (%v, %v)", test.out, test.a, test.b)
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
		{"Max should return False for (False, False)", False, False, False},
		{"Max should return Nil for (False, Nil)", False, Nil, Nil},
		{"Max should return True for (False, True)", False, True, True},
		{"Max should return Nil for (Nil, False)", Nil, False, Nil},
		{"Max should return Nil for (Nil, Nil)", Nil, Nil, Nil},
		{"Max should return True for (Nil, True)", Nil, True, True},
		{"Max should return True for (True, False)", True, False, True},
		{"Max should return True for (True, Nil)", True, Nil, True},
		{"Max should return True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Max(test.b)
			if result != test.out {
				t.Errorf("Max did not return %v for (%v, %v)", test.out, test.a, test.b)
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
		{"Imp should return True for (False, False)", False, False, True},
		{"Imp should return True for (False, Nil)", False, Nil, True},
		{"Imp should return True for (False, True)", False, True, True},
		{"Imp should return Nil for (Nil, False)", Nil, False, Nil},
		{"Imp should return True for (Nil, Nil)", Nil, Nil, True},
		{"Imp should return True for (Nil, True)", Nil, True, True},
		{"Imp should return False for (True, False)", True, False, False},
		{"Imp should return Nil for (True, Nil)", True, Nil, Nil},
		{"Imp should return True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.a.Imp(test.b)
			if result != test.out {
				t.Errorf("Imp did not return %v for (%v, %v)", test.out, test.a, test.b)
			}
		})
	}
}
