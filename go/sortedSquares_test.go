package sortedsquares

import (
	"testing"
)

func BenchmarkSortedSquares__goto(b *testing.B) {
	input := []int{-3, -2, 0, 1, 5}

	for i := 0; i < b.N; i++ {
		SortedSquares__goto(input)
	}
}

func BenchmarkSortedSquares__tail(b *testing.B) {
	input := []int{-3, -2, 0, 1, 5}

	for i := 0; i < b.N; i++ {
		SortedSquares__tail(input)
	}
}

func BenchmarkSortedSquares__recursive(b *testing.B) {
	input := []int{-3, -2, 0, 1, 5}

	for i := 0; i < b.N; i++ {
		SortedSquares__recursive(input)
	}
}

// func TestSortedSquares(t *testing.T) {
// 	for _, c := range []struct {
// 		in, want []int
// 	}{
// 		{[]int{}, []int{}},
// 		{[]int{1, 2, 3}, []int{1, 4, 9}},
// 		{[]int{-5, -3, -2}, []int{4, 9, 25}},
// 		{[]int{-3, -2, 0, 1, 5}, []int{0, 1, 4, 9, 25}},
// 	} {
// 		got := SortedSquares(c.in)

// 		if len(got) != len(c.want) {
// 			t.Errorf("SortedSquares(%v) == %v, want %v", c.in, got, c.want)
// 		}

// 		for index, value := range got {
// 			if value != c.want[index] {
// 				t.Errorf("SortedSquares(%v) == %v, want %v", c.in, got, c.want)
// 			}
// 		}
// 	}
// }
