package day02

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"os"
)

func readReports(inputPath string) [][]int {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}
	defer file.Close()

	reports, err := utils.ReadRowsToSlices(file, " ")
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}
	return reports
}

func SolveFirst(inputPath string) {
	reports := readReports(inputPath)
	result := CountSafeReports(reports, IsSafe)
	fmt.Printf("The number of safe reports are: %v\n", result)
}

func SolveSecond(inputPath string) {
	reports := readReports(inputPath)
	result := CountSafeReports(reports, IsSafeWithOneBadLevel)
	fmt.Printf("The number of safe reports when accepting one bad level are: %v\n", result)
}

func IsSafe(report []int, rules []Rule) bool {
	for _, rule := range rules {
		if !rule.Evaluate(report) {
			return false
		}
	}
	return true
}

func IsSafeWithOneBadLevel(report []int, rules []Rule) bool {
	for i := 0; i < len(report); i++ {
		var left, right []int
		left = append(left, report[:i]...)
		right = append(right, report[i+1:]...)
		if IsSafe(append(left, right...), rules) {
			return true
		}
	}
	return false
}

func CountSafeReports(reports [][]int, isSafeFunc func([]int, []Rule) bool) int {
	rules := []Rule{new(DifferenceRule), new(DirectionalRule)}
	count := 0
	for _, report := range reports {
		if isSafeFunc(report, rules) {
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
	initialDifference := list[0] - list[1]
	for i := 1; i < len(list)-1; i++ {
		if initialDifference*(list[i]-list[i+1]) < 0 {
			return false
		}
	}
	return true
}
