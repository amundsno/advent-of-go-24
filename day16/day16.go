package day16

import (
	"advent-of-code/utils/geom"
	"advent-of-code/utils/maze"
	"fmt"

	"github.com/oleiade/lane/v2"
)

type Pose = maze.Pose
type Vec2D = geom.Vec2D

type ReindeerMaze struct {
	*maze.Maze
}

// Dijkstra (BFS with min priority queue) to explore paths from S to E that minimize the score
func (m ReindeerMaze) Explore() (seen map[Pose]int, minScore int) {
	seen = make(map[Pose]int)
	todo := lane.NewMinPriorityQueue[Pose, int]()

	start := m.StartNode()
	todo.Push(maze.NewPose(start, maze.EAST), 0)

	minScore = -1
	for todo.Size() > 0 {
		pose, score, ok := todo.Pop()
		if !ok {
			panic("could not reach the end")
		}

		symbol := m.Get(pose.POS.Y, pose.POS.X)
		if prevScore, visited := seen[pose]; (visited && prevScore <= score) ||
			symbol == maze.WALL ||
			(minScore > 0 && score > minScore) {
			continue
		}

		seen[pose] = score

		if symbol == maze.END && (minScore < 0 || score < minScore) {
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
func (m ReindeerMaze) BestTiles(seen map[Pose]int) map[Vec2D]struct{} {
	start := m.StartNode()
	tiles := make(map[Vec2D]struct{})

	var trace func(Pose, int) bool
	trace = func(pose Pose, score int) bool {
		if seenScore, exist := seen[pose]; !exist || seenScore != score {
			return false
		}
		if m.Get(pose.POS.Y, pose.POS.X) == maze.END {
			tiles[pose.POS] = struct{}{}
			return true
		}

		next := []bool{
			trace(pose.StepForward(), score+1),
			trace(pose.StepLeft(), score+1001),
			trace(pose.StepRight(), score+1001),
		}

		for _, b := range next {
			if b {
				tiles[pose.POS] = struct{}{}
				return true
			}
		}
		return false
	}
	trace(maze.NewPose(start, maze.EAST), 0)

	return tiles
}

func Solve(inputPath string) {
	m := maze.ParseFileToMaze(inputPath)
	rm := ReindeerMaze{&m}

	seen, score := rm.Explore()
	tiles := rm.BestTiles(seen)

	fmt.Printf("Part 01: %v\n", score)
	fmt.Printf("Part 02: %v\n", len(tiles))
}
