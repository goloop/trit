[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/trit?v=2)](https://goreportcard.com/report/github.com/goloop/trit) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/trit/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/trit) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)


# trit

Package for working with three-valued (trinary) logic in Go.

Package defines the data type and basic operations of a ternary logic system, often referred to as trinary or ternary logic. It supports three states, namely False, Unknown, and True, with False represented by  negative number (-1), Unknown represented by zero (0), and True represented by positive number (1).

The logic operations NOT, AND, OR, XOR, NAND, NOR, and XNOR are implemented as methods of the Trit type, with each method applying the logic operation on the Trit receiver and a Trit argument to produce a Trit result according to the trinary truth table.

There are also some useful methods provided for checking the state of a Trit value (IsFalse, IsUnknown, IsTrue), setting a Trit value from an integer (Set), returning the underlying int8 value of a Trit (Int), and returning a string representation of a Trit value (String).

Overall, this package can be beneficial in scenarios where a "maybe" or "unknown" state is needed, such as in database systems and logic circuits, and in the construction of trinary computers and balanced ternary systems.

## Installation
```bash
go get github.com/goloop/trit
```

## Features
- Three-valued logic operations (True, False, Unknown)
- Safe bool conversions with Unknown state handling
- JSON marshaling (Unknown → null)
- Thread-safe parallel operations for slices
- Zero allocations for basic operations
- Extensive test coverage
- Comprehensive truth tables documentation

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/goloop/trit"
)

func main() {
	t1 := trit.True
	t2 := trit.False

	// IsTrue
	fmt.Println(t1.IsTrue()) // Output: True

	// IsFalse
	fmt.Println(t2.IsFalse()) // Output: True

	// Default
	t3 := trit.Unknown
	t3.Default(trit.True)
	fmt.Println(t3) // Output: True

	// TrueIfUnknown
	t4 := trit.Unknown
	t4.TrueIfUnknown()
	fmt.Println(t4) // Output: True

	// FalseIfUnknown
	t4 = trit.Unknown
	t4.FalseIfUnknown()
	fmt.Println(t4) // Output: False

	// And
	result := t1.And(t2)
	fmt.Println(result) // Output: False

	// Or
	result = t1.Or(t2)
	fmt.Println(result) // Output: True

	// Xor
	result = t1.Xor(t2)
	fmt.Println(result) // Output: True

	// Nand
	result = t1.Nand(t2)
	fmt.Println(result) // Output: True

	// Nor
	result = t1.Nor(t2)
	fmt.Println(result) // Output: False

	// Xnor
	result = t1.Nxor(t2)
	fmt.Println(result) // Output: False

	// Int
	fmt.Println(t1.Int()) // Output: 1

	// String
	fmt.Println(t1.String()) // Output: True
}
```

## Truth Tables

The truth table for the three-valued logic system is shown here. It's a mathematical table used in logic—specifically in connection with three-valued algebra, three-valued functions, and propositional calculus—which sets out the functional values of logical expressions on each of their functional arguments, that is, for each combination of values taken by their logical variables.

```shell
 Truth Tables of Three-valued logic
 (T=True, N=Unknown, F=False)

 NA   - Not
 MA   - Modus Ponens Absorption
 LA   - Law of Absorption
 IA   - Implication Absorption

 AND  - Logical AND
 OR   - Logical OR
 XOR  - Exclusive OR

 NAND - Logical not AND
 NOR  - Logical not OR
 NXOR - Logical not XOR

 IMP  - Implication in Lukasevich's Logic
 EQ   - If and only if
 MIN  - Minimum

 NIMP - NOT IMP
 NEQ  - NOT EQ
 MAX  - Maximum

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

## Explanation

Here's an explanation of some key parts of this package:

  - The Trit type is defined as an int8. This is a signed 8-bit integer, which means it can hold values between -128 and 127.

  - The Trit type is used to represent a "trinary digit," which can take on three states: False, Unknown, and True.

  - Package defineds various methods on the Trit type that allow you to perform operations on trinary digits, including determining if a trinary digit represents False, Unknown, or True, setting the value of a trinary digit based on an integer, normalizing a trinary digit, converting a trinary digit to an integer or a string.

  - The Trit type is defined as an alias for int8, and it can take one of three constants as its value: False, Unknown, or True. False corresponds to any negative number (including -1), Unknown corresponds to 0, and True corresponds to any positive number (including 1).

  - Provides methods default state (like Default, TrueIfUnknown, FalseIFUnknown etc.,) that check if the value of a Trit is Unknown and change its value based on the method called. For instance, TrueIfUnknown will set the Trit to True if its current value is Unknown.

  - There are some methods (like IsFalse, IsUnknown, IsTrue, etc.,) that check the state of a Trit and return a boolean indicating if the Trit is in the corresponding state.

  - Some methods perform various operations like assigning a value to a Trit (Set), returning the normalized value of a Trit (Val), normalizing the Trit in place (Norm), getting the integer representation of the Trit (Int), and getting the string representation of the Trit (String).

  - There are several methods for performing logic operations on Trit values including Not, And, Or, Xor, Nand,  Nor etc. These methods implement trinary logic versions of their respective boolean logic operations.

  - The logic operation methods follow truth tables for ternary logic which are defined in the package comment section.

## A real example of use

This package can be used, for example, to improve methods of updating information. Let's look at the following example:

```go
package main

import (
	"fmt"
)

// The defaultEnabled variable is used to set the default value
// for the Enabled field of the Config struct.
const defaultEnabled = true

// Config is a some configuration struct.
type Config struct {
	Name    string
	Port    int
	Eanbled bool
}

// Service is a some service struct.
type Service struct {
	name    string
	port    int
	enabled bool
}

// Configure executes the service configuration without changing
// the previously set parameters if the value for a specific
// field is not specified.
//
// If the field is left blank, set the default value.
func (s *Service) Configure(config Config) {
	// We have no problem determining if the Name field was passed or
	// if it is empty and we have to set it as the default value.
	if config.Name != "" {
		s.name = config.Name
	}

	if s.name == "" {
		s.name = "default"
	}

	// Everything is also fine for the Port parameter.
	if config.Port != 0 {
		s.port = config.Port
	}

	if s.port == 0 {
		s.port = 8080
	}

	// PROBLEM HERE:
	// But with bool the situation is more complicated.
	// Not sure if we passed False or just left the field blank.
	// Also, we can't figure out if the field is empty and should
	// be set by default, or if it's just set to False.
	if config.Eanbled != false {
		s.enabled = config.Eanbled
	}

	if s.enabled == false {
		s.enabled = defaultEnabled
	}
}

func main() {
	// Default service.
	s := &Service{}
	s.Configure(Config{Name: "default", Port: 8081, Eanbled: true})
	fmt.Println(s) // &{default 8081 true}

	// Change port only.
	s.Configure(Config{Port: 8082})
	fmt.Println(s) // &{default 8082 true}

	// PROBLEM HERE:
	// Change enabled status only.
	s.Configure(Config{Eanbled: false})
	fmt.Println(s) // &{default 8082 true}  --- not changed !!!

	// Note:
	// In this case, you need to write a configurator to which all
	// configuration fields (or at least bool fields) are always transferred.
}
```

As you can see, Boolean logic is not always enough even in such simple tasks.

The following example solves this problem:

```go
package main

import (
	"fmt"

	"github.com/goloop/trit"
)

// The defaultEnabled variable is used to set the default value
// for the Enabled field of the Config struct.
const defaultEnabled = true

// Config is a some configuration struct.
type Config struct {
	Name    string
	Port    int
	Eanbled trit.Trit // use three-valued logic
}

// Service is a some service struct.
type Service struct {
	name    string
	port    int
	enabled trit.Trit // use three-valued logic
}

// Configure executes the service configuration without changing
// the previously set parameters if the value for a specific
// field is not specified.
//
// If the field is left blank, set the default value.
func (s *Service) Configure(config Config) {
	// ... some code ...

	// Now we can easily determine if a field was passed as
	// a configuration update and if it is defined.
	if !config.Eanbled.IsUnknown() {
		s.enabled = config.Eanbled
	}

    // It looks like this: if the trit object is not defined,
    // set it to the default value, otherwise - do not touch it.
	trit.Default(&s.enabled, defaultEnabled)
}

// IsEnabled returns the current status of the service.
func (s *Service) IsEnabled() bool {
	return s.enabled.IsTrue()
}

func main() {
	// ... some code ...

	// Change enabled status only.
	s.Configure(Config{Eanbled: -1})
	fmt.Println(s) // &{default 8082 False} --- now it works !!!

	// Next code outputs: Service is disabled
	if s.IsEnabled() {
		fmt.Println("Service is enabled")
	} else {
		fmt.Println("Service is disabled")
	}
}
```

This is a very simple and demonstrative example of using three-valued logic in everyday life.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.