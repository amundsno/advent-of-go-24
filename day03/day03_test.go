package day03

import (
	"reflect"
	"testing"
)

func TestSumValidMultiplications(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			"sample",
			"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			161,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := SumValidMultiplications(tc.input)
			if err != nil {
				t.Fatalf("unexpected error ocurred: %v", err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected: %v\ngot:%v\n", tc.expected, result)
			}
		})
	}
}

func TestExtractEnabledRegions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			"sample",
			"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))!?mul(3,14)!don't()mul(15,24)do()!mul(3,8)!",
			[]string{
				"xmul(2,4)&mul[3,7]!^don't()",
				"do()?mul(8,5))!?mul(3,14)!don't()",
				"do()!mul(3,8)!",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ExtractEnabledRegions(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected: %v\ngot:%v\n", tc.expected, result)
			}
		})
	}
}
