package utils

import (
	"bufio"
	"fmt"
	"io"
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

		fields := strings.Split(line, delimiter)
		if len(fields) != columnCount {
			return nil, fmt.Errorf("line %v: column count mismatch; expected %v, got %v (%v", lineNumber, columnCount, len(fields), line)
		}

		for i, field := range fields {
			value, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("line: %v: failed to convert %v to int: %v", lineNumber, value, err)
			}
			columns[i] = append(columns[i], value)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}
