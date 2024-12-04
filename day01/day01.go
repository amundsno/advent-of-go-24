package day01

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"os"
	"sort"
)

func SolveFirst(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}
	defer file.Close()

	columns, err := utils.ReadColumnsToSlices(file, "   ", 2)
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}

	left, right := columns[0], columns[1]
	sum := CalculateSumOfLeastDistances(left, right)

	fmt.Printf("The sum of distances is: %v\n", sum)
}

func CalculateSumOfLeastDistances(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	distances := ComputeDistances(left, right)
	sum := utils.SumSlice(distances)
	return sum
}

func ComputeDistances(sortedList1, sortedList2 []int) []int {
	n := len(sortedList1)
	distances := make([]int, n)
	for i := range sortedList1 {
		distances[i] = sortedList1[i] - sortedList2[i]
		if distances[i] < 0 {
			distances[i] = -distances[i]
		}
	}
	return distances
}

func SolveSecond(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}
	defer file.Close()

	columns, err := utils.ReadColumnsToSlices(file, "   ", 2)
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}

	left, right := columns[0], columns[1]
	sum := CalculateSimilarityScore(left, right)

	fmt.Printf("The similarity score is: %v\n", sum)

}

func CalculateSimilarityScore(left, right []int) int {
	valueCountMap := CountValues(right)
	sum := 0
	for _, v := range left {
		sum += v * valueCountMap[v]
	}
	return sum
}

func CountValues(a []int) map[int]int {
	vcmap := make(map[int]int)

	for _, v := range a {
		vcmap[v]++
	}

	return vcmap
}
