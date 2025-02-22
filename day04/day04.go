package day04

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func SolveFirst(inputPath string) {
	content := utils.ReadFileToString(inputPath)
	count := CountAllWords(content)
	fmt.Printf("The number of words are: %v\n", count)
}

func CountAllWords(content string) int {
	w := GetContentWidth(content)
	steps := []int{1, w, w + 1, w - 1} // horizontal, vertical, diagonal right, diagonal left
	count := 0
	for i := range content {
		for _, step := range steps {
			count += utils.BoolToInt(IsWord(i, step, &content))
		}
	}
	return count
}

func GetContentWidth(content string) int {
	return len(strings.Split(content, "\n")[0]) + 1
}

func IsWord(startIndex, step int, content *string) bool {
	letters := make([]string, 4)
	maxIndex := len(*content) - 1
	for i := 0; i < 4; i++ {
		index := startIndex + i*step
		if index < 0 || index > maxIndex {
			return false
		}
		letters[i] = string((*content)[index])
	}
	word := strings.Join(letters, "")
	return word == "XMAS" || word == "SAMX"
}

func SolveSecond(inputPath string) {
	content := utils.ReadFileToString(inputPath)
	count := CountAllCrosses(content)
	fmt.Printf("The number of crosses are: %v\n", count)
}

func IsCrossCentre(startIndex, contentWidth int, content *string) bool {
	letter := string((*content)[startIndex])
	if letter != "A" {
		return false
	}
	crossIndices := []int{
		startIndex - contentWidth - 1,
		startIndex - contentWidth + 1,
		startIndex + contentWidth - 1,
		startIndex + contentWidth + 1,
	}
	maxIndex := len(*content) - 1
	crossLetters := make([]string, 4)
	for i, index := range crossIndices {
		if index < 0 || index > maxIndex {
			return false
		}
		crossLetters[i] = string((*content)[index])
	}
	word := strings.Join(crossLetters, "")
	return word == "MMSS" || word == "SSMM" || word == "MSMS" || word == "SMSM"
}

func CountAllCrosses(content string) int {
	w := GetContentWidth(content)
	count := 0
	for i := range content {
		count += utils.BoolToInt(IsCrossCentre(i, w, &content))
	}
	return count
}
