package trit_test

import (
	"testing"

	"github.com/goloop/trit"
)

// TestTruthTables tests of truth tables.
func TestTruthTables(t *testing.T) {
	values := []trit.Trit{trit.False, trit.Unknown, trit.True}
	names := map[trit.Trit]string{
		trit.False:   "F",
		trit.Unknown: "U",
		trit.True:    "T",
	}

	// Test unary operations
	t.Run("Unary Operations", func(t *testing.T) {
		type testCase struct {
			op   string
			fn   func(trit.Trit) trit.Trit
			want map[trit.Trit]trit.Trit
		}

		tests := []testCase{
			{
				op: "NA",
				fn: func(a trit.Trit) trit.Trit { return a.Not() },
				want: map[trit.Trit]trit.Trit{
					trit.False:   trit.True,
					trit.Unknown: trit.Unknown,
					trit.True:    trit.False,
				},
			},
			{
				op: "MA",
				fn: func(a trit.Trit) trit.Trit { return a.Ma() },
				want: map[trit.Trit]trit.Trit{
					trit.False:   trit.False,
					trit.Unknown: trit.True,
					trit.True:    trit.True,
				},
			},
			{
				op: "LA",
				fn: func(a trit.Trit) trit.Trit { return a.La() },
				want: map[trit.Trit]trit.Trit{
					trit.False:   trit.False,
					trit.Unknown: trit.False,
					trit.True:    trit.True,
				},
			},
			{
				op: "IA",
				fn: func(a trit.Trit) trit.Trit { return a.Ia() },
				want: map[trit.Trit]trit.Trit{
					trit.False:   trit.False,
					trit.Unknown: trit.True,
					trit.True:    trit.False,
				},
			},
		}

		for _, test := range tests {
			t.Run(test.op, func(t *testing.T) {
				for _, v := range values {
					got := test.fn(v)
					want := test.want[v]
					if got != want {
						t.Errorf("%s(%s) = %s, want %s",
							test.op, names[v], names[got], names[want])
					}
				}
			})
		}
	})

	// Test binary operations
	t.Run("Binary Operations", func(t *testing.T) {
		type testCase struct {
			op   string
			fn   func(trit.Trit, trit.Trit) trit.Trit
			want map[string]trit.Trit
		}

		tests := []testCase{
			{
				op: "AND",
				fn: func(a, b trit.Trit) trit.Trit { return a.And(b) },
				want: map[string]trit.Trit{
					"F,F": trit.False, "F,U": trit.False, "F,T": trit.False,
					"U,F": trit.False, "U,U": trit.Unknown, "U,T": trit.Unknown,
					"T,F": trit.False, "T,U": trit.Unknown, "T,T": trit.True,
				},
			},
			{
				op: "OR",
				fn: func(a, b trit.Trit) trit.Trit { return a.Or(b) },
				want: map[string]trit.Trit{
					"F,F": trit.False, "F,U": trit.Unknown, "F,T": trit.True,
					"U,F": trit.Unknown, "U,U": trit.Unknown, "U,T": trit.True,
					"T,F": trit.True, "T,U": trit.True, "T,T": trit.True,
				},
			},
			{
				op: "XOR",
				fn: func(a, b trit.Trit) trit.Trit { return a.Xor(b) },
				want: map[string]trit.Trit{
					"F,F": trit.False, "F,U": trit.Unknown, "F,T": trit.True,
					"U,F": trit.Unknown, "U,U": trit.Unknown, "U,T": trit.Unknown,
					"T,F": trit.True, "T,U": trit.Unknown, "T,T": trit.False,
				},
			},
			{
				op: "NAND",
				fn: func(a, b trit.Trit) trit.Trit { return a.Nand(b) },
				want: map[string]trit.Trit{
					"F,F": trit.True, "F,U": trit.True, "F,T": trit.True,
					"U,F": trit.True, "U,U": trit.Unknown, "U,T": trit.Unknown,
					"T,F": trit.True, "T,U": trit.Unknown, "T,T": trit.False,
				},
			},
			{
				op: "NOR",
				fn: func(a, b trit.Trit) trit.Trit { return a.Nor(b) },
				want: map[string]trit.Trit{
					"F,F": trit.True, "F,U": trit.Unknown, "F,T": trit.False,
					"U,F": trit.Unknown, "U,U": trit.Unknown, "U,T": trit.False,
					"T,F": trit.False, "T,U": trit.False, "T,T": trit.False,
				},
			},
			{
				op: "NXOR",
				fn: func(a, b trit.Trit) trit.Trit { return a.Nxor(b) },
				want: map[string]trit.Trit{
					"F,F": trit.True, "F,U": trit.Unknown, "F,T": trit.False,
					"U,F": trit.Unknown, "U,U": trit.Unknown, "U,T": trit.Unknown,
					"T,F": trit.False, "T,U": trit.Unknown, "T,T": trit.True,
				},
			},
			{
				op: "IMP",
				fn: func(a, b trit.Trit) trit.Trit { return a.Imp(b) },
				want: map[string]trit.Trit{
					"F,F": trit.True, "F,U": trit.True, "F,T": trit.True,
					"U,F": trit.Unknown, "U,U": trit.True, "U,T": trit.True,
					"T,F": trit.False, "T,U": trit.Unknown, "T,T": trit.True,
				},
			},
			{
				op: "EQ",
				fn: func(a, b trit.Trit) trit.Trit { return a.Eq(b) },
				want: map[string]trit.Trit{
					"F,F": trit.True, "F,U": trit.Unknown, "F,T": trit.False,
					"U,F": trit.Unknown, "U,U": trit.Unknown, "U,T": trit.Unknown,
					"T,F": trit.False, "T,U": trit.Unknown, "T,T": trit.True,
				},
			},
			{
				op: "MIN",
				fn: func(a, b trit.Trit) trit.Trit { return a.Min(b) },
				want: map[string]trit.Trit{
					"F,F": trit.False, "F,U": trit.False, "F,T": trit.False,
					"U,F": trit.False, "U,U": trit.Unknown, "U,T": trit.Unknown,
					"T,F": trit.False, "T,U": trit.Unknown, "T,T": trit.True,
				},
			},
			{
				op: "NIMP",
				fn: func(a, b trit.Trit) trit.Trit { return a.Nimp(b) },
				want: map[string]trit.Trit{
					"F,F": trit.False, "F,U": trit.False, "F,T": trit.False,
					"U,F": trit.Unknown, "U,U": trit.False, "U,T": trit.False,
					"T,F": trit.True, "T,U": trit.Unknown, "T,T": trit.False,
				},
			},
			{
				op: "NEQ",
				fn: func(a, b trit.Trit) trit.Trit { return a.Neq(b) },
				want: map[string]trit.Trit{
					"F,F": trit.False, "F,U": trit.Unknown, "F,T": trit.True,
					"U,F": trit.Unknown, "U,U": trit.Unknown, "U,T": trit.Unknown,
					"T,F": trit.True, "T,U": trit.Unknown, "T,T": trit.False,
				},
			},
			{
				op: "MAX",
				fn: func(a, b trit.Trit) trit.Trit { return a.Max(b) },
				want: map[string]trit.Trit{
					"F,F": trit.False, "F,U": trit.Unknown, "F,T": trit.True,
					"U,F": trit.Unknown, "U,U": trit.Unknown, "U,T": trit.True,
					"T,F": trit.True, "T,U": trit.True, "T,T": trit.True,
				},
			},
		}

		for _, test := range tests {
			t.Run(test.op, func(t *testing.T) {
				for _, a := range values {
					for _, b := range values {
						key := names[a] + "," + names[b]
						got := test.fn(a, b)
						want := test.want[key]
						if got != want {
							t.Errorf("%s(%s,%s) = %s, want %s",
								test.op, names[a], names[b], names[got], names[want])
						}
					}
				}
			})
		}
	})
}
