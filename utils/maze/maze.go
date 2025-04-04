package maze

import (
	"advent-of-code/utils"
	"advent-of-code/utils/geom"
	"advent-of-code/utils/matrix"
	"fmt"
	"iter"
)

type Maze struct {
	*matrix.Matrix[string]
}

func NewMaze(m matrix.Matrix[string]) Maze {
	return Maze{&m}
}

func ParseFileToMaze(filepath string) Maze {
	slice2D := utils.ReadFileTo2D(filepath, "")
	m := matrix.New(slice2D)
	return NewMaze(m)
}

// Scan the maze for the first occurrence of 'symbol' (e.g. START or END)
func (m Maze) First(symbol string) (geom.Vec2D, error) {
	for y := 0; y < m.Rows(); y++ {
		for x := 0; x < m.Cols(); x++ {
			if m.Get(y, x) == symbol {
				return geom.NewVec2D(x, y), nil
			}
		}
	}
	return geom.Vec2D{}, fmt.Errorf("could not find symbol '%v'", symbol)
}

func (m Maze) StartNode() geom.Vec2D {
	node, _ := m.First(START)
	return node
}

func (m Maze) EndNode() geom.Vec2D {
	node, _ := m.First(END)
	return node
}

// Iterate over all nodes in range (i.e. in step radius)
// (scanning from top left to bottom right)
func (m Maze) NodesInRange(from geom.Vec2D, steps int) iter.Seq[geom.Vec2D] {
	return func(yield func(geom.Vec2D) bool) {
		for y := max(0, from.Y-steps); y <= min(m.Rows()-1, from.Y+steps); y++ {
			dy := from.Y - y
			if dy < 0 {
				dy = -dy
			}
			dx := steps - dy
			for x := max(0, from.X-dx); x <= min(m.Cols()-1, from.X+dx); x++ {
				if !yield(geom.NewVec2D(x, y)) {
					return
				}
			}
		}
	}
}

// The number of steps between two points as if no walls
// Manhattan distance
func (m Maze) DirectSteps(from, to geom.Vec2D) int {
	dx := max(from.X, to.X) - min(from.X, to.X)
	dy := max(from.Y, to.Y) - min(from.Y, to.Y)
	return dy + dx
}

type Pose struct {
	POS, DIR geom.Vec2D
}

func NewPose(position, direction geom.Vec2D) Pose {
	return Pose{POS: position, DIR: direction}
}

func (p Pose) StepForward() Pose {
	return Pose{p.POS.Add(p.DIR), p.DIR}
}

func (p Pose) StepRight() Pose {
	return Pose{p.POS.Add(p.DIR.TurnRight()), p.DIR.TurnRight()}
}

func (p Pose) StepLeft() Pose {
	return Pose{p.POS.Add(p.DIR.TurnLeft()), p.DIR.TurnLeft()}
}

var NORTH, SOUTH, EAST, WEST = geom.NewVec2D(0, -1),
	geom.NewVec2D(0, 1),
	geom.NewVec2D(1, 0),
	geom.NewVec2D(-1, 0)

var DIRECTIONS []geom.Vec2D = []geom.Vec2D{NORTH, SOUTH, EAST, WEST}

const START, END, WALL, SPACE string = "S", "E", "#", "."
