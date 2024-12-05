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

func TestIsSafe(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"safe decreasing", []int{7, 6, 4, 2, 1}, true},
		{"safe increasing", []int{1, 3, 6, 7, 9}, true},
		{"unsafe difference", []int{1, 2, 7, 8, 9}, false},
		{"unsafe difference", []int{9, 7, 6, 2, 1}, false},
		{"unsafe direction change", []int{1, 3, 2, 4, 5}, false},
		{"unsafe no direction", []int{8, 6, 4, 4, 1}, false},
	}
	rules := []Rule{new(DifferenceRule), new(DirectionalRule)}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsSafe(tc.input, rules)
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestIsSafeWithOneBadLevel(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"safe decreasing without removing", []int{7, 6, 4, 2, 1}, true},
		{"safe increasing without removing", []int{1, 3, 6, 7, 9}, true},
		{"unsafe differences", []int{1, 2, 7, 8, 9}, false},
		{"unsafe differences", []int{9, 7, 6, 2, 1}, false},
		{"safe by removing report[1]", []int{1, 3, 2, 4, 5}, true},
		{"safe by removing report[2]", []int{8, 6, 4, 4, 1}, true},
	}
	rules := []Rule{new(DifferenceRule), new(DirectionalRule)}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsSafeWithOneBadLevel(tc.input, rules)
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
