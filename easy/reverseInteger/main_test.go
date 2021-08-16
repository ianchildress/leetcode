package reverseInteger

import "testing"

func Test_reverse(t *testing.T) {
	var val,want int

	val = 123
	want = 321
	if got := reverse(val);got!=want{
		t.Errorf("val: %v want: %v got: %v",val,want,got)
	}


	val = -123
	want = -321
	if got := reverse(val);got!=want{
		t.Errorf("val: %v want: %v got: %v",val,want,got)
	}

	val = -1234
	want = -4321
	if got := reverse(val);got!=want{
		t.Errorf("val: %v want: %v got: %v",val,want,got)
	}

	val = 120
	want = 21
	if got := reverse(val);got!=want{
		t.Errorf("val: %v want: %v got: %v",val,want,got)
	}


	val = 0
	want = 0
	if got := reverse(val);got!=want{
		t.Errorf("val: %v want: %v got: %v",val,want,got)
	}


	val = 1534236469
	want = 0
	if got := reverse(val);got!=want{
		t.Errorf("val: %v want: %v got: %v",val,want,got)
	}


}
