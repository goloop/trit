package trit

import "testing"

// TestDef tests the Def method.
func TestDef(t *testing.T) {
	t.Run("Def should update Nil to True", func(t *testing.T) {
		t1 := Nil
		result := t1.Def(True)
		if result != True {
			t.Errorf("Def did not update Nil to True")
		}
	})

	t.Run("Def should update Nil to False", func(t *testing.T) {
		t1 := Nil
		result := t1.Def(False)
		if result != False {
			t.Errorf("Def did not update Nil to False")
		}
	})

	t.Run("Def should not update non-Nil Trit", func(t *testing.T) {
		t1 := True
		result := t1.Def(False)
		if result != True {
			t.Errorf("Def updated non-Nil Trit")
		}
	})
}

// TestDefTrue tests the DefTrue method.
func TestDefTrue(t *testing.T) {
	t.Run("DefTrue should set Nil to True", func(t *testing.T) {
		tr := Nil
		tr.DefTrue()
		if tr != True {
			t.Errorf("DefTrue did not set Nil to True")
		}
	})

	t.Run("DefTrue should not change True to False", func(t *testing.T) {
		tr := True
		tr.DefTrue()
		if tr != True {
			t.Errorf("DefTrue changed True to False")
		}
	})

	t.Run("DefTrue should not change False to True", func(t *testing.T) {
		tr := False
		tr.DefTrue()
		if tr != False {
			t.Errorf("DefTrue changed False to True")
		}
	})
}

// TestDefFalse tests the DefFalse method.
func TestDefFalse(t *testing.T) {
	t.Run("DefFalse should set Nil to False", func(t *testing.T) {
		tr := Nil
		tr.DefFalse()
		if tr != False {
			t.Errorf("DefFalse did not set Nil to False")
		}
	})

	t.Run("DefFalse should not change False to True", func(t *testing.T) {
		tr := False
		tr.DefFalse()
		if tr != False {
			t.Errorf("DefFalse changed False to True")
		}
	})

	t.Run("DefFalse should not change True to False", func(t *testing.T) {
		tr := True
		tr.DefFalse()
		if tr != True {
			t.Errorf("DefFalse changed True to False")
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

// TestIsFalse tests the IsFalse method.
func TestIsFalse(t *testing.T) {
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

// TestIsNil tests the IsNil method.
func TestIsNil(t *testing.T) {
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

// TestIsTrue tests the IsTrue method.
func TestIsTrue(t *testing.T) {
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

// TestSet tests the Set method.
func TestSet(t *testing.T) {
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

// TestNot tests the Not method.
func TestNot(t *testing.T) {
	t.Run("Not should return False for True", func(t *testing.T) {
		tr := True
		result := tr.Not()
		if result != False {
			t.Errorf("Not did not return False for True")
		}
	})

	t.Run("Not should return Nil for Nil", func(t *testing.T) {
		tr := Nil
		result := tr.Not()
		if result != Nil {
			t.Errorf("Not did not return Nil for Nil")
		}
	})

	t.Run("Not should return True for False", func(t *testing.T) {
		tr := False
		result := tr.Not()
		if result != True {
			t.Errorf("Not did not return True for False")
		}
	})
}

// TestAnd tests the And method.
func TestAnd(t *testing.T) {
	t.Run("And should return True for (True, True)", func(t *testing.T) {
		t1 := True
		t2 := True
		result := t1.And(t2)
		if result != True {
			t.Errorf("And did not return True for (True, True)")
		}
	})

	t.Run("And should return Nil for (True, Nil)", func(t *testing.T) {
		t1 := True
		t2 := Nil
		result := t1.And(t2)
		if result != Nil {
			t.Errorf("And did not return Nil for (True, Nil)")
		}
	})

	t.Run("And should return False for (True, False)", func(t *testing.T) {
		t1 := True
		t2 := False
		result := t1.And(t2)
		if result != False {
			t.Errorf("And did not return False for (True, False)")
		}
	})
}

// TestOr tests the Or method.
func TestOr(t *testing.T) {
	t.Run("Or should return True for (True, True)", func(t *testing.T) {
		t1 := True
		t2 := True
		result := t1.Or(t2)
		if result != True {
			t.Errorf("Or did not return True for (True, True)")
		}
	})

	t.Run("Or should return True for (True, Nil)", func(t *testing.T) {
		t1 := True
		t2 := Nil
		result := t1.Or(t2)
		if result != True {
			t.Errorf("Or did not return True for (True, Nil)")
		}
	})

	t.Run("Or should return True for (True, False)", func(t *testing.T) {
		t1 := True
		t2 := False
		result := t1.Or(t2)
		if result != True {
			t.Errorf("Or did not return True for (True, False)")
		}
	})
}

// TestXor tests the Xor method.
func TestXor(t *testing.T) {
	t.Run("Xor should return False for (True, True)", func(t *testing.T) {
		t1 := True
		t2 := True
		result := t1.Xor(t2)
		if result != False {
			t.Errorf("Xor did not return False for (True, True)")
		}
	})

	t.Run("Xor should return Nil for (True, Nil)", func(t *testing.T) {
		t1 := True
		t2 := Nil
		result := t1.Xor(t2)
		if result != Nil {
			t.Errorf("Xor did not return Nil for (True, Nil)")
		}
	})

	t.Run("Xor should return True for (True, False)", func(t *testing.T) {
		t1 := True
		t2 := False
		result := t1.Xor(t2)
		if result != True {
			t.Errorf("Xor did not return True for (True, False)")
		}
	})

	// Add more test cases to cover all possible scenarios
}

// TestNand tests the Nand method.
func TestNand(t *testing.T) {
	t.Run("Nand should return False for (True, True)", func(t *testing.T) {
		t1 := True
		t2 := True
		result := t1.Nand(t2)
		if result != False {
			t.Errorf("Nand did not return False for (True, True)")
		}
	})

	t.Run("Nand should return True for (True, Nil)", func(t *testing.T) {
		t1 := True
		t2 := Nil
		result := t1.Nand(t2)
		if result != True {
			t.Errorf("Nand did not return True for (True, Nil)")
		}
	})

	t.Run("Nand should return True for (True, False)", func(t *testing.T) {
		t1 := True
		t2 := False
		result := t1.Nand(t2)
		if result != True {
			t.Errorf("Nand did not return True for (True, False)")
		}
	})
}

// TestNor tests the Nor method.
func TestNor(t *testing.T) {
	t.Run("Nor should return False for (True, True)", func(t *testing.T) {
		t1 := True
		t2 := True
		result := t1.Nor(t2)
		if result != False {
			t.Errorf("Nor did not return False for (True, True)")
		}
	})

	t.Run("Nor should return False for (True, Nil)", func(t *testing.T) {
		t1 := True
		t2 := Nil
		result := t1.Nor(t2)
		if result != False {
			t.Errorf("Nor did not return False for (True, Nil)")
		}
	})

	t.Run("Nor should return False for (True, False)", func(t *testing.T) {
		t1 := True
		t2 := False
		result := t1.Nor(t2)
		if result != False {
			t.Errorf("Nor did not return False for (True, False)")
		}
	})
}

// TestXnor tests the Xnor method.
func TestXnor(t *testing.T) {
	t.Run("Xnor should return True for (True, True)", func(t *testing.T) {
		t1 := True
		t2 := True
		result := t1.Xnor(t2)
		if result != True {
			t.Errorf("Xnor did not return True for (True, True)")
		}
	})

	t.Run("Xnor should return Nil for (True, Nil)", func(t *testing.T) {
		t1 := True
		t2 := Nil
		result := t1.Xnor(t2)
		if result != Nil {
			t.Errorf("Xnor did not return Nil for (True, Nil)")
		}
	})

	t.Run("Xnor should return False for (True, False)", func(t *testing.T) {
		t1 := True
		t2 := False
		result := t1.Xnor(t2)
		if result != False {
			t.Errorf("Xnor did not return False for (True, False)")
		}
	})
}
