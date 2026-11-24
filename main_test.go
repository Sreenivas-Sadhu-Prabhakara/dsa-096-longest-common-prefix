package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"same", "same"}, "same"},
		{[]string{"alone"}, "alone"},
		{[]string{}, ""},
		{[]string{"ab", "a"}, "a"},
		{[]string{"", "abc"}, ""},
	}
	for _, c := range cases {
		if got := LongestCommonPrefix(c.in); got != c.want {
			t.Errorf("LongestCommonPrefix(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
