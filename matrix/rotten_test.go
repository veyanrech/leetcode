package matrix

import "testing"

func TestRunRotten(t *testing.T) {
	test_case := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	result := orangesRotting(test_case)
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}
