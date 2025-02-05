package day19

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

// Recursive memoized DFS
func CanDesign(design string, options []string, memo map[string]bool) bool {
	if result, tried := memo[design]; tried {
		return result
	}
	if design == "" {
		return true
	}
	for _, opt := range options {
		n := len(opt)
		if n <= len(design) && design[:n] == opt && CanDesign(design[n:], options, memo) {
			memo[design] = true
			return true
		}
	}
	memo[design] = false
	return false
}

func ParseInput(filepath string) (options, designs []string) {
	content := utils.ReadFileToString(filepath)
	parts := strings.Split(content, "\n\n")
	options = strings.Split(parts[0], ", ")
	designs = strings.Split(parts[1], "\n")
	return
}

func Solve(filepath string) {
	options, designs := ParseInput(filepath)

	count := 0
	for _, design := range designs {
		memo := make(map[string]bool)
		if CanDesign(design, options, memo) {
			count++
		}
	}

	fmt.Printf("Part 01: %v\n", count)
}
