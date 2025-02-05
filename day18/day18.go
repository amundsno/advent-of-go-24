package day18

import (
	"advent-of-code/utils"
	"fmt"

	"github.com/oleiade/lane/v2"
)

type Vec2D struct {
	x, y int
}

func (v Vec2D) Add(w Vec2D) Vec2D {
	return Vec2D{v.x + w.x, v.y + w.y}
}

var UP, DOWN, LEFT, RIGHT = Vec2D{0, -1}, Vec2D{0, 1}, Vec2D{-1, 0}, Vec2D{1, 0}
var DIRECTIONS = []Vec2D{DOWN, RIGHT, LEFT, UP}

type Space map[Vec2D]struct{}

func IsInbounds(pos Vec2D) bool {
	return pos.x >= 0 && pos.x <= MAXPOS && pos.y >= 0 && pos.y <= MAXPOS
}

// Dijkstra (BFS with min-priority queue)
func (s Space) MinSteps() int {
	seen := make(map[Vec2D]struct{})

	todo := lane.NewMinPriorityQueue[Vec2D, int]()
	todo.Push(Vec2D{0, 0}, 0)

	for todo.Size() > 0 {
		pos, dist, ok := todo.Pop()
		if !ok {
			panic("could not reach the end")
		}

		if pos.x == MAXPOS && pos.y == MAXPOS {
			return dist
		}

		_, visited := seen[pos]
		seen[pos] = struct{}{}

		_, noGo := s[pos]
		if visited || !IsInbounds(pos) || noGo {
			continue
		}

		for _, dir := range DIRECTIONS {
			todo.Push(pos.Add(dir), dist+1)
		}
	}
	panic("could not reach the end")
}

func ParseInput(filepath string, nRows int) Space {
	space := make(Space)
	rowsStr := utils.ReadFileTo2D(filepath, ",")
	rows, _ := utils.SliceAtoi2D(rowsStr)
	for i, row := range rows {
		space[Vec2D{row[0], row[1]}] = struct{}{}
		if i > 0 && i+1 == nRows {
			break
		}
	}
	return space
}

const MAXPOS int = 70

func Solve(filepath string) {
	space := ParseInput(filepath, 1024)

	fmt.Printf("Part 01: %v\n", space.MinSteps())
}
