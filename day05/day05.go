package day05

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func ExtractPageOrder(rule string) (left, right int, err error) {
	order := strings.Split(rule, "|")
	if len(order) != 2 {
		return 0, 0, fmt.Errorf("invalid format: rule must contain exactly one '|' character")
	}
	left, err = strconv.Atoi(order[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid left number: %v", err)
	}
	right, err = strconv.Atoi(order[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid right number: %v", err)
	}
	return left, right, nil
}

func ComputeSortFunction(rules []string) func(a, b int) int {
	comesAfter := make(map[int]map[int]bool)
	for _, rule := range rules {
		left, right, err := ExtractPageOrder(rule)
		if err != nil {
			panic(fmt.Sprintf("failed to extract page order from rule: %v", err))
		}
		if comesAfter[left] == nil {
			comesAfter[left] = make(map[int]bool)
		}
		comesAfter[left][right] = true
	}
	return func(a, b int) int {
		if comesAfter[a][b] {
			return -1 // a < b
		} else if comesAfter[b][a] {
			return 1 // a > b
		}
		return 0
	}
}

func ReadFileToProblemComponents(inputPath string) (rules []string, orders [][]int) {
	content := utils.ReadFileToString(inputPath)
	contentParts := strings.Split(content, "\n\n")
	rules = strings.Split(contentParts[0], "\n")
	orders, err := utils.ReadRowsToSlices(strings.NewReader(contentParts[1]), ",")
	if err != nil {
		panic(err)
	}
	return rules, orders
}

func Solve(inputPath string) {
	rules, orders := ReadFileToProblemComponents(inputPath)
	cmpFunc := ComputeSortFunction(rules)

	sortedSum, unsortedSum := 0, 0
	for _, order := range orders {
		if slices.IsSortedFunc(order, cmpFunc) {
			// Part 01
			sortedSum += order[len(order)/2]
		} else {
			// Part 02
			slices.SortFunc(order, cmpFunc)
			unsortedSum += order[len(order)/2]
		}
	}
	fmt.Printf("Part 01 (sum of middle numbers in valid page orders): %v\n", sortedSum)
	fmt.Printf("Part 02 (sum of middle numbers after reordering invalid page orders): %v\n", unsortedSum)
}
