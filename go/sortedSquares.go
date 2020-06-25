package sortedsquares

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sq(x int) int {
	return x * x
}

func runSortedSquares__tail(input []int, accum []int) []int {
	if len(input) == 0 {
		return accum
	}

	first := input[0]
	lastIndex := len(input) - 1
	last := input[lastIndex]

	if abs(first) > last {
		return runSortedSquares__tail(input[1:], append([]int{sq(first)}, accum...))
	}

	return runSortedSquares__tail(input[0:lastIndex], append([]int{sq(last)}, accum...))
}

func runSortedSquares__goto(input []int, accum []int) []int {
Entry:

	if len(input) == 0 {
		return accum
	}

	first := input[0]
	lastIndex := len(input) - 1
	last := input[lastIndex]

	if abs(first) > last {
		accum = append([]int{sq(first)}, accum...)
		input = input[1:]
		goto Entry
	}

	accum = append([]int{sq(last)}, accum...)
	input = input[0:lastIndex]
	goto Entry
}

// SortedSquares__tail creates a sorted list of squares from the input
func SortedSquares__tail(input []int) []int {
	return runSortedSquares__tail(input, []int{})
}

// SortedSquares__tail creates a sorted list of squares from the input
func SortedSquares__goto(input []int) []int {
	return runSortedSquares__goto(input, []int{})
}

// SortedSquares__recursive creates a sorted list of squares from the input
func SortedSquares__recursive(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}

	first := input[0]
	lastIndex := len(input) - 1
	last := input[lastIndex]

	if abs(first) > last {
		return append(
			SortedSquares__recursive(input[1:]),
			sq(first),
		)
	}

	return append(
		SortedSquares__recursive(input[0:lastIndex]),
		sq(last),
	)
}
