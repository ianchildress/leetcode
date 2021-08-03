package main

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := "ab"
	want := 2
	if got := lengthOfLongestSubstring(s); got != want {
		t.Errorf("string: %s want: %v got: %v\n", s, want, got)
	}

	s = "abcabcbb"
	want = 3
	if got := lengthOfLongestSubstring(s); got != want {
		t.Errorf("string: %s want: %v got: %v\n", s, want, got)
	}

	s = "dvdf"
	want = 3
	if got := lengthOfLongestSubstring(s); got != want {
		t.Errorf("string: %s want: %v got: %v\n", s, want, got)
	}

	s = "bbbbb"
	want = 1
	if got := lengthOfLongestSubstring(s); got != want {
		t.Errorf("string: %s want: %v got: %v\n", s, want, got)
	}

	s = ""
	want = 0
	if got := lengthOfLongestSubstring(s); got != want {
		t.Errorf("string: %s want: %v got: %v\n", s, want, got)
	}

	s = "pwwkew"
	want = 3
	if got := lengthOfLongestSubstring(s); got != want {
		t.Errorf("string: %s want: %v got: %v\n", s, want, got)
	}
}
