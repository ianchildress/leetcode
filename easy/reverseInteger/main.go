package reverseInteger

import "math"

func reverse(x int) int {
	var out int

	for x != 0 {
		// select the least significant digit of x and assign it to y
		y := x % 10
		// shift the current value of out over a column and add y as the new one's column
		out = out*10 + y
		// remove the least significant digit from x
		x = x / 10
	}

	// check if out of bounds
	if float64(out) > math.Pow(2, 31) || float64(out) < math.Pow(-2, 31) {
		return 0
	}

	return out
}
