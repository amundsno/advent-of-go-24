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

const WALL, SPACE, BOT, BOX, BOX_LEFT, BOX_RIGHT string = "#", ".", "@", "O", "[", "]"

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

func (g *Grid) GetMoveHandler(pos, dir Vec2D, memo map[Vec2D]struct{}) func() {
	if _, visited := memo[pos]; visited {
		return func() {}
	}
	memo[pos] = struct{}{}

	sym := g.Get(pos.y, pos.x)
	if sym == SPACE {
		return func() {}
	}
	if sym == WALL {
		return nil
	}

	posNext := pos.Add(dir)
	moveFunc := g.GetMoveHandler(posNext, dir, memo)

	if moveFunc == nil {
		return nil
	}

	if sym != BOX_LEFT && sym != BOX_RIGHT {
		return func() {
			moveFunc()
			g.Set(posNext.y, posNext.x, sym)
			g.Set(pos.y, pos.x, SPACE)
		}
	}

	posLinked := pos.Add(RIGHT)
	if sym == BOX_RIGHT {
		posLinked = pos.Add(LEFT)
	}

	moveFuncLinked := g.GetMoveHandler(posLinked, dir, memo)
	if moveFuncLinked == nil {
		return nil
	}

	return func() {
		moveFunc()
		moveFuncLinked()
		g.Set(posNext.y, posNext.x, sym)
		g.Set(pos.y, pos.x, SPACE)
	}
}

func DoubleGrid(g Grid) Grid {
	newWidth := g.Cols() * 2
	newGrid := make([][]string, g.Rows())
	for i := range g.Rows() {
		newRow := make([]string, newWidth)
		for j := range g.Cols() {
			var newTiles []string
			switch g.Get(i, j) {
			case WALL:
				newTiles = []string{WALL, WALL}
			case BOX:
				newTiles = []string{BOX_LEFT, BOX_RIGHT}
			case SPACE:
				newTiles = []string{SPACE, SPACE}
			case BOT:
				newTiles = []string{BOT, SPACE}
			}
			newRow[2*j] = newTiles[0]
			newRow[2*j+1] = newTiles[1]
		}
		newGrid[i] = newRow
	}
	m := matrix.New(newGrid)
	return Grid{&m}
}

func (g *Grid) StartPosition() Vec2D {
	for y := range g.Rows() {
		for x := range g.Cols() {
			if g.Get(y, x) == BOT {
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
			if g.Get(y, x) == BOX || g.Get(y, x) == BOX_LEFT {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func SolvePart01(inputPath string) {
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

func (g *Grid) DoMove(pos, dir Vec2D) {
	next := pos.Add(dir)
	a := g.Get(pos.y, pos.x)
	b := g.Get(next.y, next.x)
	g.Set(pos.y, pos.x, b)
	g.Set(next.y, next.x, a)
}

func SolvePart02(inputPath string) {
	grid, moves := parseFileInput(inputPath)
	grid = DoubleGrid(grid)
	pos := grid.StartPosition()

	// r := bufio.NewReader(os.Stdin)

	for _, move := range moves {
		dir := ParseDirection(move)

		// Debugging
		// fmt.Println(grid)
		// fmt.Println(dir)
		// r.ReadString('\n')

		moveHandler := grid.GetMoveHandler(pos, dir, make(map[Vec2D]struct{}))
		if moveHandler != nil {
			moveHandler()
			pos = pos.Add(dir)
		}
	}
	fmt.Printf("Part 02: %v\n", grid.SumGpsCoordinates())
	// fmt.Println(grid)
}
