package trit

import (
	"encoding/json"
	"testing"
)

// isCanonical reports whether v is exactly one of the three canonical states.
// Every conversion/parse path must land on a canonical value; a non-canonical
// leak would be a correctness bug.
func isCanonical(v Trit) bool {
	return v == False || v == Unknown || v == True
}

// FuzzUnmarshalJSON feeds arbitrary bytes to the JSON decoder. The invariants:
// it must never panic, and on success the value must be canonical.
func FuzzUnmarshalJSON(f *testing.F) {
	for _, s := range []string{
		"null", "true", "false", "1", "-1", "0", "3.14", "-2.7",
		"", "  ", "abc", `"true"`, "[1,2]", "{}", "1e3", "99999999999999999999",
	} {
		f.Add([]byte(s))
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		var v Trit
		if err := json.Unmarshal(data, &v); err == nil {
			if !isCanonical(v) {
				t.Fatalf("Unmarshal(%q) produced non-canonical %d", data, int8(v))
			}
		}
	})
}

// FuzzParseTrit ensures ParseTrit never panics and always yields a canonical
// value; on error it must return Unknown (documented behaviour).
func FuzzParseTrit(f *testing.F) {
	for _, s := range []string{
		"true", "False", " UNKNOWN ", "maybe", "1", "-1", "0", "", "x", "yes",
	} {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		v, err := ParseTrit(s)
		if !isCanonical(v) {
			t.Fatalf("ParseTrit(%q) produced non-canonical %d", s, int8(v))
		}
		if err != nil && v != Unknown {
			t.Fatalf("ParseTrit(%q) error case returned %s, want Unknown", s, v)
		}
	})
}

// FuzzDefineInt checks the reflection-free integer conversion is total and
// sign-correct for any int64.
func FuzzDefineInt(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(1))
	f.Add(int64(-1))
	f.Add(int64(9223372036854775807))
	f.Add(int64(-9223372036854775808))

	f.Fuzz(func(t *testing.T, n int64) {
		v := Define(n)
		if !isCanonical(v) {
			t.Fatalf("Define(%d) non-canonical: %d", n, int8(v))
		}
		want := Unknown
		switch {
		case n > 0:
			want = True
		case n < 0:
			want = False
		}
		if v != want {
			t.Fatalf("Define(%d) = %s, want %s", n, v, want)
		}
	})
}

// FuzzScan ensures the sql.Scanner path never panics on arbitrary strings and
// stays canonical.
func FuzzScan(f *testing.F) {
	for _, s := range []string{"true", "false", "null", "maybe", "", "42", "!!"} {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		var v Trit
		if err := v.Scan(s); err == nil && !isCanonical(v) {
			t.Fatalf("Scan(%q) produced non-canonical %d", s, int8(v))
		}
	})
}
