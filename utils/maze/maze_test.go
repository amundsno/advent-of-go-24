package maze

import (
	"advent-of-code/utils/geom"
	"advent-of-code/utils/matrix"
	"slices"
	"testing"
)

func TestNodesInRadius(t *testing.T) {
	tests := []struct {
		name         string
		n, m, radius int
		from         geom.Vec2D
		expectedLen  int
	}{
		{"7x7, (3,3), r = 3", 7, 7, 3, geom.NewVec2D(3, 3), 25},
		{"4x4, (1,1), r = 2", 4, 4, 2, geom.NewVec2D(1, 1), 11},
		{"6x6, (5,5), r = 6", 6, 6, 6, geom.NewVec2D(5, 5), 26},
		{"4x4, (2,2), r = 0", 4, 4, 0, geom.NewVec2D(2, 2), 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			grid2D := make([][]string, tc.n)
			for i := range tc.n {
				grid2D[i] = make([]string, tc.m)
			}
			mat := matrix.New(grid2D)
			maze := NewMaze(mat)
			points := slices.Collect(maze.NodesInRange(tc.from, tc.radius))
			if tc.expectedLen != len(points) {
				t.Errorf("Expected length '%v' got '%v' (%v)\n", tc.expectedLen, len(points), points)
			}
		})
	}
}
