package day02

import "testing"

func TestDifferenceRule(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"in range positive", []int{1, 2, 4, 7}, true},
		{"in range negative", []int{7, 4, 2, 1}, true},
		{"out of range positive", []int{1, 2, 4, 8}, false},
		{"out of range negative", []int{8, 4, 2, 1}, false},
		{"no change", []int{1, 1, 1, 1}, false},
	}

	dr := new(DifferenceRule)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := dr.Evaluate(tc.input)
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestDirectionalRule(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"increasing", []int{1, 2, 4, 7}, true},
		{"decreasing", []int{7, 4, 2, 1}, true},
		{"variable", []int{1, 2, 4, 3}, false},
	}

	dr := new(DirectionalRule)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := dr.Evaluate(tc.input)
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
