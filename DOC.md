# trit — reference

The full reference for the `trit` package: the mental model, states and
conversions, the logic operators, the default-state helpers, serialization,
slice aggregates, the truth tables and practical recipes.

Ukrainian version: **[DOC.UK.md](DOC.UK.md)**.

## Contents

- [Mental model](#mental-model)
- [States and the Logicable constraint](#states-and-the-logicable-constraint)
- [Inspecting and setting a Trit](#inspecting-and-setting-a-trit)
- [Default-state helpers](#default-state-helpers)
- [Conversions](#conversions)
- [Logic operators](#logic-operators)
- [Comparison and ordering](#comparison-and-ordering)
- [Parsing](#parsing)
- [Serialization](#serialization)
- [Slice aggregates](#slice-aggregates)
- [Errors](#errors)
- [Truth tables](#truth-tables)
- [Recipes and tips](#recipes-and-tips)

## Mental model

`trit` implements three-valued (ternary) logic. A `Trit` is an `int8` with three
meaningful states:

| State | Value |
|-------|-------|
| `False`   | any negative number (canonically `-1`) |
| `Unknown` | `0` |
| `True`    | any positive number (canonically `1`) |

The key property: **the zero value is `Unknown`**, so an uninitialized `Trit` is
already meaningful. That is what makes `trit` useful wherever a "maybe" or "not
set" state matters — config merging, partial updates, nullable database columns,
logic circuits — where a plain `bool` cannot tell "false" apart from "unset".

Because any negative/positive integer counts as False/True, normalize to the
canonical `-1/0/1` with `Norm`/`Val` when you need the exact value.

```go
import "github.com/goloop/trit/v2"
```

## States and the Logicable constraint

Many functions are generic over `Logicable` — the set of types that can stand in
for a ternary value: `bool`, the integer kinds, and `Trit` itself. This lets you
pass a `bool`, an `int` or a `Trit` interchangeably:

```go
trit.Define(true)  // True
trit.Define(0)     // Unknown
trit.Define(-5)    // False
```

`Define(v)` converts any `Logicable` to the corresponding `Trit`.

## Inspecting and setting a Trit

```go
func (t Trit) IsTrue() bool
func (t Trit) IsFalse() bool
func (t Trit) IsUnknown() bool
func (t *Trit) Set(v int) Trit
func (t Trit) Val() Trit    // normalized copy (-1/0/1)
func (t *Trit) Norm() Trit  // normalize in place
func (t Trit) Int() int
func (t Trit) String() string
```

There are also generic package-level predicates that accept any `Logicable`:

```go
func IsTrue[T Logicable](t T) bool
func IsFalse[T Logicable](t T) bool
func IsUnknown[T Logicable](t T) bool
```

```go
t := trit.True
t.IsTrue()   // true
t.Int()      // 1
t.String()   // "True"
```

## Default-state helpers

These resolve the `Unknown` state; they are the reason to reach for `trit`:

```go
func (t *Trit) Default(trit Trit) Trit
func (t *Trit) TrueIfUnknown() Trit
func (t *Trit) FalseIfUnknown() Trit
func (t *Trit) Clean() Trit

func Default[T Logicable](t *Trit, v T) Trit
```

`Default` sets the value only if it is currently `Unknown`, otherwise leaves it
untouched — perfect for "apply a default without overwriting an explicit
choice". `TrueIfUnknown`/`FalseIfUnknown` collapse `Unknown` to a definite state;
`Clean` resets to `Unknown`.

```go
t := trit.Unknown
t.Default(trit.True) // now True

trit.Default(&cfg.Enabled, true) // set a default only if unset
```

## Conversions

```go
func (t Trit) CanBeBool() bool
func (t Trit) ToBool() (bool, error)
```

`ToBool` returns the boolean value, or [`ErrUnknownValue`](#errors) for
`Unknown` — so converting an unresolved state to `bool` is an explicit,
checkable operation rather than a silent guess. `CanBeBool` reports whether the
conversion would succeed.

```go
b, err := trit.True.ToBool() // true, nil
_, err = trit.Unknown.ToBool() // errors.Is(err, trit.ErrUnknownValue)
```

## Logic operators

Each operator is a method taking a `Trit` and returning a `Trit`, following the
ternary [truth tables](#truth-tables):

| Method | Operation |
|--------|-----------|
| `Not()`         | logical NOT |
| `And`, `Or`, `Xor` | AND / OR / exclusive OR |
| `Nand`, `Nor`, `Nxor` | negated AND / OR / XOR |
| `Imp`, `Nimp`   | implication (Łukasiewicz) and its negation |
| `Eq`, `Neq`     | equivalence (iff) and its negation |
| `Min`, `Max`    | minimum / maximum (False < Unknown < True) |
| `Ma`, `La`, `Ia` | Modus-Ponens / Law-of / Implication absorption |

```go
t1, t2 := trit.True, trit.False
t1.And(t2) // False
t1.Or(t2)  // True
t1.Xor(t2) // True
t1.Not()   // False
```

## Comparison and ordering

```go
func (t Trit) Compare(o Trit) int
```

`Compare` orders `False < Unknown < True` and returns `-1`, `0` or `1` — the
same contract as `cmp.Compare`, so it plugs directly into `slices.SortFunc`.

## Parsing

```go
func ParseTrit(s string) (Trit, error)
```

Parses a textual representation into a `Trit`, returning
[`ErrInvalidTrit`](#errors) for an unrecognised string. It accepts the common
spellings of the three states (e.g. `"true"`, `"false"`, `"unknown"`, and their
short/numeric forms).

## Serialization

`Trit` implements the standard serialization interfaces, mapping `Unknown` to
the "null" of each format:

| Interface | Unknown becomes |
|-----------|-----------------|
| `json.Marshaler`/`Unmarshaler` (`MarshalJSON`/`UnmarshalJSON`) | JSON `null` |
| `encoding.TextMarshaler`/`TextUnmarshaler` (`MarshalText`/`UnmarshalText`) | text form |
| `database/sql` (`Value`/`Scan`) | SQL `NULL` |

```go
data, _ := json.Marshal(trit.Unknown) // null
var t trit.Trit
_ = json.Unmarshal([]byte("true"), &t) // True
```

This makes `Trit` a natural fit for a nullable boolean column: `NULL` ⇄
`Unknown`, and the two definite states round-trip.

## Slice aggregates

Reduce many `Logicable` values (or an `iter.Seq`) to a single `Trit`:

```go
func All[T Logicable](t ...T) Trit        func AllSeq[T Logicable](seq iter.Seq[T]) Trit
func Any[T Logicable](t ...T) Trit        func AnySeq[T Logicable](seq iter.Seq[T]) Trit
func None[T Logicable](t ...T) Trit       func NoneSeq[T Logicable](seq iter.Seq[T]) Trit
func Known[T Logicable](ts ...T) Trit     func KnownSeq[T Logicable](seq iter.Seq[T]) Trit
func Consensus[T Logicable](trits ...T) Trit
func Majority[T Logicable](trits ...T) Trit
```

- `All` / `Any` / `None` — the ternary quantifiers.
- `Known` — whether every value is definite (not `Unknown`).
- `Consensus` — the shared value when all agree, else `Unknown`.
- `Majority` — the value held by more than half.

```go
trit.All(true, true, trit.Unknown) // Unknown
trit.Any(false, trit.True)         // True
trit.Majority(true, true, false)   // True
```

## Errors

```go
var ErrInvalidTrit  = errors.New("invalid trit value")
var ErrUnknownValue = errors.New("cannot convert Unknown to bool")
```

`ParseTrit` returns `ErrInvalidTrit`; `ToBool` returns `ErrUnknownValue` for
`Unknown`. Match either with `errors.Is`.

## Truth tables

```
 Truth Tables of Three-valued logic
 (T=True, N=Unknown, F=False)

  A  | NA      A  | MA      A  | LA      A  | IA
 ----+----    ----+----    ----+----    ----+----
  F  |  T      F  |  F      F  |  F      F  |  F
  U  |  U      U  |  T      U  |  F      U  |  T
  T  |  F      T  |  T      T  |  T      T  |  F

  A | B | AND       A | B |  OR       A | B | XOR
 ---+---+------    ---+---+------    ---+---+------
  F | F |  F        F | F |  F        F | F |  F
  F | U |  F        F | U |  U        F | U |  U
  F | T |  F        F | T |  T        F | T |  T
  U | F |  F        U | F |  U        U | F |  U
  U | U |  U        U | U |  U        U | U |  U
  U | T |  U        U | T |  T        U | T |  U
  T | F |  F        T | F |  T        T | F |  T
  T | U |  U        T | U |  T        T | U |  U
  T | T |  T        T | T |  T        T | T |  F

  A | B | NAND      A | B | NOR       A | B | NXOR
 ---+---+------    ---+---+------    ---+---+------
  F | F |  T        F | F |  T        F | F |  T
  F | U |  T        F | U |  U        F | U |  U
  F | T |  T        F | T |  F        F | T |  F
  U | F |  T        U | F |  U        U | F |  U
  U | U |  U        U | U |  U        U | U |  U
  U | T |  U        U | T |  F        U | T |  U
  T | F |  T        T | F |  F        T | F |  F
  T | U |  U        T | U |  F        T | U |  U
  T | T |  F        T | T |  F        T | T |  T

  A | B | IMP       A | B |  EQ       A | B | MIN
 ---+---+------    ---+---+------    ---+---+------
  F | F |  T        F | F |  T        F | F |  F
  F | U |  T        F | U |  U        F | U |  F
  F | T |  T        F | T |  F        F | T |  F
  U | F |  U        U | F |  U        U | F |  F
  U | U |  T        U | U |  U        U | U |  U
  U | T |  T        U | T |  U        U | T |  U
  T | F |  F        T | F |  F        T | F |  F
  T | U |  U        T | U |  U        T | U |  U
  T | T |  T        T | T |  T        T | T |  T

  A | B | NIMP      A | B | NEQ       A | B | MAX
 ---+---+------    ---+---+------    ---+---+------
  F | F |  F        F | F |  F        F | F |  F
  F | U |  F        F | U |  U        F | U |  U
  F | T |  F        F | T |  T        F | T |  T
  U | F |  U        U | F |  U        U | F |  U
  U | U |  F        U | U |  U        U | U |  U
  U | T |  F        U | T |  U        U | T |  T
  T | F |  T        T | F |  T        T | F |  T
  T | U |  U        T | U |  U        T | U |  T
  T | T |  F        T | T |  F        T | T |  T
```

## Recipes and tips

**Partial config updates.** Give optional flags the type `trit.Trit`. An unset
field stays `Unknown`, so you can apply `trit.Default(&field, fallback)` without
clobbering an explicit `False` — the problem a plain `bool` cannot solve.

**Nullable booleans.** Store a `Trit` in a nullable SQL column: `NULL` ⇄
`Unknown` via `Value`/`Scan`, and JSON `null` ⇄ `Unknown` via the JSON methods.

**Aggregate votes.** Use `Consensus` when every input must agree, `Majority` for
a simple vote, and `Known` to check whether all inputs are decided before acting.

**Normalize before comparing raw values.** Any negative/positive integer is
False/True; call `Norm`/`Val` to get the canonical `-1/0/1` when the exact
`int8` matters.

**Convert deliberately.** Use `ToBool` (and check the error) rather than assuming
`Unknown` is false — the explicit `ErrUnknownValue` keeps the ambiguity visible.
