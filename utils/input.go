package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadTo1D(reader io.Reader) ([]string, error) {
	var rows []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return rows, nil
}

func ReadFileTo1D(inputPath string) []string {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ReadTo1D(file)
	if err != nil {
		panic(err)
	}
	return content
}

func ReadTo2D(reader io.Reader, delimiter string) ([][]string, error) {
	grid := [][]string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), delimiter)
		grid = append(grid, tokens)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func ReadFileTo2D(inputPath, delimiter string) [][]string {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ReadTo2D(file, delimiter)
	if err != nil {
		panic(err)
	}
	return content
}

// TODO: Refactor I/O. Create a utils for converting to int, instead of doing it in the reader.
// Create a general readFile. Check stash.
func ReadRowsToSlices(reader io.Reader, delimiter string) ([][]int, error) {
	rows := [][]int{}
	scanner := bufio.NewScanner(reader)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		row := []int{}
		tokens := strings.Split(line, delimiter)
		for _, token := range tokens {
			value, err := strconv.Atoi(token)
			if err != nil {
				return nil, fmt.Errorf("line: %v: failed to convert %v to int: %v", lineNumber, token, err)
			}
			row = append(row, value)
		}

		rows = append(rows, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return rows, nil
}

func readToString(reader io.Reader) (string, error) {
	content, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ReadFileToString(inputPath string) string {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := readToString(file)
	if err != nil {
		panic(err)
	}
	return content
}

func ReadToMap(reader io.Reader, delimiter string) (map[int]map[int]string, error) {
	m := make(map[int]map[int]string)
	scanner := bufio.NewScanner(reader)
	iRow := -1
	for scanner.Scan() {
		iRow++
		text := scanner.Text()

		row := strings.Split(text, delimiter)
		m[iRow] = make(map[int]string)

		for iCol, val := range row {
			m[iRow][iCol] = val
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return m, nil
}

func ReadFileToMap(inputPath, delimiter string) map[int]map[int]string {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ReadToMap(file, delimiter)
	if err != nil {
		panic(err)
	}
	return content
}
