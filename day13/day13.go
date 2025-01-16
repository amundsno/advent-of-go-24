package day13

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

const PRICE_A, PRICE_B = 3, 1

func SolvePart01(inputPath string) {
	machines := ParseInput(inputPath)

	sum := 0
	for _, m := range machines {
		ctx := RecursionContext{
			memo:   map[Point]int{},
			depthA: 0,
			depthB: 0,
		}
		cost := MinCostToPriceDynamic(m.Target, m.ButtonA, m.ButtonB, ctx)
		if cost >= 0 {
			sum += cost
		}
	}

	fmt.Printf("Part 01: %v (minimum cost to reach all prizes)\n", sum)
}

func SolvePart02(inputPath string) {
	machines := ParseInput(inputPath)

	sum := 0
	for _, m := range machines {
		m.Target.x += 10000000000000
		m.Target.y += 10000000000000
		cost := CostToPriceLinAlg(m)
		if cost >= 0 {
			sum += cost
		}
	}
	fmt.Printf("Part 02: %v (minimum cost to reach all prizes)\n", sum)
}

// Use linear algebra to find the number of times to press the A & B button, then compute the cost
func CostToPriceLinAlg(cm ClawMachine) int {
	// NB! This only works if A & B are linearly independent!
	// There might be solutions, even if A & B are linearly dependent.

	ax, ay := cm.ButtonA.x, cm.ButtonA.y
	bx, by := cm.ButtonB.x, cm.ButtonB.y
	tx, ty := cm.Target.x, cm.Target.y

	// Determinant
	det := ax*by - bx*ay

	aTimesDet := tx*by - ty*bx
	bTimesDet := ty*ax - tx*ay

	if aTimesDet%det == 0 && bTimesDet%det == 0 {
		return PRICE_A*(aTimesDet/det) + PRICE_B*(bTimesDet/det)
	}

	return -1
}

type ClawMachine struct {
	ButtonA, ButtonB, Target Point
}

func ParseInput(inputPath string) []ClawMachine {
	content := utils.ReadFileToString(inputPath)
	pattern := regexp.MustCompile(`.*X.(\d+).*Y.(\d+)`)
	matches := pattern.FindAllStringSubmatch(content, -1)

	problems := make([]ClawMachine, 0)
	for i := 0; i < len(matches); i += 3 {
		ax, _ := strconv.Atoi(matches[i][1])
		ay, _ := strconv.Atoi(matches[i][2])
		bx, _ := strconv.Atoi(matches[i+1][1])
		by, _ := strconv.Atoi(matches[i+1][2])
		tx, _ := strconv.Atoi(matches[i+2][1])
		ty, _ := strconv.Atoi(matches[i+2][2])

		problems = append(problems, ClawMachine{
			ButtonA: Point{ax, ay},
			ButtonB: Point{bx, by},
			Target:  Point{tx, ty}})
	}

	return problems
}

type Point struct {
	x, y int
}

func (p *Point) Minus(q Point) Point {
	return Point{p.x - q.x, p.y - q.y}
}

type RecursionContext struct {
	memo           map[Point]int
	depthA, depthB int
}

// Use dynamic programming to find the cheapest way to reach the target
func MinCostToPriceDynamic(target, a, b Point, ctx RecursionContext) int {
	if ctx.depthA > 100 || ctx.depthB > 100 {
		return -1
	}

	if cost, exist := ctx.memo[target]; exist {
		return cost
	}

	if target.x < 0 || target.y < 0 {
		return -1
	}
	if target.x == 0 && target.y == 0 {
		return 0
	}

	costPathA := MinCostToPriceDynamic(target.Minus(a), a, b, RecursionContext{ctx.memo, ctx.depthA + 1, ctx.depthB})
	costPathB := MinCostToPriceDynamic(target.Minus(b), a, b, RecursionContext{ctx.memo, ctx.depthA, ctx.depthB + 1})

	var cost int
	if costPathA < 0 && costPathB < 0 {
		cost = -1
	} else if costPathA >= 0 && costPathB >= 0 {
		cost = min(costPathA+PRICE_A, costPathB+PRICE_B)
	} else if costPathA >= 0 {
		cost = costPathA + PRICE_A
	} else if costPathB >= 0 {
		cost = costPathB + PRICE_B
	}

	ctx.memo[target] = cost
	return cost
}
