package utils

import (
	"testing"
)

func TestSumSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"basic case", []int{1, 2, 3}, 6},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sum := SumSlice(tc.input)
			if sum != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, sum)
			}
		})
	}
}
