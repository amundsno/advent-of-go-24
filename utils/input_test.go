package utils

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadRowsToSlices(t *testing.T) {
	tests := []struct {
		name        string
		delimiter   string
		input       string
		expected    [][]int
		shouldError bool
	}{
		{"basic case", " ", "1 2 3\n4 5 6\n", [][]int{{1, 2, 3}, {4, 5, 6}}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			rows, err := ReadRowsToSlices(reader, " ")
			if err != nil && !tc.shouldError {
				t.Fatalf("unexpected error: %v", err)
			}
			if err == nil && !reflect.DeepEqual(rows, tc.expected) {
				t.Errorf("expected: %v, got: %v", tc.expected, rows)
			}
		})
	}
}

func TestReadToString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic case", "hello\nthis is a\ttest", "hello\nthis is a\ttest"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			result, err := readToString(reader)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestReadToMap(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter string
		expected  map[int]map[int]string
	}{
		{"basic case", "a b\nc d", " ", map[int]map[int]string{0: {0: "a", 1: "b"}, 1: {0: "c", 1: "d"}}},
		{"no delimiter case", "ab\ncd", "", map[int]map[int]string{0: {0: "a", 1: "b"}, 1: {0: "c", 1: "d"}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			result, err := readToMap(reader, tc.delimiter)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
