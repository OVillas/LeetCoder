package main

import "testing"

func TestIsPalindrome(t *testing.T) {

	t.Run("Positive numbers", func(t *testing.T) {
		got := IsPalindrome(121)
		want := true

		if got != want {
			t.Errorf("got %t, but want %t", got, want)
		}
	})

	t.Run("Negative numbers", func(t *testing.T) {
		got := IsPalindrome(-121)
		want := false

		if got != want {
			t.Errorf("got %t, but want %t", got, want)
		}
	})

	t.Run("Giant numbers", func(t *testing.T) {
		got := IsPalindrome(2147447412)
		want := true

		if got != want {
			t.Errorf("got %t, but want %t", got, want)
		}
	})

	t.Run("Giant negative numbers", func(t *testing.T) {
		got := IsPalindrome(-2147447412)
		want := false

		if got != want {
			t.Errorf("got %t, but want %t", got, want)
		}
	})

}
