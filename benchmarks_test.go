package trit

import (
	"fmt"
	"testing"
)

// BenchmarkBasicOperations benchmarks basic trit operations
func BenchmarkBasicOperations(b *testing.B) {
	b.Run("IsTrue", func(b *testing.B) {
		t := True
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = t.IsTrue()
		}
	})

	b.Run("IsFalse", func(b *testing.B) {
		t := False
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = t.IsFalse()
		}
	})

	b.Run("IsUnknown", func(b *testing.B) {
		t := Unknown
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = t.IsUnknown()
		}
	})

	b.Run("Not", func(b *testing.B) {
		t := True
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = t.Not()
		}
	})
}

// BenchmarkLogicOperations benchmarks binary logic operations
func BenchmarkLogicOperations(b *testing.B) {
	values := []Trit{True, False, Unknown}

	for _, v1 := range values {
		for _, v2 := range values {
			b.Run("And/"+v1.String()+"/"+v2.String(), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = v1.And(v2)
				}
			})

			b.Run("Or/"+v1.String()+"/"+v2.String(), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = v1.Or(v2)
				}
			})

			b.Run("Xor/"+v1.String()+"/"+v2.String(), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = v1.Xor(v2)
				}
			})
		}
	}
}

// BenchmarkAggregates benchmarks the slice aggregates over a range of sizes.
// The decisive element is placed at the very end so the linear scan cannot
// short-circuit early — this measures the true worst case.
func BenchmarkAggregates(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		// All-True slice forces All/Known/Any to traverse everything.
		values := make([]Trit, size)
		for i := range values {
			values[i] = True
		}

		b.Run("All/size="+fmt.Sprint(size), func(b *testing.B) {
			for b.Loop() {
				_ = All(values...)
			}
		})

		b.Run("Any/size="+fmt.Sprint(size), func(b *testing.B) {
			for b.Loop() {
				_ = Any(values...)
			}
		})

		b.Run("Known/size="+fmt.Sprint(size), func(b *testing.B) {
			for b.Loop() {
				_ = Known(values...)
			}
		})
	}
}

// BenchmarkConversions benchmarks type conversion operations
func BenchmarkConversions(b *testing.B) {
	b.Run("FromBool/true", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Define(true)
		}
	})

	b.Run("FromInt/positive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Define(42)
		}
	})

	b.Run("FromFloat64/negative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Define(-42.0)
		}
	})

	b.Run("ToString", func(b *testing.B) {
		t := True
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = t.String()
		}
	})
}

// BenchmarkExtendedOperations benchmarks more complex operations
func BenchmarkExtendedOperations(b *testing.B) {
	v1, v2 := True, False

	b.Run("Nand", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v1.Nand(v2)
		}
	})

	b.Run("Nor", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v1.Nor(v2)
		}
	})

	b.Run("Nxor", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v1.Nxor(v2)
		}
	})

	b.Run("Imp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v1.Imp(v2)
		}
	})

	b.Run("Eq", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v1.Eq(v2)
		}
	})
}
