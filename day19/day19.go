package day19

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

// Recursive memoized DFS
func DesignPermutations(design string, options []string, memo map[string]int) int {
	if result, tried := memo[design]; tried {
		return result
	}
	if design == "" {
		return 1
	}
	permutations := 0
	for _, opt := range options {
		n := len(opt)
		if n <= len(design) && design[:n] == opt {
			permutations += DesignPermutations(design[n:], options, memo)
		}
	}
	memo[design] = permutations
	return permutations
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

	canDesign, permutations := 0, 0
	memo := make(map[string]int)
	for _, design := range designs {
		if p := DesignPermutations(design, options, memo); p > 0 {
			canDesign++
			permutations += p
		}
	}

	fmt.Printf("Part 01: %v\n", canDesign)
	fmt.Printf("Part 02: %v\n", permutations)
}
