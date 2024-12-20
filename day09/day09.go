package day09

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Solve(inputPath string) {
	input := utils.ReadFileToString(inputPath)
	blocks := ParseInput(input)
	Reorder(blocks)
	sum := ComputeChecksum(blocks)
	fmt.Printf("Part 01: %v (checksum of ordered blocks)\n", sum)
}

func ComputeChecksum(blocks []string) int {
	blocksInt, err := utils.SliceAtoi(blocks[:slices.Index(blocks, ".")])
	if err != nil {
		panic(fmt.Errorf("failed to truncate and convert to int '%v': %v", blocks, err))
	}
	sum := 0
	for i, val := range blocksInt {
		sum += i * val
	}
	return sum

}

func Reorder(blocks []string) {
	f, b := 0, len(blocks)-1
	for {
		for blocks[f] != "." {
			f++
		}
		for blocks[b] == "." {
			b--
			if b <= f {
				return
			}
		}
		blocks[f], blocks[b] = blocks[b], blocks[f]
	}
}

func ParseInput(input string) (blocks []string) {
	inputStr := strings.Split(input, "")
	inputInt, err := utils.SliceAtoi(inputStr)
	if err != nil {
		panic(fmt.Errorf("failed to convert []string (%v) to []int: %v", inputStr, err))
	}

	blocksSize := utils.SliceSum(inputInt)
	blocks = make([]string, blocksSize)

	id := 0
	bi := 0
	for i := range len(inputInt) {
		size := inputInt[i]
		var token string

		if i%2 == 0 {
			token = strconv.Itoa(id)
			id++
		} else {
			token = "."
		}
		for range size {
			blocks[bi] = token
			bi++
		}
	}
	return blocks
}
