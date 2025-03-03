package graph

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
	// bob := 3
	amount := []int{-5644, -6018, 1188, -8502}
	result := mostProfitablePath(edges, 3, amount)
	fmt.Println(result) // Output: 110
}
