package day12

import (
	"advent-of-code/utils"
	"advent-of-code/utils/collections"
	"advent-of-code/utils/matrix"
	"fmt"
	"maps"
)

var NORTH, EAST, SOUTH, WEST = Point{-1, 0}, Point{0, 1}, Point{1, 0}, Point{0, -1}
var NESW = []Point{NORTH, EAST, SOUTH, WEST}

type Garden struct {
	*matrix.Matrix[string]
}

type Region map[Point]struct{}

func (r *Region) Area() int {
	return len(*r)
}

func (r *Region) Perimeter() int {
	perim := 0
	for p := range *r {
		for _, dir := range NESW {
			if _, exist := (*r)[p.Add(dir)]; !exist {
				perim++
			}
		}
	}
	return perim
}

func (r *Region) Sides() int {
	sides := 0
	for p := range maps.Keys(*r) {
		sides += r.cornerCount(p)
	}
	return sides
}

func (r *Region) cornerCount(p Point) int {
	count := 0
	for i := 0; i < len(NESW); i++ {
		normal := NESW[i]
		orthog := NESW[(i+1)%len(NESW)]
		diagon := normal.Add(orthog)

		_, normalExist := (*r)[p.Add(normal)]
		_, orthogExist := (*r)[p.Add(orthog)]
		_, diagonExist := (*r)[p.Add(diagon)]

		// Outer corners - e.g. if the point to the North and East are not in the region
		// 00
		// 10
		if !normalExist && !orthogExist {
			count++
		}

		// Inner corners - e.g. if the point ot the North and East are in the region, but the NE point is not
		// 10
		// 11
		if normalExist && orthogExist && !diagonExist {
			count++
		}

	}
	return count
}

type Point struct {
	i, j int
}

func (p Point) Add(q Point) Point {
	return Point{p.i + q.i, p.j + q.j}
}

func (g *Garden) ExploreRegion(start Point) Region {
	region := make(map[Point]struct{})

	unexplored := collections.Queue[Point]{}
	unexplored.Enqueue(start)

	for unexplored.Len() > 0 {
		point := unexplored.Dequeue()
		if _, visited := region[point]; visited {
			continue
		}
		region[point] = struct{}{}

		label := g.Get(point.i, point.j)
		for _, dir := range NESW {
			nextPoint := point.Add(dir)
			if !g.IsInbounds(nextPoint.i, nextPoint.j) {
				continue
			}

			nextLabel := g.Get(nextPoint.i, nextPoint.j)
			if label == nextLabel {
				unexplored.Enqueue(nextPoint)
			}
		}
	}

	return region
}

func (g *Garden) Explore() []Region {
	visited := map[Point]struct{}{}
	regions := []Region{}

	for i := range g.Rows() {
		for j := range g.Cols() {
			start := Point{i, j}
			if _, seen := visited[start]; seen {
				continue
			}
			region := g.ExploreRegion(start)
			regions = append(regions, region)
			for p := range region {
				visited[p] = struct{}{}
			}
		}
	}

	return regions
}

func Solve(inputPath string) {
	m := matrix.New(utils.ReadFileTo2D(inputPath, ""))
	garden := Garden{&m}
	regions := garden.Explore()

	sumPart01, sumPart02 := 0, 0
	for _, region := range regions {
		area := region.Area()
		perimeter := region.Perimeter()
		sumPart01 += area * perimeter

		sides := region.Sides()
		sumPart02 += area * sides
	}

	fmt.Printf("Part 01: %v (sum perimeter*area for each region)\n", sumPart01)
	fmt.Printf("Part 02: %v (sum sides*area for each region)\n", sumPart02)
}
