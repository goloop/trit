// Package trit implements three-valued logic operations with False,
// Unknown, and True states.
//
// # Overview
//
// Trit (short for "trinary digit") is an information unit that can take three
// states: False (-1), Unknown (0), and True (1). It serves as the fundamental
// unit of trinary/ternary logic systems, with applications in:
//   - Database systems (SQL NULL handling)
//   - Logic circuits and digital design
//   - Decision systems with uncertainty
//   - Artificial intelligence and expert systems
//   - Configuration management
//
// Key Features
//   - Basic type conversion (from/to bool, int, float)
//   - Fundamental unary operations (NOT, MA, LA, IA)
//   - Binary logic operations (AND, OR, XOR, etc.)
//   - Extended operations (IMP, EQ, MIN, MAX)
//   - Thread-safe parallel operations for slices
//   - Full set of comparison and testing methods
//
// # Quick Start
//
// Basic usage example:
//
//	t1 := trit.True
//	t2 := trit.Unknown
//	result := t1.And(t2)      // Unknown
//	isTrue := result.IsTrue() // false
//
// # Type System
//
// The package implements type Trit as int8 with three main states:
//   - False:   Any negative value (-1 and below)
//   - Unknown: Zero (0)
//   - True:    Any positive value (1 and above)
//
// # Truth Tables
//
// The package implements the following truth tables for three-valued logic:
//
// 1. Unary Operations:
//
//   - NA (Not): Logical negation
//
//   - MA (Modus Ponens Absorption)
//
//   - LA (Law of Absorption)
//
//   - IA (Implication Absorption)
//
//     A  | NA      A  | MA      A  | LA      A  | IA
//     ----+----    ----+----    ----+----    ----+----
//     F  |  T      F  |  F      F  |  F      F  |  F
//     U  |  U      U  |  T      U  |  F      U  |  T
//     T  |  F      T  |  T      T  |  T      T  |  F
//
// 2. Basic Binary Operations:
//
//   - AND: Logical conjunction
//
//   - OR:  Logical disjunction
//
//   - XOR: Exclusive OR
//
//     A | B | AND       A | B |  OR       A | B | XOR
//     ---+---+------    ---+---+------    ---+---+------
//     F | F |  F        F | F |  F        F | F |  F
//     F | U |  F        F | U |  U        F | U |  U
//     F | T |  F        F | T |  T        F | T |  T
//     U | F |  F        U | F |  U        U | F |  U
//     U | U |  U        U | U |  U        U | U |  U
//     U | T |  U        U | T |  T        U | T |  U
//     T | F |  F        T | F |  T        T | F |  T
//     T | U |  U        T | U |  T        T | U |  U
//     T | T |  T        T | T |  T        T | T |  F
//
// 3. Negative Binary Operations:
//
//   - NAND: Logical NAND (NOT AND)
//
//   - NOR:  Logical NOR (NOT OR)
//
//   - NXOR: Logical XNOR (NOT XOR)
//
//     A | B | NAND      A | B | NOR       A | B | NXOR
//     ---+---+------    ---+---+------    ---+---+------
//     F | F |  T        F | F |  T        F | F |  T
//     F | U |  T        F | U |  U        F | U |  U
//     F | T |  T        F | T |  F        F | T |  F
//     U | F |  T        U | F |  U        U | F |  U
//     U | U |  U        U | U |  U        U | U |  U
//     U | T |  U        U | T |  F        U | T |  U
//     T | F |  T        T | F |  F        T | F |  F
//     T | U |  U        T | U |  F        T | U |  U
//     T | T |  F        T | T |  F        T | T |  T
//
// 4. Extended Operations:
//
//   - IMP:  Implication (Lukasiewicz Logic)
//
//   - EQ:   Equivalence (If and only if)
//
//   - MIN:  Minimum value
//
//   - NIMP: Inverse implication
//
//   - NEQ:  Non-equivalence
//
//   - MAX:  Maximum value
//
//     A | B | IMP       A | B |  EQ       A | B | MIN
//     ---+---+------    ---+---+------    ---+---+------
//     F | F |  T        F | F |  T        F | F |  F
//     F | U |  T        F | U |  U        F | U |  F
//     F | T |  T        F | T |  F        F | T |  F
//     U | F |  U        U | F |  U        U | F |  F
//     U | U |  T        U | U |  U        U | U |  U
//     U | T |  T        U | T |  U        U | T |  U
//     T | F |  F        T | F |  F        T | F |  F
//     T | U |  U        T | U |  U        T | U |  U
//     T | T |  T        T | T |  T        T | T |  T
//
//     A | B | NIMP      A | B | NEQ       A | B | MAX
//     ---+---+------    ---+---+------    ---+---+------
//     F | F |  F        F | F |  F        F | F |  F
//     F | U |  F        F | U |  U        F | U |  U
//     F | T |  F        F | T |  T        F | T |  T
//     U | F |  U        U | F |  U        U | F |  U
//     U | U |  F        U | U |  U        U | U |  U
//     U | T |  F        U | T |  U        U | T |  T
//     T | F |  T        T | F |  T        T | F |  T
//     T | U |  U        T | U |  U        T | U |  T
//     T | T |  F        T | T |  F        T | T |  T
//
// # Thread Safety
//
// All operations in this package are thread-safe. The parallel operations
// (All, Any, Known) use goroutines for processing large slices and include
// proper synchronization mechanisms.
//
// # Performance Considerations
//
// The package optimizes performance by:
//   - Using int8 as the underlying type
//   - Implementing efficient parallel processing for slice operations
//   - Providing direct value access methods
//   - Minimizing memory allocations
//
// For more examples and detailed API documentation, see the individual method
// documentation and the package examples.
package trit
