package matrix

import "fmt"

func PrintMatrixClockWisely(numbers [][]int) {
	if row := len(numbers); row > 0 {
		printMatrixClockWisely(numbers, row, len(numbers[0]))
	}
}

func printMatrixClockWisely(numbers [][]int, rows, columns int) {
	start := 0
	for start<<1 < columns && start<<1 < rows {
		printMatrixInCircle(numbers, rows, columns, start)
		start++
	}
	fmt.Print("(:")
}

func printMatrixInCircle(numbers [][]int, rows, columns, start int) {
	// 向右 向下 向左 向上
	if columns-start>>1 > 0 {
		for row, col := start, start; col < columns-start; col++ {
			fmt.Print(numbers[row][col], " ")
		}

	}
	// 大于一行
	if rows-start>>1 > 1 {
		for row, col := start+1, columns-start-1; row < rows-start; row++ {
			fmt.Print(numbers[row][col], " ")
		}
	}
	// 大于一列并且大于一行
	if (columns-start>>1 > 1) && (rows-start>>1 > 1) {
		for row, col := rows-start-1, columns-start-2; start <= col; col-- {
			fmt.Print(numbers[row][col], " ")
		}
	}
	// 大于一列并且大于两行
	if (columns-start>>1 > 1) && (rows-start>>1 > 2) {
		for row, col := rows-start-2, start; start < row; row-- {
			fmt.Print(numbers[row][col], " ")
		}
	}
}
