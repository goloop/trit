[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/trit/v2)](https://goreportcard.com/report/github.com/goloop/trit/v2) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/trit/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://pkg.go.dev/github.com/goloop/trit/v2) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# trit

`trit` implements three-valued (ternary) logic in Go. A `Trit` is an `int8`
with three states — `False` (any negative number), `Unknown` (`0`) and `True`
(any positive number) — and the full set of ternary logic operators (`Not`,
`And`, `Or`, `Xor`, `Nand`, `Nor`, `Nxor`, implication, equivalence and more).

The key property is that **the zero value is `Unknown`**, so an uninitialized
`Trit` is already meaningful. That is what makes it useful wherever a "maybe" or
"not set" state matters — config merging, partial updates, nullable database
columns, logic circuits — where a plain `bool` cannot tell "false" apart from
"unset".

## Features

- Three-valued logic operations (True, False, Unknown); zero value is Unknown.
- Safe bool conversions with explicit `Unknown` handling.
- Serialization: JSON (Unknown → `null`), text, and `database/sql`
  (Unknown → `NULL`).
- `ParseTrit`, `Compare` (ordering False < Unknown < True), and `Default`.
- Slice aggregates: `All`, `Any`, `None`, `Known`, `Consensus`, `Majority`
  (plus `iter.Seq` forms).
- Zero allocations for basic operations; property and fuzz tested.

## Installation

```bash
go get github.com/goloop/trit/v2
```

```go
import "github.com/goloop/trit/v2"
```

Requires Go 1.24 or newer.

## Quick start

```go
package main

import (
    "fmt"

    "github.com/goloop/trit/v2"
)

func main() {
    t1, t2 := trit.True, trit.False

    fmt.Println(t1.IsTrue())  // true
    fmt.Println(t1.And(t2))   // False
    fmt.Println(t1.Or(t2))    // True
    fmt.Println(t1.Xor(t2))   // True

    // Resolve the Unknown state.
    t3 := trit.Unknown
    t3.Default(trit.True)     // set only if currently Unknown
    fmt.Println(t3)           // True
}
```

Why three-valued logic? A `bool` config field cannot distinguish "explicitly
false" from "not provided". A `trit.Trit` field can — an unset field stays
`Unknown`, so you can apply a default without clobbering an explicit choice:

```go
type Config struct {
    Enabled trit.Trit // Unknown until the caller sets it
}

func (s *Service) Configure(c Config) {
    if !c.Enabled.IsUnknown() {
        s.enabled = c.Enabled
    }
    trit.Default(&s.enabled, true) // default only if still unset
}
```

## Documentation

- Full reference, truth tables and recipes: [DOC.md](DOC.md) · [DOC.UK.md](DOC.UK.md)
- Package API: [pkg.go.dev/github.com/goloop/trit/v2](https://pkg.go.dev/github.com/goloop/trit/v2)
- Changes between versions: [CHANGELOG.md](CHANGELOG.md)

## Contributing

Contributions are welcome. Please run `go test ./...`, `go vet ./...` and
`gofmt -l .` before submitting a pull request.

## License

`trit` is released under the MIT License. See [LICENSE](LICENSE).
