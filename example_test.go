package trit_test

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/goloop/trit/v2"
)

func ExampleAllSeq() {
	values := []trit.Trit{trit.True, trit.True, trit.False, trit.True}
	// AllSeq stops at the first non-True value without draining the rest.
	fmt.Println(trit.AllSeq(slices.Values(values)))
	// Output: False
}

func ExampleTrit_And() {
	a := trit.True
	b := trit.Unknown
	fmt.Println(a.And(b)) // conjunction with an unknown is unknown
	// Output: Unknown
}

func ExampleTrit_Not() {
	fmt.Println(trit.False.Not(), trit.Unknown.Not(), trit.True.Not())
	// Output: True Unknown False
}

func ExampleDefine() {
	fmt.Println(trit.Define(5), trit.Define(0), trit.Define(-3))
	// Output: True Unknown False
}

func ExampleParseTrit() {
	for _, s := range []string{"yes", "maybe", "off"} {
		v, _ := trit.ParseTrit(s)
		fmt.Println(v)
	}
	// Output:
	// True
	// Unknown
	// False
}

func ExampleConsensus() {
	fmt.Println(trit.Consensus(trit.True, trit.True, trit.True))
	fmt.Println(trit.Consensus(trit.True, trit.True, trit.Unknown))
	// Output:
	// True
	// Unknown
}

func ExampleMajority() {
	fmt.Println(trit.Majority(trit.True, trit.True, trit.False))
	// Output: True
}

func ExampleTrit_MarshalJSON() {
	type record struct {
		Active trit.Trit `json:"active"`
	}
	b, _ := json.Marshal(record{Active: trit.Unknown})
	fmt.Println(string(b))
	// Output: {"active":null}
}

func ExampleTrit_Compare() {
	fmt.Println(trit.False.Compare(trit.True))
	fmt.Println(trit.True.Compare(trit.True))
	fmt.Println(trit.True.Compare(trit.Unknown))
	// Output:
	// -1
	// 0
	// 1
}
