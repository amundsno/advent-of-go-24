package day06

import (
	"advent-of-code/utils"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
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

	// Part 01
	uniquePositionsMap := make(map[Position]bool)
	for pose := range gp.Steps(startPose) {
		uniquePositionsMap[Position{pose.i, pose.j}] = true
	}
	count := utils.IterLength(maps.Keys(uniquePositionsMap))
	fmt.Printf("Part 01: %v (unique positions visited)\n", count)
}

func (gp *GuardPath) Steps(startPose Pose) iter.Seq[Pose] {
	iMax := len(*gp) - 1
	jMax := len((*gp)[0]) - 1
	di, dj := 0, 0
	pose := startPose

	return func(yield func(Pose) bool) {
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
