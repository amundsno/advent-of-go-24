package day12

import (
	"advent-of-code/utils"
	"advent-of-code/utils/collections"
	"advent-of-code/utils/matrix"
	"fmt"
)

type Garden struct {
	*matrix.Matrix[string]
	state *ExploreState
}

type Point struct {
	i, j int
}

func (p Point) Add(q Point) Point {
	return Point{p.i + q.i, p.j + q.j}
}

type ExploreState struct {
	unexplored collections.Stack[Point]
	visited    map[Point]struct{}
}

func (g *Garden) ExploreRegion(start Point) (area, perimiter int) {
	regionState := ExploreState{
		unexplored: collections.Stack[Point]{},
		visited:    make(map[Point]struct{}),
	}

	regionState.unexplored.Push(start)

	directions := []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for regionState.unexplored.Len() > 0 {
		point := regionState.unexplored.Pop()
		if _, visited := regionState.visited[point]; visited {
			continue
		}
		area++

		regionState.visited[point] = struct{}{}
		g.state.visited[point] = struct{}{}

		region := *g.At(point.i, point.j)

		for _, dir := range directions {
			nextPoint := point.Add(dir)
			if _, visited := regionState.visited[nextPoint]; visited {
				continue
			}
			if !g.IsInbounds(nextPoint.i, nextPoint.j) {
				perimiter++
				continue
			}

			nextRegion := *g.At(nextPoint.i, nextPoint.j)
			if region != nextRegion {
				perimiter++
				g.state.unexplored.Push(nextPoint)
			} else {
				regionState.unexplored.Push(nextPoint)
			}
		}
	}

	return area, perimiter
}

func Solve(inputPath string) {
	m := matrix.New(utils.ReadFileTo2D(inputPath, ""))

	state := ExploreState{visited: make(map[Point]struct{})}
	state.unexplored.Push(Point{0, 0})

	garden := Garden{&m, &state}

	total := 0
	for state.unexplored.Len() > 0 {
		start := state.unexplored.Pop()
		if _, visited := state.visited[start]; visited {
			continue
		}

		area, perimeter := garden.ExploreRegion(start)
		total += area * perimeter

		// fmt.Println(start, *garden.At(start.i, start.j), area, perimeter)
		// fmt.Println(state.unexplored)
		// fmt.Println(state.visited)
		// fmt.Println()
	}

	fmt.Println(total)
}

// Rewrite to return area: map[Point]struct{}{}, perimiter: map[Point]int

// perimiter length is sum(perimiter.values)

// perimeter sides can be computed by counting corners.
// Take one point. Find the closest point by trying directions (N, NE, E, SE, S, SW, W, NW). The delta is d1.
// Try to apply d1 again to the second point. If it works, it is either a stair shape or straight shape.
// if d1.i == d1.j it is a stair, otherwise it is straight
// Stair = 3 corners
// Straight = 0 corners
// Otherwise it is an inside corner or outside corner, both of which has 1 corner.
