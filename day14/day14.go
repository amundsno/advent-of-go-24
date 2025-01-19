package day14

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Vec struct {
	x, y int
}

type Bot struct {
	position, speed Vec
}

const WIDTH, HEIGHT int = 101, 103

// const WIDTH, HEIGHT int = 11, 7

func (b Bot) PositionAfterSeconds(sec int) Vec {
	return Vec{
		x: mod(b.position.x+b.speed.x*sec, WIDTH),
		y: mod(b.position.y+b.speed.y*sec, HEIGHT)}
}

func QuadrantIndex(position Vec) int {
	if position.x < WIDTH/2 {
		if position.y < HEIGHT/2 {
			return 0
		}
		return 2
	}
	if position.y < HEIGHT/2 {
		return 1
	}
	return 3
}

func Solve(inputPath string) {
	bots := ParseInput(inputPath)
	quadrantCounts := make([]int, 4)
	for _, bot := range bots {
		pos := bot.PositionAfterSeconds(100)

		// Do not count bots in the middle
		if pos.x == WIDTH/2 || pos.y == HEIGHT/2 {
			continue
		}

		quadrantCounts[QuadrantIndex(pos)]++
	}

	safetyFactor := 1
	for _, count := range quadrantCounts {
		safetyFactor *= count
	}

	fmt.Printf("Part 01: %v\n", safetyFactor)

}

// Positive remainder modulo
func mod(a, b int) int {
	return ((a % b) + b) % b
}

func ParseInput(inputPath string) []Bot {
	rows := utils.ReadFileTo1D(inputPath)
	pattern := regexp.MustCompile(`.*?,?(-?\d+).*?`)

	bots := make([]Bot, len(rows))
	for i, row := range rows {
		matches := pattern.FindAllStringSubmatch(row, -1)
		x, _ := strconv.Atoi(matches[0][1])
		y, _ := strconv.Atoi(matches[1][1])
		vx, _ := strconv.Atoi(matches[2][1])
		vy, _ := strconv.Atoi(matches[3][1])

		bots[i] = Bot{
			position: Vec{x, y},
			speed:    Vec{vx, vy},
		}
	}

	return bots
}
