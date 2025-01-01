package day10

import (
	"advent-of-code/utils"
	"fmt"
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

func (p *Point) Add(q Point) Point {
	return Point{p.i + q.i, p.j + q.j}
}

func Solve(inputPath string) {
	gridStr := utils.ReadFileTo2D(inputPath, "")
	grid, err := utils.SliceAtoi2D(gridStr)
	if err != nil {
		panic(err)
	}

	m := NewMatrix(grid)
	directions := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	scoreSum := 0
	ratingSum := 0
	for i := range m.n {
		for j := range m.m {
			if !m.isTrailhead(i, j) {
				continue
			}

			trailheadState := TrailState{
				reachableHilltops: map[Point]struct{}{},
				trailheadRating:   0,
			}

			m.ExploreTrail(Point{i, j}, directions, &trailheadState)

			scoreSum += len(trailheadState.reachableHilltops)
			ratingSum += trailheadState.trailheadRating

		}
	}
	fmt.Printf("Part 01: %v (sum of all trailhead scores)\n", scoreSum)
	fmt.Printf("Part 02: %v (sum of all trailhead ratings)\n", ratingSum)
}

func (m *Matrix) isTrailhead(i, j int) bool {
	return *m.At(i, j) == 0
}

func (m *Matrix) isValidStep(start, step Point) bool {
	next := start.Add(step)
	return m.IsInbounds(next.i, next.j) &&
		*m.At(next.i, next.j) == *m.At(start.i, start.j)+1
}

type TrailState struct {
	reachableHilltops map[Point]struct{}
	trailheadRating   int
}

// Recursive DFS in grid
func (m *Matrix) ExploreTrail(from Point, directions []Point, state *TrailState) {
	if *m.At(from.i, from.j) == 9 {
		state.reachableHilltops[from] = struct{}{}
		state.trailheadRating++
		return
	}

	for _, dir := range directions {
		if !m.isValidStep(from, dir) {
			continue
		}

		next := from.Add(dir)
		m.ExploreTrail(next, directions, state)
	}
}
