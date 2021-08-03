package main

import "fmt"

func main() {
	c := lengthOfLongestSubstring("dvdf")
	fmt.Println(c)
}

func lengthOfLongestSubstring(s string) int {
	var complete, current, winner []uint8

	// split looks for a repeat character in the given list. if the given character already exists in the list
	split := func(s uint8, list []uint8) ([]uint8, []uint8) {
		for i := range list {
			if s == list[i] {
				if len(list) == 1 {
					return list, list
				}

				return list, append(list[i+1:], s)
			}
		}

		return nil, append(list, s)
	}

	for i := range s {
		// check if this character exists in current substring
		complete, current =split(s[i], current)

		if len(complete)>len(winner)  {
			winner = complete
		}
	}

	if len(current) > len(winner) {
		winner = current
	}

	return len(winner)
}
