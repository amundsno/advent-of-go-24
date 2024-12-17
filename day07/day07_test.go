package day07

import "testing"

func TestIsReachable(t *testing.T) {
	tests := []struct {
		name        string
		goal, start int
		weights     []int
		expected    bool
	}{
		{"sample 1 - TRUE", 190, 10, []int{19}, true},
		{"sample 2 - TRUE", 3267, 81, []int{40, 27}, true},
		{"sample 3 - TRUE", 292, 11, []int{6, 16, 20}, true},
		{"sample 1 - FALSE", 83, 17, []int{5}, false},
		{"sample 2 - FALSE", 192, 17, []int{8, 14}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsReachable(tc.goal, tc.start, tc.weights)
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
