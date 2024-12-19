package day07

import "testing"

func TestIsReachable(t *testing.T) {
	tests := []struct {
		name      string
		target    int
		chain     []int
		operators []Operator
		expected  bool
	}{
		{"sample 1 - TRUE", 190, []int{10, 19}, []Operator{ADD, MUL}, true},
		{"sample 2 - TRUE", 3267, []int{81, 40, 27}, []Operator{ADD, MUL}, true},
		{"sample 3 - TRUE", 292, []int{11, 6, 16, 20}, []Operator{ADD, MUL}, true},
		{"sample 4 - FALSE", 83, []int{17, 5}, []Operator{ADD, MUL}, false},
		{"sample 5 - FALSE", 192, []int{17, 8, 14}, []Operator{ADD, MUL}, false},
		{"sample 7 - TRUE", 156, []int{15, 6}, []Operator{ADD, MUL, CONC}, true},
		{"sample 8 - FALSE", 156, []int{15, 6}, []Operator{ADD, MUL}, false},
		{"sample 9 - TRUE", 7290, []int{6, 8, 6, 15}, []Operator{ADD, MUL, CONC}, true},
		{"sample 10 - TRUE", 192, []int{17, 8, 14}, []Operator{ADD, MUL, CONC}, true},
		{"sample 11 - FALSE", 21037, []int{9, 7, 18, 13}, []Operator{ADD, MUL, CONC}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsReachable(tc.target, tc.chain, tc.operators)
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
