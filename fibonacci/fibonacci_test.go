package fibonacci

import (
	"testing"
)

func TestFibonacciRecursive(t *testing.T) {
	t.Run("test sucess", func(t *testing.T) {
		got := FibonacciRecursive(20)
		want := 6765

		if got != want {
			t.Fatalf("want %d, but got %d", want, got)
		}
	})

	t.Run("test fail", func(t *testing.T) {
		got := FibonacciRecursive(20)
		want := 6760

		if got == want {
			t.Fatalf("want %d, but got %d", want, got)
		}
	})

}

func TestFibonacciDynamic(t *testing.T) {
	t.Run("test sucess", func(t *testing.T) {
		got := FibonacciDynamic(20)
		want := 6765

		if got != want {
			t.Fatalf("want %d, but got %d", want, got)
		}
	})

	t.Run("test fail", func(t *testing.T) {
		got := FibonacciDynamic(20)
		want := 6760

		if got == want {
			t.Fatalf("want %d, but got %d", want, got)
		}
	})

}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciRecursive(35)
	}
}

func BenchmarkFibonacciDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibonacciRecursive(35)
	}
}
