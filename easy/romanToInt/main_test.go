package romanToInt

import "testing"

func Test_romanToInt(t *testing.T) {
	var s string
	var got, want int

	s = "III"
	want = 3
	got = romanToInt(s)
	if got != want {
		t.Errorf("string: %s want: %v got: %v", s, want, got)
	}

	s = "IV"
	want = 4
	got = romanToInt(s)
	if got != want {
		t.Errorf("string: %s want: %v got: %v", s, want, got)
	}

	s = "IX"
	want = 9
	got = romanToInt(s)
	if got != want {
		t.Errorf("string: %s want: %v got: %v", s, want, got)
	}

	s = "LVIII"
	want = 58
	got = romanToInt(s)
	if got != want {
		t.Errorf("string: %s want: %v got: %v", s, want, got)
	}

	s = "MCMXCIV"
	want = 1994
	got = romanToInt(s)
	if got != want {
		t.Errorf("string: %s want: %v got: %v", s, want, got)
	}
}
