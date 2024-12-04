package day01

import (
	"reflect"
	"testing"
)

func TestComputeDistances(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   []int
		expected []int
	}{
		{"basic case", []int{0, 2, 6}, []int{1, 4, 9}, []int{1, 2, 3}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			distances := ComputeDistances(tc.input1, tc.input2)
			if !reflect.DeepEqual(distances, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, distances)
			}
		})
	}
}

func TestCountValues(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected map[int]int
	}{
		{"basic case", []int{1, 1, 3, 5}, map[int]int{1: 2, 3: 1, 5: 1}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			vcmap := CountValues(tc.input)
			if !reflect.DeepEqual(vcmap, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, vcmap)
			}
		})
	}
}
