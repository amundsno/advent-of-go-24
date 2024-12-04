package day02

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"os"
)

func SolveFirst(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}
	defer file.Close()

	reports, err := utils.ReadRowsToSlices(file, " ")
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}

	rules := []Rule{new(DifferenceRule), new(DirectionalRule)}
	result := CountSafeReports(reports, rules)
	fmt.Printf("The number of safe reports are: %v\n", result)

}

func IsSafe(report []int, rules []Rule) bool {
	for _, rule := range rules {
		if !rule.Evaluate(report) {
			return false
		}
	}
	return true
}

func CountSafeReports(reports [][]int, rules []Rule) int {
	count := 0
	for _, report := range reports {
		if IsSafe(report, rules) {
			count++
		}
	}
	return count
}

type Rule interface {
	Evaluate(list []int) bool
}

type DifferenceRule struct{}

func (r DifferenceRule) Evaluate(list []int) bool {
	const minDifference, maxDifference = 1, 3
	for i := 0; i < len(list)-1; i++ {
		difference := list[i] - list[i+1]
		if difference < 0 {
			difference = -difference
		}
		if difference < minDifference || difference > maxDifference {
			return false
		}
	}
	return true
}

type DirectionalRule struct{}

func (r DirectionalRule) Evaluate(list []int) bool {
	if len(list) < 2 {
		return false
	}
	isIncreasing := list[1] > list[0]
	for i := 1; i < len(list)-1; i++ {
		if (isIncreasing && list[i] > list[i+1]) || (!isIncreasing && list[i] < list[i+1]) {
			return false
		}
	}
	return true
}
