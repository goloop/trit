package trit

import (
	"context"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"time"
)

var (
	// The parallelTasks the number of parallel tasks.
	parallelTasks = 1

	// The maxParallelTasks is the maximum number of parallel tasks.
	maxParallelTasks = runtime.NumCPU() * 3

	// The minLoadPerGoroutine is the minimum slice size for processing
	// in an individual goroutine. Essentially, it delineates the threshold
	// at which it becomes worthwhile to divide the slice processing amongst
	// multiple goroutines. If each goroutine isn't handling a sufficiently
	// large subslice, the overhead of goroutine creation and management
	// may outweigh the benefits of concurrent processing. This variable
	// specifies the minimum number of iterations per goroutine to ensure
	// an efficient division of labor.
	minLoadPeGoroutine = 1024
)

// Logicable is a special data type from which to determine the state of Trit
// in the context of three-valued logic.
type Logicable interface {
	bool | Trit |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Tritter is a special data type that implements the Trit interface.
type Tritter interface {
	IsTrue() bool
	IsFalse() bool
	IsUnknown() bool
}

// The logicFoundValue is a helper struct that holds a boolean value
// and a Mutex to protect it from concurrent access.
//
// They are used in the In function to detect the desired result
// in a separate goroutine.
type logicFoundValue struct {
	m     sync.Mutex
	value Trit
}

// SetValue sets a new value for the Found. It locks the Mutex before
// changing the value and unlocks it after the change is complete.
func (f *logicFoundValue) SetValue(value Trit) {
	f.m.Lock()
	defer f.m.Unlock()
	f.value = value
}

// GetValue retrieves the current value of the Found. It locks the Mutex
// before reading the value and unlocks it after the read is complete.
func (f *logicFoundValue) GetValue() Trit {
	f.m.Lock()
	defer f.m.Unlock()
	return f.value
}

// The init initializes the randomGenerator variable.
func init() {
	parallelTasks = runtime.NumCPU() * 2
}

// ParallelTasks returns the number of parallel tasks.
//
// If the function is called without parameters, it returns the
// current value of parallelTasks.
//
// A function can receive one or more values for parallelTasks,
// these values are added together to form the final result for
// parallelTasks. If the new value for parallelTasks is less than
// or equal to zero - it will be set to 1, if it is greater than
// maxParallelTasks - it will be set to maxParallelTasks.
func ParallelTasks(v ...int) int {
	if len(v) > 0 {
		n := 0
		for _, value := range v {
			n += value
		}

		if n <= 0 {
			parallelTasks = 1
		} else if n > maxParallelTasks {
			parallelTasks = maxParallelTasks
		} else {
			parallelTasks = n
		}
	}

	return parallelTasks
}

// The logicToTrit function converts any logic type to Trit.
func logicToTrit[T Logicable](v T) Trit {
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
			value := reflect.ValueOf(v).Int()
			if value > 0 {
				return True
			} else if value < 0 {
				return False
			}

			return Unknown
		}
	case uint, uint8, uint16, uint32, uint64:
		switch reflect.TypeOf(v).Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64:
			value := reflect.ValueOf(v).Uint()
			if value > 0 {
				return True
			}

			// Can't be less than 0
			return Unknown
		}
	case float32, float64:
		switch reflect.TypeOf(v).Kind() {
		case reflect.Float32, reflect.Float64:
			value := reflect.ValueOf(v).Float()
			if value > 0 {
				return True
			} else if value < 0 {
				return False
			}

			return Unknown
		}
	}

	return any(v).(Trit)
}

// Default sets the default value for the trit-object
// if this one has a Unknown state.
//
// Example usage:
//
//	t := trit.Unknown
//	trit.Default(&t, trit.True)
//	fmt.Println(t.String()) // Output: True
func Default[T Logicable](t *Trit, v T) Trit {
	// If the trit is not Unknown, return the trit.
	if t.Val() != Unknown {
		return *t
	}

	trit := logicToTrit(v)
	*t = trit
	return *t
}

// IsFalse checks if the trit-object is False.
//
// See Trit.IsFalse() for more information.
func IsFalse[T Logicable](t T) bool {
	trit := logicToTrit(t)
	return trit.IsFalse()
}

// IsUnknown checks if the trit-object is Unknown.
//
// See Trit.IsUnknown() for more information.
func IsUnknown[T Logicable](t T) bool {
	trit := logicToTrit(t)
	return trit.IsUnknown()
}

// IsTrue checks if the trit-object is True.
//
// See Trit.IsTrue() for more information.
func IsTrue[T Logicable](t T) bool {
	trit := logicToTrit(t)
	return trit.IsTrue()
}

// Set sets the value of the trit-object.
//
// See Trit.Set() for more information.
func Set[T Logicable](t *Trit, v T) Trit {
	trit := logicToTrit(v)
	*t = trit
	return *t
}

// Convert converts the any Logicable type to Trit.
//
// Example usage:
//
//	t := trit.Convert(true)
//	fmt.Println(t.String()) // Output: True
func Convert[T Logicable](v T) Trit {
	trit := logicToTrit(v)
	return trit
}

// All returns True if all the trit-objects are True.
//
// Example usage:
//
//	t := trit.All(trit.True, trit.True, trit.True)
//	fmt.Println(t.String()) // Output: True
func All[T Logicable](t ...T) Trit {
	var wg sync.WaitGroup

	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := parallelTasks
	found := &logicFoundValue{value: True}

	// If the length of the slice is less than or equal to
	// the minLoadPeGoroutine, then we do not need
	// to use goroutines.
	if l := len(t); l == 0 {
		return False
	} else if l/p < minLoadPeGoroutine {
		for _, v := range t {
			trit := logicToTrit(v)
			if trit.IsFalse() || trit.IsUnknown() {
				return False
			}
		}

		return True
	}

	chunkSize := len(t) / p
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(t)
		}

		go func(start, end int) {
			defer wg.Done()

			for _, b := range t[start:end] {
				trit := logicToTrit(b)
				// Check if the context has been cancelled.
				select {
				case <-ctx.Done():
					return
				default:
				}

				if trit.IsFalse() || trit.IsUnknown() {
					found.SetValue(False)
					cancel() // stop all other goroutines
					return
				}
			}
		}(start, end)
	}

	wg.Wait()
	return found.GetValue()
}

// Any returns True if any of the trit-objects are True.
//
// Example usage:
//
//	t := trit.Any(trit.True, trit.False, trit.False)
//	fmt.Println(t.String()) // Output: True
func Any[T Logicable](t ...T) Trit {
	var wg sync.WaitGroup

	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := parallelTasks
	found := &logicFoundValue{value: False}

	// If the length of the slice is less than or equal to
	// the minLoadPeGoroutine, then we do not need
	// to use goroutines.
	if l := len(t); l == 0 {
		return False
	} else if l/p < minLoadPeGoroutine {
		for _, v := range t {
			trit := logicToTrit(v)
			if trit.IsTrue() {
				return True
			}
		}

		return False
	}

	chunkSize := len(t) / p
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(t)
		}

		go func(start, end int) {
			defer wg.Done()

			for _, b := range t[start:end] {
				trit := logicToTrit(b)

				// Check if the context has been cancelled.
				select {
				case <-ctx.Done():
					return
				default:
				}

				if trit.IsTrue() {
					found.SetValue(True)
					cancel() // stop all other goroutines
					return
				}
			}
		}(start, end)
	}

	wg.Wait()
	return found.GetValue()
}

// None returns True if none of the trit-objects are True.
//
// Example usage:
//
//	t := trit.None(trit.False, trit.False, trit.False)
//	fmt.Println(t.String()) // Output: True
func None[T Logicable](t ...T) Trit {
	for _, v := range t {
		trit := logicToTrit(v)
		if trit.IsTrue() {
			return False
		}
	}

	return True
}

// Not performs a logical NOT operation on a Trit-Like value and
// returns the result as Trit.
//
// See Trit.Not() for more information.
func Not[T Logicable](t T) Trit {
	trit := logicToTrit(t)
	return trit.Not()
}

// Ma performs a logical MA (Modus Ponens Absorption) operation
// on a Trit-Like values and returns the result as Trit.
//
// See Trit.Ma() for more information.
func Ma[T Logicable](t T) Trit {
	trit := logicToTrit(t)
	return trit.Ma()
}

// La performs a logical LA (Law of Absorption) operation on a Trit-Like
// value and returns the result as Trit.
//
// See Trit.La() for more information.
func La[T Logicable](t T) Trit {
	trit := logicToTrit(t)
	return trit.La()
}

// Ia performs a logical IA (Idempotent Absorption) operation on a Trit-Like
// value and returns the result as Trit.
//
// See Trit.Ia() for more information.
func Ia[T Logicable](t T) Trit {
	trit := logicToTrit(t)
	return trit.Ia()
}

// And performs a logical AND operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.And() for more information.
func And[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.And(tb)
}

// Or performs a logical OR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Or() for more information.
func Or[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Or(tb)
}

// Xor performs a logical XOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Xor() for more information.
func Xor[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Xor(tb)
}

// Nand performs a logical NAND operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nand() for more information.
func Nand[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nand(tb)
}

// Nor performs a logical NOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nor() for more information.
func Nor[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nor(tb)
}

// Nxor performs a logical NXOR operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nxor() for more information.
func Nxor[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nxor(tb)
}

// Min performs a logical MIN operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Min() for more information.
func Min[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Min(tb)
}

// Max performs a logical MAX operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Max() for more information.
func Max[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Max(tb)
}

// Imp performs a logical IMP operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Imp() for more information.
func Imp[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Imp(tb)
}

// Nimp performs a logical NIMP operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Nimp() for more information.
func Nimp[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Nimp(tb)
}

// Eq performs a logical EQ operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Eq() for more information.
func Eq[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Eq(tb)
}

// Neq performs a logical NEQ operation between two Trit-Like values
// and returns the result as Trit.
//
// See Trit.Neq() for more information.
func Neq[T, U Logicable](a T, b U) Trit {
	ta := logicToTrit(a)
	tb := logicToTrit(b)
	return ta.Neq(tb)
}

// Known returns boolean true if all Trit-Like values has definiteness,
// i.e. is either True or False.
//
// Example usage:
//
//	a := trit.Known(trit.True, trit.False, trit.Unknown)
//	fmt.Println(a.String()) // Output: False
//
//	b := trit.Known(trit.True, trit.True, trit.False)
//	fmt.Println(b.String()) // Output: True
func Known[T Logicable](ts ...T) Trit {
	var wg sync.WaitGroup

	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := parallelTasks
	found := &logicFoundValue{value: True}

	// If the length of the slice is less than or equal to
	// the minLoadPeGoroutine, then we do not need
	// to use goroutines.
	if l := len(ts); l == 0 {
		return False
	} else if l/p < minLoadPeGoroutine {
		for _, t := range ts {
			trit := logicToTrit(t)
			if trit == Unknown {
				return False
			}
		}

		return True
	}

	chunkSize := len(ts) / p
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(ts)
		}

		go func(start, end int) {
			defer wg.Done()

			for _, b := range ts[start:end] {
				trit := logicToTrit(b)

				// Check if the context has been cancelled.
				select {
				case <-ctx.Done():
					return
				default:
				}

				if trit.IsUnknown() {
					found.SetValue(False)
					cancel() // stop all other goroutines
					return
				}
			}
		}(start, end)
	}

	wg.Wait()
	return found.GetValue()
}

// IsConfidence returns boolean true if all Trit-Like values has definiteness,
// i.e. is either True or False.
//
// Example usage:
//
//	a := trit.IsConfidence(trit.True, trit.False, trit.Unknown)
//	fmt.Println(a.String()) // Output: False
//
//	b := trit.IsConfidence(trit.True, trit.True, trit.False)
//	fmt.Println(b.String()) // Output: True
func IsConfidence[T Logicable](ts ...T) bool {
	for _, t := range ts {
		trit := logicToTrit(t)
		if trit == Unknown {
			return false
		}
	}

	return true
}

// Random returns a random Trit value.
// The function can accept an optional argument that indicates the
// percentage probability of the occurrence of the Unknown event.
//
// Example usage:
//
//	 a := trit.Random()
//	 fmt.Println(a.String()) // Output: True, False or Unknown
//
//	b := trit.Random(0)
//	fmt.Println(b.String()) // Output: True or False
//
//	c := trit.Random(50)
//	fmt.Println(c.String()) // Output: With a probability of 50% is Unknown
func Random(up ...uint8) Trit {
	// Determination of the probability of occurrence of the event Unknown.
	var p int

	if len(up) == 0 {
		p = 33
	} else {
		for _, v := range up {
			p += int(v)
		}
	}

	if p > 100 {
		p = 100
	}

	// Generate random value.
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(100)

	if value < p {
		return Unknown
	}

	if value < (100-p)/2 {
		return True
	}

	return False
}
