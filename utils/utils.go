package utils

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
