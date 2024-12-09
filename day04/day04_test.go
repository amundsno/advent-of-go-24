package day04

import "testing"

const content = `XMAS
X...
M...
A...
S...
X...
.M..
..A.
...S
...X
..M.
.A..
S...
...X
MAS.`

const contentRev = `SAMX
S...
A...
M...
X...
S...
.A..
..M.
...X
...S
..A.
.M..
X...
...S
AMX.`

const contentWidth = 5
const contentHorizontalStep = 1
const contentVerticalStep = contentWidth
const contentDiagonalRightStep = contentWidth + 1
const contentDiagonalLeftStep = contentWidth - 1

func TestIsWord(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		startIndex int
		step       int
		expected   bool
	}{
		{"horizontal", content, 0, contentHorizontalStep, true},
		{"horizontal reverse", contentRev, 0, contentHorizontalStep, true},
		{"vertical", content, contentWidth * 1, contentVerticalStep, true},
		{"vertical reverse", contentRev, contentWidth * 1, contentVerticalStep, true},
		{"diagonal right", content, contentWidth * 5, contentDiagonalRightStep, true},
		{"diagonal right reverse", contentRev, contentWidth * 5, contentDiagonalRightStep, true},
		{"diagonal left", content, contentWidth*9 + 3, contentDiagonalLeftStep, true},
		{"diagonal left reverse", contentRev, contentWidth*9 + 3, contentDiagonalLeftStep, true},
		{"accross line break", content, contentWidth*13 + 3, contentHorizontalStep, false},
		{"accross line break reverse", contentRev, contentWidth*13 + 3, contentHorizontalStep, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsWord(tc.startIndex, tc.step, &tc.content)
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestGetGridWidth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"sample", sample, 11},
		{"content", content, 5},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			width := GetGridWidth(tc.input)
			if width != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, width)
			}
		})
	}
}

const sample = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestCountAllWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"sample", sample, 18},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			count := CountAllWords(tc.input)
			if count != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, count)
			}
		})
	}
}
