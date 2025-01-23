package day16

import (
	"advent-of-code/utils"
	"advent-of-code/utils/collections"
	"advent-of-code/utils/matrix"
	"fmt"
)

type Grid struct {
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

const WALL, SPACE, START, END string = "#", ".", "S", "E"

var EAST, SOUTH, WEST, NORTH Vec2D = Vec2D{1, 0}, Vec2D{0, 1}, Vec2D{-1, 0}, Vec2D{0, -1}

func parseInput(inputPath string) Grid {
	slice := utils.ReadFileTo2D(inputPath, "")
	m := matrix.New(slice)
	return Grid{&m}
}

func (g Grid) StartPosition() Vec2D {
	for y := range g.Rows() {
		for x := range g.Cols() {
			if g.Get(y, x) == START {
				return Vec2D{x, y}
			}
		}
	}
	panic("could not find start position")
}

type Step struct {
	pos, dir Vec2D
	score    int
}

// BFS Flood Fill
// By queuing the turn itself as a step, we ensure that the first to reach the target
// is in the least amount of steps and turns. NB! There might be more than one path that is best.
// Might be possible if the memo included direction?
// Just memoize the whole step? No... Minus the score
// New struct pose{pos, dir}, score int
// memo[pos, dir.flip()] - i.e. if visited before in the other direction
// But then we need another memo for the points?

// Or just count the number of tiles traversed as well?

// Or extend the memo to include the whole step object
// Then create a function to reverse the memo from the END
// Try the different directions, if it exist in the memo, the step is valid
// DFS apporach for the reverse, use another memo to see if a position has been counted already
func (g Grid) MinScore(start Vec2D) int {
	memo := make(map[Vec2D]int)

	todo := collections.Queue[Step]{}
	todo.Enqueue(Step{start, EAST, 0})

	minScore := -1
	for todo.Len() > 0 {
		step := todo.Dequeue()
		pos, dir, score := step.pos, step.dir, step.score
		symbol := g.Get(pos.y, pos.x)

		if prevScore, visited := memo[pos]; (visited && prevScore < score) ||
			symbol == WALL ||
			(minScore > 0 && score >= minScore) {
			continue
		}

		if symbol == END {
			if minScore < 0 || score < minScore {
				minScore = score
			}
			continue
		}

		memo[pos] = score

		todo.Enqueue(Step{pos.Add(dir), dir, score + 1})
		todo.Enqueue(Step{pos.Add(dir.TurnRight()), dir.TurnRight(), score + 1001})
		todo.Enqueue(Step{pos.Add(dir.TurnLeft()), dir.TurnLeft(), score + 1001})
	}
	if minScore < 0 {
		panic("could not reach the end")
	}
	return minScore
}

func Solve(inputPath string) {
	grid := parseInput(inputPath)
	start := grid.StartPosition()

	minPoints := grid.MinScore(start)
	fmt.Printf("Part 01: %v\n", minPoints)

}
