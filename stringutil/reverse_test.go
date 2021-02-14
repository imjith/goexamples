package stringutil

import "testing"

func TestReverse(t *testing.T) {
	for _, c := range []struct {
		in, ex string
	}{
		{"Hello World", "dlroW olleH"},
		{"Go Lang", "gnaL oG"},
		{"", ""},
	} {
		actual := Reverse(c.in)
		if c.ex != actual {
			t.Errorf("Reverse(%q), Expected %q, Actual %q", c.in, c.ex, actual)
		}
	}
}
