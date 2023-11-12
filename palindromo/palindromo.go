package main

import "fmt"

func isPalindrome(str string) bool {
	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

func isPalindrome2(str string) bool {
	size := len(str) - 1

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[size-i] {
			return false
		}
	}

	return true
}

func main() {
	word := "arara"
	if isPalindrome(word) {
		fmt.Println("É palindromo!")
	} else {
		fmt.Println("Não é um palindromo!")
	}
}
