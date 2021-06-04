package matrix

import "fmt"

func ExamplePrintMatrixClockWisely() {
	PrintMatrixClockWisely([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})
	fmt.Println()
	PrintMatrixClockWisely([][]int{
		{1, 2, 3, 4},
	})
	fmt.Println()
	PrintMatrixClockWisely([][]int{
		{1},
		{2},
		{3},
		{4},
	})
	fmt.Println()
	PrintMatrixClockWisely([][]int{
		{1},
	})
	fmt.Println()
	PrintMatrixClockWisely([][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	})

	// Output:
	// 1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10 (:
	// 1 2 3 4 (:
	// 1 2 3 4 (:
	// 1 (:
	// 1 2 4 6 8 7 5 3 (:
}
