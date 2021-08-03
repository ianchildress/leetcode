package romanToInt

func romanToInt(s string) int {
	var result int
	// important: roman numerals are read from left to right and from large to small. if we process from left to right and
	// we go from small to large then we know we need to subtract.

	lookup := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	// check for empty string
	if len(s) == 0 {
		return 0
	}

	// check for single character string
	if len(s) == 1 {
		return lookup[s]
	}

	// i'm choosing to do a look forward approach. it seems to have the simplest design.
	for i := 0; i < len(s); i++ {
		char := string(s[i])

		// check if this is the last character in the string. we could pull this out of the loop and calculate the
		// last character at the end for a performance increase but for readability/simplicity i have chosen to
		// include it in the loop.
		// this will always be a blind addition at the end because in our subtraction case we increment the index
		// so this will never get evaluated if the last characters are a subtraction.
		if i == len(s)-1 {
			result += lookup[char]
			continue
		}

		nextChar := string(s[i+1])

		// check to see if this character is larger or equal to the next character. if true, we add.
		// this handles the majority of our calculations
		if lookup[char] >= lookup[nextChar] {
			result += lookup[char]
			continue
		}

		// this must be a subtraction since we've ruled out all other possibilities.
		result += lookup[nextChar] - lookup[char]
		i++ // increment index since we did a look ahead
	}

	return result
}

func firstIteration(s string) int {
	var result int
	var lastChar, nextChar string
	var subset []string

	lookup := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	// find break points
	for i := 0; i < len(s); i++ {
		char := string(s[i])

		// check for repeat chars
		if char == lastChar {
			subset = append(subset, char)
			continue
		}

		// check if there is another character
		if i+1 < len(s) {
			nextChar = string(s[i+1])
		} else {
			nextChar = ""
		}

		switch {
		case char == "I" && nextChar == "V":
			result += 4
			i++
			continue
		case char == "I" && nextChar == "X":
			result += 9
			i++
			continue
		case char == "X" && nextChar == "L":
			result += 40
			i++
			continue
		case char == "X" && nextChar == "C":
			result += 90
			i++
			continue
		case char == "C" && nextChar == "D":
			result += 400
			i++
			continue
		case char == "C" && nextChar == "M":
			result += 900
			i++
			continue
		}

		result += lookup[char]
	}

	return result
}
