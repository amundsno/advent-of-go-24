package day15

import (
	"advent-of-code/utils"
	"advent-of-code/utils/matrix"
	"fmt"
	"strings"
)

type Grid struct {
	*matrix.Matrix[string]
}

type Vec2D struct {
	x, y int
}

func (v *Vec2D) Add(w Vec2D) Vec2D {
	return Vec2D{v.x + w.x, v.y + w.y}
}

const WALL, SPACE, BOT, BOX string = "#", ".", "@", "O"

var RIGHT, DOWN, LEFT, UP Vec2D = Vec2D{1, 0}, Vec2D{0, 1}, Vec2D{-1, 0}, Vec2D{0, -1}

func ParseDirection(symbol rune) Vec2D {
	switch symbol {
	case '<':
		return LEFT
	case '^':
		return UP
	case '>':
		return RIGHT
	case 'v':
		return DOWN
	}
	panic("could not parse direction")
}

func (g *Grid) Move(pos, dir Vec2D) bool {
	sym := g.Get(pos.y, pos.x)
	if sym == SPACE {
		return true
	}
	if sym == WALL {
		return false
	}

	next := pos.Add(dir)
	if g.Move(next, dir) {
		g.Set(next.y, next.x, sym)
		return true
	}
	return false
}

func (g *Grid) StartPosition() Vec2D {
	for y := range g.Rows() {
		for x := range g.Cols() {
			if g.Get(x, y) == BOT {
				return Vec2D{x, y}
			}
		}
	}
	panic("could not find start position")
}

func (g Grid) String() string {
	s := ""
	for y := range g.Rows() {
		s += strings.Join(g.Row(y), "") + "\n"
	}
	return s
}

func parseFileInput(path string) (grid Grid, moves string) {
	content := utils.ReadFileToString(path)
	parts := strings.Split(content, "\n\n")

	g, _ := utils.ReadTo2D(strings.NewReader(parts[0]), "")
	m := matrix.New(g)

	return Grid{&m}, strings.ReplaceAll(parts[1], "\n", "")
}

func (g *Grid) SumGpsCoordinates() int {
	sum := 0
	for y := range g.Rows() {
		for x := range g.Cols() {
			if g.Get(y, x) == BOX {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func Solve(inputPath string) {
	grid, moves := parseFileInput(inputPath)
	pos := grid.StartPosition()

	// r := bufio.NewReader(os.Stdin)

	for _, move := range moves {
		dir := ParseDirection(move)

		// Debugging
		// fmt.Println(grid)
		// fmt.Println(dir)
		// r.ReadString('\n')

		if grid.Move(pos, dir) {
			grid.Set(pos.y, pos.x, SPACE)
			pos = pos.Add(dir)
		}
	}

	fmt.Printf("Part 01: %v\n", grid.SumGpsCoordinates())
	// fmt.Println(grid)
}

// To solve the wide case, it will only be relevant to check neighbours when moving up/down
// Simply split the recursion and require that both is True before moving
