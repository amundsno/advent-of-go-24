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
			return -1
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
	return -1
}

func ParseInput(filepath string) []Vec2D {
	rowsStr := utils.ReadFileTo2D(filepath, ",")
	rows, _ := utils.SliceAtoi2D(rowsStr)

	noGoPos := make([]Vec2D, len(rows))
	for i, row := range rows {
		noGoPos[i] = Vec2D{row[0], row[1]}
	}
	return noGoPos
}

func NewSpace(noGo []Vec2D) Space {
	space := make(Space)
	for _, pos := range noGo {
		space[pos] = struct{}{}
	}
	return space
}

// Binary search to find the first position that would block
// a path to the exit
func FirstBlockPosition(noGoPos []Vec2D) Vec2D {
	lower, upper := BASE, len(noGoPos)-1

	for lower < upper-1 {
		i := lower + (upper-lower)/2
		space := NewSpace(noGoPos[:i])

		if space.MinSteps() < 0 {
			upper = i
		} else {
			lower = i
		}
	}
	return noGoPos[lower]
}

const MAXPOS, BASE int = 70, 1024

func Solve(filepath string) {
	noGoPos := ParseInput(filepath)
	space := NewSpace(noGoPos[:BASE])

	fmt.Printf("Part 01: %v\n", space.MinSteps())
	fmt.Printf("Part 02: %v\n", FirstBlockPosition(noGoPos))
}
