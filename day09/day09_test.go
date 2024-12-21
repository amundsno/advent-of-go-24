package day09

import (
	"reflect"
	"strings"
	"testing"
)

func TestReorderBits(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"small sample", "0..111....22222", "022111222......"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			blocks := strings.Split(tc.input, "")
			ReorderBits(blocks)

			expected := strings.Split(tc.expected, "")
			if !reflect.DeepEqual(blocks, expected) {
				t.Errorf("expected: %v, got: %v", expected, blocks)
			}

		})
	}
}

func TestParseInputToBits(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"small sample", "12345", "0..111....22222"},
		{"big sample", "2333133121414131402", "00...111...2...333.44.5555.6666.777.888899"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			blocks := ParseInputToBits(tc.input)
			if !reflect.DeepEqual(strings.Join(blocks, ""), tc.expected) {
				t.Errorf("expected: %v, got: %v", tc.expected, blocks)
			}

		})
	}
}

func TestComputeBitChecksum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"big sample", "0099811188827773336446555566..............", 1928},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			blocks := strings.Split(tc.input, "")
			result := ComputeBitChecksum(blocks)
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, blocks)
			}
		})
	}
}
