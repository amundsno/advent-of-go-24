package day06

import (
	"advent-of-code/utils"
	"fmt"
	"iter"
	"maps"
	"slices"
)

const OBSTACLE, VISITED, LEFT, UP, RIGHT, DOWN = "#", "X", "<", "^", ">", "v"

func Solve(inputPath string) {
	guardMap := utils.ReadFileToMap(inputPath, "")
	Traverse(&guardMap)
	count := CountUniquePositionsTraversed(&guardMap)
	fmt.Printf("Part 01: %v (unique positions visited)\n", count)
}

func PrintGuardMap(m *map[int]map[int]string) {
	for i := range slices.Sorted(maps.Keys(*m)) {
		fmt.Print(i)
		for j := range slices.Sorted(maps.Keys((*m)[i])) {
			fmt.Print((*m)[i][j])
		}
		fmt.Println()
	}
}

func Traverse(m *map[int]map[int]string) {
	i, j := FindStartPos(m)
	iMax, jMax := GetIndexLimits(m)
	di, dj := 0, 0
	dir := (*m)[i][j]

	outOfBounds := false
	for !outOfBounds {
		(*m)[i][j] = dir
		for {
			di, dj = ComputeNextStep(dir)
			if (*m)[i+di][j+dj] != OBSTACLE {
				break
			}
			dir = GetNextDirection(dir)
		}
		(*m)[i][j] = VISITED
		i += di
		j += dj
		outOfBounds = IsOutOfBounds(i, j, iMax, jMax)
	}
}

func IsOutOfBounds(i, j, iMax, jMax int) bool {
	return i < 0 || i > iMax || j < 0 || j > jMax
}

func seqLen[V any](s iter.Seq[V]) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func GetIndexLimits(m *map[int]map[int]string) (iMax, jMax int) {
	height := seqLen(maps.Keys(*m))
	width := seqLen(maps.Keys((*m)[0]))
	iMax = height - 1
	jMax = width - 1
	return iMax, jMax
}

func GetNextDirection(dir string) string {
	switch dir {
	case LEFT:
		return UP
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	default:
		return LEFT
	}
}

func ComputeNextStep(dir string) (di int, dj int) {
	switch dir {
	case LEFT:
		return 0, -1
	case UP:
		return -1, 0
	case RIGHT:
		return 0, 1
	default:
		return 1, 0
	}
}

func FindStartPos(m *map[int]map[int]string) (i, j int) {
	for i, row := range *m {
		for j, val := range row {
			if val == RIGHT || val == LEFT || val == UP || val == DOWN {
				return i, j
			}
		}
	}
	panic("failed to find start positon")
}

func CountUniquePositionsTraversed(m *map[int]map[int]string) int {
	count := 0
	for _, row := range *m {
		for _, val := range row {
			if val == VISITED {
				count++
			}
		}
	}
	return count
}
