package main

import "testing"

func TestPalindromo(t *testing.T) {
	t.Run("Is a palindrome", func(t *testing.T) {
		got := isPalindrome("arara")
		want := true

		assertCorretMessage(t, got, want)
	})

	t.Run("Is not a palindrome", func(t *testing.T) {
		got := isPalindrome("gabriel")
		want := false

		assertCorretMessage(t, got, want)
	})
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPalindrome("arara")
	}
}

func BenchmarkPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPalindrome2("arara")
	}
}
func assertCorretMessage(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
