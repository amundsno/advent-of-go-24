package day11

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve(inputPath string) {
	stonesStr := utils.ReadFileToString(inputPath)
	stoneStrSlice := strings.Split(stonesStr, " ")
	stones, err := utils.SliceAtoi(stoneStrSlice)
	if err != nil {
		panic(err)
	}

	count := CountStonesAfterBlinks(stones, 25)
	fmt.Printf("Part 01: %v (number of stones after 25 blinks)\n", count)

	count = CountStonesAfterBlinks(stones, 75)
	fmt.Printf("Part 02: %v (number of stones after 75 blinks)\n", count)
}

type MemoKey struct {
	stone, blinks int
}

func CountStonesAfterBlinks(stones []int, blinks int) int {
	memory := map[MemoKey]int{}

	var compute func(stone, blinks int) int
	compute = func(stone, blinks int) int {
		if blinks == 0 {
			return 1
		} else if value, exists := memory[MemoKey{stone, blinks}]; exists {
			return value
		}
		var count int
		if stone == 0 {
			count = compute(1, blinks-1)
		} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
			left, right := splitStone(stone)
			count = compute(left, blinks-1) + compute(right, blinks-1)
		} else {
			count = compute(stone*2024, blinks-1)
		}
		memory[MemoKey{stone, blinks}] = count
		return count
	}

	count := 0
	for _, stone := range stones {
		count += compute(stone, blinks)
	}

	return count
}

func splitStone(stone int) (left, right int) {
	numDigits := int(math.Floor(math.Log10(float64(stone))) + 1)
	left = stone / int(math.Pow10(numDigits/2))
	right = stone % int(math.Pow10(numDigits/2))
	return left, right
}

/*
func splitStoneAsDigits(stone int) (left, right int) {
	digits := []int{}
	for stone > 0 {
		digits = append([]int{stone % 10}, digits...) // Clever way to append in front
		stone /= 10
	}
	mid := len(digits) / 2
	for i, digit := range digits {
		if i < mid {
			left = left*10 + digit
		} else {
			right = right*10 + digit
		}
	}
	return left, right
}
*/
