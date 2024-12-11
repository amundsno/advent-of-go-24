package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadColumnsToSlices(reader io.Reader, delimiter string, columnCount int) ([][]int, error) {
	columns := make([][]int, columnCount)
	scanner := bufio.NewScanner(reader)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		tokens := strings.Split(line, delimiter)
		if len(tokens) != columnCount {
			return nil, fmt.Errorf("line %v: column count mismatch; expected %v, got %v (%v", lineNumber, columnCount, len(tokens), line)
		}

		for i, token := range tokens {
			value, err := strconv.Atoi(token)
			if err != nil {
				return nil, fmt.Errorf("line: %v: failed to convert %v to int: %v", lineNumber, token, err)
			}
			columns[i] = append(columns[i], value)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}

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
