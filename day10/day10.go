package day10

import (
	"advent-of-code/utils"
	"fmt"
	"maps"
)

type Matrix struct {
	matrix [][]int
	n, m   int
}

func NewMatrix(m [][]int) Matrix {
	return Matrix{m, len(m), len(m[0])}
}

func (m *Matrix) IsInbounds(i, j int) bool {
	return !utils.IsOutOfBounds2D(i, j, m.n-1, m.m-1)
}

func (m *Matrix) At(i, j int) *int {
	return &m.matrix[i][j]
}

type Point struct {
	i, j int
}

func Solve(inputPath string) {
	gridStr := utils.ReadFileTo2D(inputPath, "")
	grid, err := utils.SliceAtoi2D(gridStr)
	if err != nil {
		panic(err)
	}

	m := NewMatrix(grid)

	directions := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	sum := 0
	for i := range m.n {
		for j := range m.m {
			ref := m.At(i, j)
			if *ref != 0 {
				continue
			}

			reachable := map[Point]struct{}{}
			Dfs(&reachable, Point{i, j}, &m, directions)
			sum += utils.IterLength(maps.Keys(reachable))
		}
	}
	fmt.Printf("Part 01: %v (sum of all trailhead scores)\n", sum)
}

func (m *Matrix) isValidStep(start, step Point) bool {
	next := Point{start.i + step.i, start.j + step.j}
	return m.IsInbounds(next.i, next.j) &&
		*m.At(next.i, next.j) == *m.At(start.i, start.j)+1
}

func Dfs(reachable *map[Point]struct{}, from Point, m *Matrix, directions []Point) {
	if *m.At(from.i, from.j) == 9 {
		(*reachable)[from] = struct{}{}
		return
	}

	for _, dir := range directions {
		if !m.isValidStep(from, dir) {
			continue
		}

		next := Point{from.i + dir.i, from.j + dir.j}
		Dfs(reachable, next, m, directions)
	}
}
