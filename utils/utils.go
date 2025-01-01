package utils

import (
	"fmt"
	"iter"
	"strconv"
)

func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Performance optimized conversion algorithm (https://dev.to/chigbeef_77/bool-int-but-stupid-in-go-3jb3)
func BoolToInt(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func Transpose[T any](s [][]T) [][]T {
	transposed := make([][]T, len(s[0]))

	for _, row := range s {
		for j, val := range row {
			transposed[j] = append(transposed[j], val)
		}
	}

	return transposed
}

func IterLength[V any](s iter.Seq[V]) int {
	count := 0
	for range s {
		count++
	}
	return count
}

// (0, 0) > (i, j) > (iMax, jMax)
func IsOutOfBounds2D(i, j, iMax, jMax int) bool {
	return i < 0 || i > iMax || j < 0 || j > jMax
}

func SliceAtoi(s []string) ([]int, error) {
	is := make([]int, len(s))
	var err error
	for i, a := range s {
		is[i], err = strconv.Atoi(a)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%v' to int: %v", a, err)
		}
	}
	return is, nil
}

func SliceAtoi2D(g [][]string) ([][]int, error) {
	gInt := make([][]int, len(g))
	for i, row := range g {
		rowInt, err := SliceAtoi(row)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%v' to ints: %v", row, err)
		}
		gInt[i] = rowInt
	}
	return gInt, nil
}

func SliceSum(s []int) (sum int) {
	for _, val := range s {
		sum += val
	}
	return sum
}
