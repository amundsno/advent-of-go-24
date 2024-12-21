package day09

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart01(inputPath string) {
	input := utils.ReadFileToString(inputPath)
	bits := ParseInputToBits(input)
	ReorderBits(bits)
	sum := ComputeBitChecksum(bits)
	fmt.Printf("Part 01: %v (checksum of ordered bits)\n", sum)
}

func ComputeBitChecksum(bits []string) int {
	blocksInt, err := utils.SliceAtoi(bits[:slices.Index(bits, ".")])
	if err != nil {
		panic(fmt.Errorf("failed to truncate and convert to int '%v': %v", bits, err))
	}
	sum := 0
	for i, val := range blocksInt {
		sum += i * val
	}
	return sum

}

func ReorderBits(bits []string) {
	f, b := 0, len(bits)-1
	for {
		for bits[f] != "." {
			f++
		}
		for bits[b] == "." {
			b--
			if b <= f {
				return
			}
		}
		bits[f], bits[b] = bits[b], bits[f]
	}
}

func ParseInputToBits(input string) (bits []string) {
	inputStr := strings.Split(input, "")
	inputInt, err := utils.SliceAtoi(inputStr)
	if err != nil {
		panic(fmt.Errorf("failed to convert []string (%v) to []int: %v", inputStr, err))
	}

	blocksSize := utils.SliceSum(inputInt)
	bits = make([]string, blocksSize)

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
			bits[bi] = token
			bi++
		}
	}
	return bits
}

func SolvePart02(inputPath string) {
	input := utils.ReadFileToString(inputPath)
	blocks := ParseInputToBlocks(input)
	ReorderBlocks(blocks)
	sum := ComputeBlockChecksum(blocks)
	fmt.Printf("Part 02: %v (checksum of ordered blocks)\n", sum)
}

type Block struct {
	pos, len int
}

func ComputeBlockChecksum(blocks []Block) int {
	// Divergent sequence:
	// id * (pos + (pos + 1) + (pos + 2) + ... + (pos + len - 1))
	// id * (pos * len + (1 + 2 + ... + len - 1))
	// id * (pos * len + (len-1)*(len-1+1)/2)
	// id * (pos * len + (len-1)*(len)/2)
	// id * len * (2 * pos + len - 1) / 2
	sum, id := 0, 0
	for iFile := 0; iFile < len(blocks); iFile += 2 {
		l, p := blocks[iFile].len, blocks[iFile].pos
		sum += id * l * (2*p + l - 1) / 2
		id++
	}
	return sum
}

func ReorderBlocks(blocks []Block) {
	for iFile := len(blocks) - 1; iFile >= 0; iFile -= 2 {
		file := &blocks[iFile]

		for iSpace := 1; iSpace < len(blocks); iSpace += 2 {
			space := &blocks[iSpace]

			if space.pos < file.pos && file.len <= space.len {
				file.pos = space.pos
				space.pos += file.len
				space.len -= file.len
			}
		}
	}
}

func ParseInputToBlocks(input string) []Block {
	blocks := make([]Block, len(input))
	pos := 0
	for i, ch := range input {
		l := int(ch - '0')
		blocks[i] = Block{pos, l}
		pos += l
	}
	return blocks
}
