package matrix

import "fmt"

func RunRotate() {
	matrix := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}

	matrix = rotate(matrix)
}

func rotate(matrix [][]int) [][]int {

	n := len(matrix) * len(matrix[0])
	rowLen := len(matrix[0])
	colLen := len(matrix)
	buf := make([]int, colLen)
	fmt.Println("rowLen:", rowLen, "colLen: ", colLen, " n: ", n, " buf: ", buf)

	for i := 0; i < n; i++ {

	}

	return matrix
}
