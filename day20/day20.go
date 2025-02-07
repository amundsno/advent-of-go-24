package day20

import (
	"advent-of-code/utils/geom"
	"advent-of-code/utils/maze"
	"fmt"
)

type RaceMaze struct {
	*maze.Maze
}

type Vec2D = geom.Vec2D

// Number of steps remaining to reach the maze END from each node
func (rm RaceMaze) TimeToEndByNode() map[Vec2D]int {
	timeByNode := make(map[Vec2D]int)
	end := rm.EndNode()

	// DFS traversal of the maze from END -> START
	var traverse func(maze.Pose, int)
	traverse = func(pose maze.Pose, step int) {
		if !rm.IsInbounds(pose.POS.Y, pose.POS.X) {
			return
		}
		symbol := rm.Get(pose.POS.Y, pose.POS.X)
		if symbol == maze.WALL {
			return
		}
		timeByNode[pose.POS] = step

		if symbol == maze.START {
			return
		}
		next := []maze.Pose{
			pose.StepForward(),
			pose.StepRight(),
			pose.StepLeft(),
		}
		for _, pose := range next {
			traverse(pose, step+1)
		}
	}

	// Find start direction
	for _, dir := range maze.DIRECTIONS {
		next := end.Add(dir)
		if rm.Get(next.Y, next.X) == maze.SPACE {
			startPose := maze.NewPose(end, dir)
			traverse(startPose, 0)
			break
		}
	}
	return timeByNode
}

type Cheat struct {
	start, end Vec2D
}

func (rm RaceMaze) CheatsByTimeSaved(timeByNode map[Vec2D]int) map[int][]Cheat {
	cheatsByTime := make(map[int][]Cheat)

	for node, time := range timeByNode {
		for _, dir := range maze.DIRECTIONS {
			cheatStart := node.Add(dir)

			// It is only a cheat if we pass through a wall
			if rm.Get(cheatStart.Y, cheatStart.X) != maze.WALL {
				continue
			}

			cheatEnd := cheatStart.Add(dir)
			cheatTime, canCheat := timeByNode[cheatEnd]

			// The cheat is valid only if we end up on the maze path
			if !canCheat {
				continue
			}

			delta := time - cheatTime - 2

			// Only cheats that yield a better time are of interest
			if delta <= 0 {
				continue
			}

			cheatsByTime[delta] = append(cheatsByTime[delta], Cheat{cheatStart, cheatEnd})
		}
	}
	return cheatsByTime
}

func Solve(filepath string) {
	m := maze.ParseFileToMaze(filepath)
	rm := RaceMaze{&m}

	timeByNode := rm.TimeToEndByNode()
	cheatsByTime := rm.CheatsByTimeSaved(timeByNode)

	count := 0
	for time, cheats := range cheatsByTime {
		if time >= 100 {
			count += len(cheats)
		}
	}

	fmt.Printf("Part 01: %v\n", count)
}
