package main

import (
	"fmt"
)

func IsValid(s string) bool {
	stack := make([]rune, 0)

	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if opening, exists := mapping[char]; exists {
			var topElement rune
			if len(stack) == 0 {
				fmt.Println("Entrou no primeiro if")
				topElement = '#'
			} else {
				topElement = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			if topElement != opening {
				return false
			}
		} else {
			stack = append(stack, char)
		}

	}

	return len(stack) == 0
}

func main() {
	s := "{[]}"
	result := IsValid(s)
	fmt.Printf("A string '%s' é válida? %t\n", s, result)
}
