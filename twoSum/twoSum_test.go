package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		got := TwoSum1(nums, 9)
		want := []int{0, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{3, 2, 4}
		got := TwoSum1(nums, 6)
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("Example 3", func(t *testing.T) {
		nums := []int{3, 3}
		got := TwoSum1(nums, 6)
		want := []int{0, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestTwoSum2(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		got := TwoSum2(nums, 9)
		want := []int{0, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{3, 2, 4}
		got := TwoSum2(nums, 6)
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("Example 3", func(t *testing.T) {
		nums := []int{3, 3}
		got := TwoSum2(nums, 6)
		want := []int{0, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

var numsBench = []int{2, 7, 11, 15, 3, 5, 8, 10, 1, 6, 12, 4, 9, 14, 18, 21, 25, 17}
var targetBench = 33

func BenchmarkTwoSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {

		_ = TwoSum1(numsBench, targetBench)
	}
	// BenchmarkTwoSum1-12    	20118448	        97.84 ns/op	      16 B/op	       1 allocs/op
}

func BenchmarkTwoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TwoSum2(numsBench, targetBench)
	}

	// BenchmarkTwoSum2-12    	  975079	      1644 ns/op	     910 B/op	       3 allocs/op
}
