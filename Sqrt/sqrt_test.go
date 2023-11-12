package sqrt

import "testing"

func TestSqrt(t *testing.T) {
	t.Run("Número possuidor de raiz quadrada inteira", func(t *testing.T) {
		got := MySqrt(4)
		want := 2

		if got != want {
			t.Errorf("got %d, but  want %d", got, want)
		}
	})

	t.Run("Número não possuidor de raiz quadrada inteira", func(t *testing.T) {
		got := MySqrt(8)
		want := 2

		if got != want {
			t.Errorf("got %d, but  want %d", got, want)
		}
	})
}
