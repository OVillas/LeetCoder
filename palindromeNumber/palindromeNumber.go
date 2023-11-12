package main

import "strconv"

func IsPalindrome(x int) bool {
	y := strconv.Itoa(x)
	lenght := len(y)

	for i := 0; i < lenght/2; i++ {
		if y[i] != y[lenght-1-i] {
			return false
		}
	}

	return true
}
