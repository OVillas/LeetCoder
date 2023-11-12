package main

import "testing"

func TestValidParentheses(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		s := "()"
		got := IsValid(s)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		s := "()[]{}"
		got := IsValid(s)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("example 1", func(t *testing.T) {
		s := "(]"
		got := IsValid(s)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("example 1", func(t *testing.T) {
		s := "{{[[(())]]}}"
		got := IsValid(s)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("example 1", func(t *testing.T) {
		s := "{{[[(())]]})}"
		got := IsValid(s)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
