package day03

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	sum, err := SumValidMultiplications(content)
	if err != nil {
		log.Fatalf("unexpected error ocurred: %v", err)
	}
	fmt.Printf("The sum of valid multiplications is: %v\n", sum)
}

func SolveSecond(inputPath string) {
	content := readInput(inputPath)
	fmt.Println(content)

	enabledRegions := ExtractEnabledRegions(content)
	sum := 0
	for _, region := range enabledRegions {
		regionSum, err := SumValidMultiplications(region)
		if err != nil {
			log.Fatalf("unexpected error ocurred: %v", err)
		}
		sum += regionSum
	}
	fmt.Printf("The sum of valid multiplications in enabled regions is: %v\n", sum)
}

func SumValidMultiplications(input string) (int, error) {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := pattern.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		left, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, fmt.Errorf("failed to convert %v to int", match[1])
		}
		right, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, fmt.Errorf("failed to convert %v to int", match[2])
		}
		sum += left * right
	}

	return sum, nil
}

func ExtractEnabledRegions(input string) []string {
	pattern := regexp.MustCompile(`(?s)(?:^.*?don't\(\))|(?:do\(\).*?don't\(\))|(?:do\(\).*?$)`)
	return pattern.FindAllString(input, -1)
}
