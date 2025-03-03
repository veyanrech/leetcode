package arrays

import "testing"

func TestCountSubarrays(t *testing.T) {
	testCase := []int{2, 5, 1, 4, 3, 6}
	k := 1
	expected := 3

	result := countSubarrays(testCase, k)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
