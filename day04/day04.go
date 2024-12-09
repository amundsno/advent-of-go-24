package day04

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(inputPath string) string {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}
	defer file.Close()

	content, err := utils.ReadToString(file)
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}
	return content
}

func SolveFirst(inputPath string) {
	content := readInput(inputPath)
	count := CountAllWords(content)
	fmt.Printf("The number of words are: %v\n", count)
}

func CountAllWords(content string) int {
	w := GetGridWidth(content)
	steps := []int{1, w, w + 1, w - 1} // horizontal, vertical, diagonal right, diagonal left
	count := 0
	for i := range content {
		for _, step := range steps {
			count += utils.BoolToInt(IsWord(i, step, &content))
		}
	}
	return count
}

func GetGridWidth(content string) int {
	return len(strings.Split(content, "\n")[0]) + 1
}

func IsWord(startIndex, step int, content *string) bool {
	letters := make([]string, 4)
	maxIndex := len(*content) - 1
	for i := 0; i < 4; i++ {
		index := startIndex + i*step
		if index > maxIndex {
			return false
		}
		letters[i] = string((*content)[index])
	}
	word := strings.Join(letters, "")
	return word == "XMAS" || word == "SAMX"
}
