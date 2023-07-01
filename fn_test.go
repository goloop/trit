package trit

import (
	"testing"
)

// TestParalellTasks tests the ParallelTasks function.
func TestParallelTasks(t *testing.T) {
	tests := []struct {
		name   string
		inputs []int
		want   int
	}{
		{
			name:   "No input values",
			inputs: []int{},
			want:   parallelTasks, // should return current value of pt
		},
		{
			name:   "Input values sum to less than 0",
			inputs: []int{-10, -5},
			want:   1, // should be set to 1
		},
		{
			name:   "Input values sum to 0",
			inputs: []int{-5, 5},
			want:   1, // should be set to 1
		},
		{
			name:   "Input values sum to more than maxParallelTasks",
			inputs: []int{10, maxParallelTasks + 1},
			want:   maxParallelTasks, // should be set to maxParallelTasks
		},
		{
			name:   "Input values sum to valid value",
			inputs: []int{3, 4},
			want:   7, // should be set to the sum of the inputs
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParallelTasks(test.inputs...)
			if got != test.want {
				t.Errorf("ParallelTasks() = %v, want %v", got, test.want)
			}
		})
	}
}

// TestIsFalse tests the IsFalse function.
func TestIsFalse(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  bool
	}{
		{"-1 should return true", -1, true},
		{"1 should return false", 1, false},
		{"0 should return false", 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsFalse(test.in)
			if result != test.out {
				t.Errorf("isFalse did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestIsUnknown tests the IsUnknown function.
func TestIsUnknown(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  bool
	}{
		{"-1 should return false", -1, false},
		{"1 should return false", 1, false},
		{"0 should return true", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsUnknown(test.in)
			if result != test.out {
				t.Errorf("IsUnknown did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestIsTrue tests the IsTrue function.
func TestIsTrue(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  bool
	}{
		{"-1 should return false", -1, false},
		{"1 should return true", 1, true},
		{"0 should return false", 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsTrue(test.in)
			if result != test.out {
				t.Errorf("isTrue did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestSet tests the Set function.
func TestSet(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  Trit
	}{
		{"-1 should return False", -1, False},
		{"1 should return True", 1, True},
		{"0 should return Unknown", 0, Unknown},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var trit Trit
			result := Set(&trit, test.in)
			if result != test.out {
				t.Errorf("Set did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestConvert tests the Convert function.
func TestConvert(t *testing.T) {
	tests := []struct {
		name string
		in   float64
		out  Trit
	}{
		{"-0.1 should return False", -0.1, False},
		{"7.7 should return True", 7.7, True},
		{"0.0 should return Unknown", 0.0, Unknown},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Convert(test.in)
			if result != test.out {
				t.Errorf("Any did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestAll tests the All function.
func TestAll(t *testing.T) {
	ParallelTasks(2)
	minLoadPeGoroutine = 20
	tests := []struct {
		name string
		in   []Trit
		out  Trit
	}{
		{"[1, 1, 1] should return True", []Trit{True, True, True}, True},
		{"[1, 0, 1] should return False", []Trit{True, False, True}, False},
		{"[0, 0, 0] should return False", []Trit{False, False, False}, False},
		{"[0, 0, 1] should return False", []Trit{False, False, True}, False},
		{"[0, 1, 0] should return False", []Trit{False, True, False}, False},
		{"[0, 1, 1] should return False", []Trit{False, True, True}, False},
		{"[1, 0, 0] should return False", []Trit{True, False, False}, False},
		{"[1, 1, 0] should return False", []Trit{True, True, False}, False},
		{
			name: "Just empty list",
			in:   []Trit{},
			out:  False,
		},
		{
			name: "A very large list for testing goroutines",
			in: []Trit{
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, Unknown, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
			},
			out: False,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := All(test.in...)
			if result != test.out {
				t.Errorf("All did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestAny tests the Any function.
func TestAny(t *testing.T) {
	ParallelTasks(2)
	minLoadPeGoroutine = 20
	tests := []struct {
		name string
		in   []Trit
		out  Trit
	}{
		{"[1, 1, 1] should return True", []Trit{True, True, True}, True},
		{"[1, 0, 1] should return True", []Trit{True, False, True}, True},
		{"[0, 0, 0] should return False", []Trit{False, False, False}, False},
		{"[0, 0, 1] should return True", []Trit{False, False, True}, True},
		{"[0, 1, 0] should return True", []Trit{False, True, False}, True},
		{"[0, 1, 1] should return True", []Trit{False, True, True}, True},
		{"[1, 0, 0] should return True", []Trit{True, False, False}, True},
		{"[1, 1, 0] should return True", []Trit{True, True, False}, True},
		{
			name: "Just empty list",
			in:   []Trit{},
			out:  False,
		},
		{
			name: "A very large list for testing goroutines",
			in: []Trit{
				False, False, False, False, False, False, False, False, False,
				False, False, False, False, False, False, False, False, False,
				False, False, False, False, False, False, False, False, False,
				False, False, False, False, False, False, False, False, False,
				False, False, False, False, False, False, False, False, False,
				False, False, False, False, False, False, False, False, False,
				False, False, False, False, False, False, False, False, False,
				False, False, False, True, False, False, False, False, False,
				False, False, False, False, False, False, False, False, False,
				False, False, False, False, False, False, False, False, False,
			},
			out: True,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Any(test.in...)
			if result != test.out {
				t.Errorf("Any did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestNone tests the None function.
func TestNone(t *testing.T) {
	tests := []struct {
		name string
		in   []Trit
		out  Trit
	}{
		{"[1, 1, 1] should return False", []Trit{True, True, True}, False},
		{"[1, 0, 1] should return False", []Trit{True, False, True}, False},
		{"[0, 0, 0] should return True", []Trit{False, False, False}, True},
		{"[0, 0, 1] should return False", []Trit{False, False, True}, False},
		{"[0, 1, 0] should return False", []Trit{False, True, False}, False},
		{"[0, 1, 1] should return False", []Trit{False, True, True}, False},
		{"[1, 0, 0] should return False", []Trit{True, False, False}, False},
		{"[1, 1, 0] should return False", []Trit{True, True, False}, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := None(test.in...)
			if result != test.out {
				t.Errorf("None did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestLogicToTrit tests the logicToTrit function.
func TestLogicToTrit(t *testing.T) {
	// Numbers.
	testsInt := []struct {
		name string
		in   int
		out  Trit
	}{
		{"1 should return True", 1, True},
		{"-1 should return False", -1, False},
		{"0 should return Unknown", 0, Unknown},
		{"-77 should return False", -77, False},
		{"1000000 should return True", 1000000, True},
	}

	for _, test := range testsInt {
		t.Run(test.name, func(t *testing.T) {
			result := logicToTrit(test.in)
			if result != test.out {
				t.Errorf("logicToTrit did not return %v for %v",
					test.out, test.in)
			}
		})
	}

	testsUint := []struct {
		name string
		in   uint
		out  Trit
	}{
		{"1 should return True", 1, True},
		{"0 should return Unknown", 0, Unknown},
		{"1000000 should return True", 1000000, True},
	}

	for _, test := range testsUint {
		t.Run(test.name, func(t *testing.T) {
			result := logicToTrit(test.in)
			if result != test.out {
				t.Errorf("logicToTrit did not return %v for %v",
					test.out, test.in)
			}
		})
	}

	testsFloat := []struct {
		name string
		in   float64
		out  Trit
	}{
		{"0.3 should return True", 0.3, True},
		{"-0.3 should return False", -0.3, False},
		{"0.0 should return Unknown", 0.0, Unknown},
		{"-77.5 should return False", -77.5, False},
		{"1000000.5 should return True", 1000000.5, True},
	}

	for _, test := range testsFloat {
		t.Run(test.name, func(t *testing.T) {
			result := logicToTrit(test.in)
			if result != test.out {
				t.Errorf("logicToTrit did not return %v for %v",
					test.out, test.in)
			}
		})
	}

	// Bool.
	testsBool := []struct {
		name string
		in   bool
		out  Trit
	}{
		{"trut should return True", true, True},
		{"false should return False", false, False},
	}

	for _, test := range testsBool {
		t.Run(test.name, func(t *testing.T) {
			result := logicToTrit(test.in)
			if result != test.out {
				t.Errorf("logicToTrit did not return %v for %v",
					test.out, test.in)
			}
		})
	}

	// Trit.
	testsTrit := []struct {
		name string
		in   Trit
		out  Trit
	}{
		{"True should return True", True, True},
		{"False should return False", False, False},
		{"Unknown should return False", Unknown, Unknown},
	}

	for _, test := range testsTrit {
		t.Run(test.name, func(t *testing.T) {
			result := logicToTrit(test.in)
			if result != test.out {
				t.Errorf("logicToTrit did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestDefault tests the Default method.
func TestDefault(t *testing.T) {
	t.Run("Default with bool value", func(t *testing.T) {
		t1 := Unknown
		Default(&t1, true)
		if t1 != True {
			t.Errorf("Default did not update Unknown to True")
		}

		t2 := Unknown
		Default(&t2, false)
		if t2 != False {
			t.Errorf("Default did not update Unknown to False")
		}
	})

	t.Run("Default with numeric value", func(t *testing.T) {
		t1 := Unknown
		Default(&t1, int32(1)) // for example int32
		if t1 != True {
			t.Errorf("Default did not update Unknown to True")
		}

		t2 := Unknown
		Default(&t2, int64(-1)) // for example int64
		if t2 != False {
			t.Errorf("Default did not update Unknown to False")
		}
	})

	t.Run("Default with Trit value", func(t *testing.T) {
		t1 := Unknown
		Default(&t1, True)
		if t1 != True {
			t.Errorf("Default did not update Unknown to True")
		}

		t2 := Unknown
		Default(&t2, False)
		if t2 != False {
			t.Errorf("Default did not update Unknown to False")
		}
	})

	t.Run("Should not update non-Unknown Trit", func(t *testing.T) {
		t1 := True
		Default(&t1, false)
		if t1 != True {
			t.Errorf("Default updated non-Unknown Trit")
		}
	})
}

// TestNot tests the Not function.
func TestNot(t *testing.T) {
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
			result := Not(test.in)
			if result != test.out {
				t.Errorf("Not did not return %v for %v", test.out, test.in)
			}
		})
	}
}

// TestMa tests the Ma function.
func TestMa(t *testing.T) {
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
			result := Ma(test.in)
			if result != test.out {
				t.Errorf("Ma did not return %v for %v", test.out, test.in)
			}
		})
	}
}

// TestLa tests the La function.
func TestLa(t *testing.T) {
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
			result := La(test.in)
			if result != test.out {
				t.Errorf("La did not return %v for %v", test.out, test.in)
			}
		})
	}
}

// TestIa tests the Ia function.
func TestIa(t *testing.T) {
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
			result := Ia(test.in)
			if result != test.out {
				t.Errorf("Ia did not return %v for %v", test.out, test.in)
			}
		})
	}
}

// TestAnd tests the And function.
func TestAnd(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Should be False for (False, False)", False, False, False},
		{"Should be False for (False, Unknown)", False, Unknown, False},
		{"Should be False for (False, True)", False, True, False},
		{"Should be False for (Unknown, False)", Unknown, False, False},
		{
			"Should be Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{"Should be Unknown for (Unknown, True)", Unknown, True, Unknown},
		{"Should be False for (True, False)", True, False, False},
		{"Should be Unknown for (True, Unknown)", True, Unknown, Unknown},
		{"Should be True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := And(test.a, test.b)
			if result != test.out {
				t.Errorf("And did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestOr tests the Or function.
func TestOr(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Should be False for (False, False)", False, False, False},
		{"Should be Unknown for (False, Unknown)", False, Unknown, Unknown},
		{"Should be True for (False, True)", False, True, True},
		{"Should be Unknown for (Unknown, False)", Unknown, False, Unknown},
		{
			"Should be Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{"Should be True for (Unknown, True)", Unknown, True, True},
		{"Should be True for (True, False)", True, False, True},
		{"Should be True for (True, Unknown)", True, Unknown, True},
		{"Should be True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Or(test.a, test.b)
			if result != test.out {
				t.Errorf("Or did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestXor tests the Xor function.
func TestXor(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Should be False for (False, False)", False, False, False},
		{"Should be Unknown for (False, Unknown)", False, Unknown, Unknown},
		{"Should be True for (False, True)", False, True, True},
		{"Should be Unknown for (Unknown, False)", Unknown, False, Unknown},
		{
			"Should be Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{"Should be Unknown for (Unknown, True)", Unknown, True, Unknown},
		{"Should be True for (True, False)", True, False, True},
		{"Should be Unknown for (True, Unknown)", True, Unknown, Unknown},
		{"Should be False for (True, True)", True, True, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Xor(test.a, test.b)
			if result != test.out {
				t.Errorf("Xor did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestNand tests the Nand function.
func TestNand(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Should be True for (False, False)", False, False, True},
		{"Should be True for (False, Unknown)", False, Unknown, True},
		{"Should be True for (False, True)", False, True, True},
		{"Should be True for (Unknown, False)", Unknown, False, True},
		{
			"Should be Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{"Should be Unknown for (Unknown, True)", Unknown, True, Unknown},
		{"Should be True for (True, False)", True, False, True},
		{"Should be Unknown for (True, Unknown)", True, Unknown, Unknown},
		{"Should be False for (True, True)", True, True, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Nand(test.a, test.b)
			if result != test.out {
				t.Errorf("Nand did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestNor tests the Nor function.
func TestNor(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Should be True for (False, False)", False, False, True},
		{"Should be Unknown for (False, Unknown)", False, Unknown, Unknown},
		{"Should be False for (False, True)", False, True, False},
		{"Should be Unknown for (Unknown, False)", Unknown, False, Unknown},
		{
			"Should be Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{"Should be False for (Unknown, True)", Unknown, True, False},
		{"Should be False for (True, False)", True, False, False},
		{"Should be False for (True, Unknown)", True, Unknown, False},
		{"Should be False for (True, True)", True, True, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Nor(test.a, test.b)
			if result != test.out {
				t.Errorf("Nor did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestNxor tests the Nxor function.
func TestNxor(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Should be True for (False, False)", False, False, True},
		{"Should be Unknown for (False, Unknown)", False, Unknown, Unknown},
		{"Should be False for (False, True)", False, True, False},
		{"Should be Unknown for (Unknown, False)", Unknown, False, Unknown},
		{
			"Should be Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{"Should be Unknown for (Unknown, True)", Unknown, True, Unknown},
		{"Should be False for (True, False)", True, False, False},
		{"Should be Unknown for (True, Unknown)", True, Unknown, Unknown},
		{"Should be True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Nxor(test.a, test.b)
			if result != test.out {
				t.Errorf("Nxor did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMin tests the Min function.
func TestMin(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Should be False for (False, False)", False, False, False},
		{"Should be False for (False, Unknown)", False, Unknown, False},
		{"Should be False for (False, True)", False, True, False},
		{"Should be False for (Unknown, False)", Unknown, False, False},
		{
			"Should be Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{"Should be Unknown for (Unknown, True)", Unknown, True, Unknown},
		{"Should be False for (True, False)", True, False, False},
		{"Should be Unknown for (True, Unknown)", True, Unknown, Unknown},
		{"Should be True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Min(test.a, test.b)
			if result != test.out {
				t.Errorf("Min did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestMax tests the Max function.
func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Should be False for (False, False)", False, False, False},
		{"Should be Unknown for (False, Unknown)", False, Unknown, Unknown},
		{"Should be True for (False, True)", False, True, True},
		{"Should be Unknown for (Unknown, False)", Unknown, False, Unknown},
		{
			"Should be Unknown for (Unknown, Unknown)",
			Unknown, Unknown, Unknown,
		},
		{"Should be True for (Unknown, True)", Unknown, True, True},
		{"Should be True for (True, False)", True, False, True},
		{"Should be True for (True, Unknown)", True, Unknown, True},
		{"Should be True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Max(test.a, test.b)
			if result != test.out {
				t.Errorf("Max did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestImp tests the Imp function.
func TestImp(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Should be True for (False, False)", False, False, True},
		{"Should be True for (False, Unknown)", False, Unknown, True},
		{"Should be True for (False, True)", False, True, True},
		{"Should be Unknown for (Unknown, False)", Unknown, False, Unknown},
		{"Should be True for (Unknown, Unknown)", Unknown, Unknown, True},
		{"Should be True for (Unknown, True)", Unknown, True, True},
		{"Should be False for (True, False)", True, False, False},
		{"Should be Unknown for (True, Unknown)", True, Unknown, Unknown},
		{"Should be True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Imp(test.a, test.b)
			if result != test.out {
				t.Errorf("Imp did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestNimp tests the Nimp function.
func TestNimp(t *testing.T) {
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
			result := Nimp(test.a, test.b)
			if result != test.out {
				t.Errorf("Nimp did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestEq tests the Eq function.
func TestEq(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Eq should return True for (False, False)", False, False, True},
		{
			"Eq should return Unknown for (False, Unknown)",
			False, Unknown, Unknown,
		},
		{"Eq should return False for (False, True)", False, True, False},
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
		{"Eq should return False for (True, False)", True, False, False},
		{
			"Eq should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{"Eq should return True for (True, True)", True, True, True},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Eq(test.a, test.b)
			if result != test.out {
				t.Errorf("Eq did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestNeq tests the Neq function.
func TestNeq(t *testing.T) {
	tests := []struct {
		name string
		a    Trit
		b    Trit
		out  Trit
	}{
		{"Neq should return False for (False, False)", False, False, False},
		{
			"Neq should return Unknown for (False, Unknown)",
			False, Unknown, Unknown,
		},
		{"Neq should return True for (False, True)", False, True, True},
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
		{"Neq should return True for (True, False)", True, False, True},
		{
			"Neq should return Unknown for (True, Unknown)",
			True, Unknown, Unknown,
		},
		{"Neq should return False for (True, True)", True, True, False},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Neq(test.a, test.b)
			if result != test.out {
				t.Errorf("Neq did not return %v for (%v, %v)",
					test.out, test.a, test.b)
			}
		})
	}
}

// TestKnown tests the Known function.
func TestKnown(t *testing.T) {
	ParallelTasks(2)
	minLoadPeGoroutine = 20
	tests := []struct {
		name string
		in   []Trit
		out  Trit
	}{
		{
			name: "Known should return True for (True, True, True)",
			in:   []Trit{True, True, True},
			out:  True,
		},
		{
			name: "Known should return True for (True, True, False)",
			in:   []Trit{True, True, False},
			out:  True,
		},
		{
			name: "Known should return True for (False, False, False)",
			in:   []Trit{False, False, False},
			out:  True,
		},
		{
			name: "Known should return False for (False, Unknown, True)",
			in:   []Trit{False, Unknown, True},
			out:  False,
		},
		{
			name: "Known should return False for (Unknown, Unknown)",
			in:   []Trit{Unknown, Unknown},
			out:  False,
		},
		{
			name: "Just empty list",
			in:   []Trit{},
			out:  False,
		},
		{
			name: "A very large list for testing goroutines",
			in: []Trit{
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, False, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, False, True, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, Unknown, True, True, True, True, True, True,
				True, True, True, True, True, True, True, True, True, True,
				True, True, True, True, True, True, False, False, False,
			},
			out: False,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Known(test.in...)
			if result != test.out {
				t.Errorf("Known did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestIsConfidence tests the IsConfidence function.
func TestIsConfidence(t *testing.T) {
	tests := []struct {
		name string
		in   []Trit
		out  bool
	}{
		{
			name: "IsConfidence should return true for (True, True, True)",
			in:   []Trit{True, True, True},
			out:  true,
		},
		{
			name: "IsConfidence should return true for (True, True, False)",
			in:   []Trit{True, True, False},
			out:  true,
		},
		{
			name: "IsConfidence should return true for (False, False, False)",
			in:   []Trit{False, False, False},
			out:  true,
		},
		{
			name: "IsConfidence should return false (False, Unknown, True)",
			in:   []Trit{False, Unknown, True},
			out:  false,
		},
		{
			name: "IsConfidence should return false for (Unknown, Unknown)",
			in:   []Trit{Unknown, Unknown},
			out:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsConfidence(test.in...)
			if result != test.out {
				t.Errorf("IsConfidence did not return %v for %v",
					test.out, test.in)
			}
		})
	}
}

// TestRandom tests the Random function.
func TestRandom(t *testing.T) {
	// No arguments.
	result := Random()
	if !(result == False || result == True || result == Unknown) {
		t.Errorf("Something went wrong with Random()")
	}

	// Unknow - 0%
	for i := 0; i < 100; i++ {
		result = Random(0)
		if result == Unknown {
			t.Errorf("Expected a random value as True or False, got %v",
				result)
		}
	}

	// Two arguments.
	for i := 0; i < 100; i++ {
		result = Random(90, 5, 5)
		if result != Unknown {
			t.Errorf("Expected a random value as Unknown, got %v", result)
		}
	}
}
