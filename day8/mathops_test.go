package day8

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if expected != 5 {
		t.Errorf("Expected %d, got %d", expected, result)
	}

}
func TestMax(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{9, 5, 1}, 9},
		{[]int{-1, -5, -2}, -1},
	}
	for _, tc := range tests {
		result := Max(tc.input)
		if result != tc.expected {
			t.Errorf("for input %v , for expected %d , got %d", tc.input, tc.expected, result)
		}
	}

}
