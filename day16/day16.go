package day16

import (
	"advent-of-code/utils"
	"advent-of-code/utils/matrix"
	"fmt"

	"github.com/oleiade/lane/v2"
)

type Maze struct {
	*matrix.Matrix[string]
}

type Vec2D struct {
	x, y int
}

func (v Vec2D) Add(w Vec2D) Vec2D {
	return Vec2D{v.x + w.x, v.y + w.y}
}

func (v Vec2D) TurnRight() Vec2D {
	return Vec2D{-v.y, v.x}
}

func (v Vec2D) TurnLeft() Vec2D {
	return Vec2D{v.y, -v.x}
}

type Pose struct {
	pos, dir Vec2D
}

func (p Pose) StepForward() Pose {
	return Pose{p.pos.Add(p.dir), p.dir}
}

func (p Pose) StepRight() Pose {
	return Pose{p.pos.Add(p.dir.TurnRight()), p.dir.TurnRight()}
}

func (p Pose) StepLeft() Pose {
	return Pose{p.pos.Add(p.dir.TurnLeft()), p.dir.TurnLeft()}
}

const WALL, START, END string = "#", "S", "E"

var EAST = Vec2D{1, 0}

func (m Maze) startPosition() Vec2D {
	// Iterate from bottom left corner
	for y := m.Rows() - 2; y >= 0; y-- {
		for x := 1; x < m.Cols(); x++ {
			if m.Get(y, x) == START {
				return Vec2D{x, y}
			}
		}
	}
	panic("could not find start position")
}

func parseInput(inputPath string) Maze {
	slice := utils.ReadFileTo2D(inputPath, "")
	m := matrix.New(slice)
	return Maze{&m}
}

// Dijkstra (BFS with min priority queue) to explore paths from S to E that minimize the score
func (m Maze) Explore() (seen map[Pose]int, minScore int) {
	seen = make(map[Pose]int)
	todo := lane.NewMinPriorityQueue[Pose, int]()

	start := m.startPosition()
	todo.Push(Pose{start, EAST}, 0)

	minScore = -1
	for todo.Size() > 0 {
		pose, score, ok := todo.Pop()
		if !ok {
			panic("could not reach the end")
		}

		symbol := m.Get(pose.pos.y, pose.pos.x)
		if prevScore, visited := seen[pose]; (visited && prevScore <= score) ||
			symbol == WALL ||
			(minScore > 0 && score > minScore) {
			continue
		}

		seen[pose] = score

		if symbol == END && (minScore < 0 || score < minScore) {
			minScore = score
			continue
		}

		todo.Push(pose.StepForward(), score+1)
		todo.Push(pose.StepRight(), score+1001)
		todo.Push(pose.StepLeft(), score+1001)
	}
	if minScore < 0 {
		panic("could not reach the end")
	}
	return seen, minScore
}

// Recursive DFS on the paths explored by Dijkstra to return tiles on a path which minimize the score
func (m Maze) BestTiles(seen map[Pose]int) map[Vec2D]struct{} {
	start := m.startPosition()
	tiles := make(map[Vec2D]struct{})

	var trace func(Pose, int) bool
	trace = func(pose Pose, score int) bool {
		if seenScore, exist := seen[pose]; !exist || seenScore != score {
			return false
		}
		if m.Get(pose.pos.y, pose.pos.x) == END {
			tiles[pose.pos] = struct{}{}
			return true
		}

		next := []bool{
			trace(pose.StepForward(), score+1),
			trace(pose.StepLeft(), score+1001),
			trace(pose.StepRight(), score+1001),
		}

		for _, b := range next {
			if b {
				tiles[pose.pos] = struct{}{}
				return true
			}
		}
		return false
	}
	trace(Pose{start, EAST}, 0)

	return tiles
}

func Solve(inputPath string) {
	maze := parseInput(inputPath)

	seen, score := maze.Explore()
	tiles := maze.BestTiles(seen)

	fmt.Printf("Part 01: %v\n", score)
	fmt.Printf("Part 02: %v\n", len(tiles))
}
