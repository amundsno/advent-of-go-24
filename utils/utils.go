package utils

import "iter"

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
