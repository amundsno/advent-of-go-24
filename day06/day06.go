package day06

import (
	"advent-of-code/utils"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
	"sync"
)

const OBSTACLE, VISITED, LEFT, UP, RIGHT, DOWN = "#", "X", "<", "^", ">", "v"

type GuardPath [][]string

type Pose struct {
	i, j int
	dir  string
}

type Position struct {
	i, j int
}

func Solve(inputPath string) {
	gp := GuardPath(utils.ReadFileToGrid(inputPath, ""))
	startPose := FindStartPose(&gp)
	prevPose := startPose

	var wg sync.WaitGroup
	var mu sync.Mutex
	obstacleCount := 0

	visitedPosition := make(map[Position]bool)

	for pose := range gp.Steps(startPose) {
		// Part 02 - Do not put an obstacle in the same position twice
		if visitedPosition[Position{pose.i, pose.j}] {
			continue
		}

		visitedPosition[Position{pose.i, pose.j}] = true

		// Part 02 - Do not put an obstacle on the start position
		if pose == startPose {
			continue
		}

		blockedPath := gp.DeepCopy()
		blockedPath[pose.i][pose.j] = OBSTACLE

		wg.Add(1)
		go func(from Pose) {
			defer wg.Done()
			if IsLoop(&blockedPath, from) {
				mu.Lock()
				obstacleCount++
				mu.Unlock()
			}
		}(prevPose)

		prevPose = pose
	}
	count := utils.IterLength(maps.Keys(visitedPosition))
	fmt.Printf("Part 01: %v (unique positions visited)\n", count)

	wg.Wait()
	fmt.Printf("Part 02: %v (number of obstructions that create a loop)\n", obstacleCount)
}

func IsLoop(gp *GuardPath, startPose Pose) bool {
	visitedPose := make(map[Pose]bool)
	for pose := range gp.Steps(startPose) {
		if visitedPose[pose] {
			return true
		}
		visitedPose[pose] = true
	}
	return false
}

func (gp *GuardPath) DeepCopy() GuardPath {
	cpy := make(GuardPath, len(*gp))
	for i, row := range *gp {
		cpy[i] = make([]string, len(row))
		copy(cpy[i], (*gp)[i])
	}
	return cpy
}

func (gp *GuardPath) Steps(startPose Pose) iter.Seq[Pose] {
	iMax := len(*gp) - 1
	jMax := len((*gp)[0]) - 1

	pose := startPose

	return func(yield func(Pose) bool) {
		di, dj := 0, 0
		for {
			if !yield(pose) {
				return
			}
			for {
				di, dj = ComputeNextStep(pose.dir)
				if IsOutOfBounds(pose.i+di, pose.j+dj, iMax, jMax) {
					return
				}
				if (*gp)[pose.i+di][pose.j+dj] != OBSTACLE {
					break
				}
				pose.dir = GetNextDirection(pose.dir)
			}
			pose.i += di
			pose.j += dj
		}
	}
}

func (gp GuardPath) String() string {
	s := ""
	for _, row := range gp {
		s += strings.Join(row, "")
		s += "\n"
	}
	return s
}

func FindStartPose(gp *GuardPath) Pose {
	directions := []string{UP, RIGHT, LEFT, DOWN}
	for i, row := range *gp {
		for _, dir := range directions {
			if j := slices.Index(row, dir); j > 0 {
				return Pose{i, j, (*gp)[i][j]}
			}
		}
	}
	panic("failed to find start pose")
}

func IsOutOfBounds(i, j, iMax, jMax int) bool {
	return i < 0 || i > iMax || j < 0 || j > jMax
}

func GetNextDirection(dir string) string {
	switch dir {
	case LEFT:
		return UP
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	default:
		return LEFT
	}
}

func ComputeNextStep(dir string) (di int, dj int) {
	switch dir {
	case LEFT:
		return 0, -1
	case UP:
		return -1, 0
	case RIGHT:
		return 0, 1
	default:
		return 1, 0
	}
}
