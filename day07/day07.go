package day07

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func Solve(inputPath string) {
	rows := utils.ReadFileToRows(inputPath)

	sum := 0
	for _, row := range rows {
		goal, start, weights, err := getComponents(row)
		if err != nil {
			panic(err)
		}
		if IsReachable(goal, start, weights) {
			sum += goal
		}
	}

	fmt.Printf("Part 01: %v\n", sum)
}

func getComponents(row string) (goal, start int, weights []int, err error) {
	goal, err = strconv.Atoi(strings.Split(row, ":")[0])
	if err != nil {
		return goal, start, weights, fmt.Errorf("failed to extract goal from '%v': %v", row, err)
	}

	sNumbers := strings.Split(row, " ")[1:]
	var iNumbers []int
	for _, sNum := range sNumbers {
		iNum, err := strconv.Atoi(sNum)
		if err != nil {
			return goal, start, weights, fmt.Errorf("failed to convert '%v' to integer: %v", sNum, err)
		}
		iNumbers = append(iNumbers, iNum)
	}

	start = iNumbers[0]
	weights = iNumbers[1:]
	return goal, start, weights, nil
}

func IsReachable(goal, start int, weights []int) bool {
	if len(weights) == 0 {
		return false
	}

	add := start + weights[0]
	mul := start * weights[0]

	if add == goal || mul == goal {
		return true
	}

	var reachableFromAdd, reachableFromMul bool
	if add < goal {
		reachableFromAdd = IsReachable(goal, add, weights[1:])
	}
	if mul < goal {
		reachableFromMul = IsReachable(goal, mul, weights[1:])
	}

	return reachableFromAdd || reachableFromMul
}
