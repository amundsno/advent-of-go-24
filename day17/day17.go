package day17

import (
	"fmt"
	"strconv"
	"strings"
)

type Computer struct {
	A, B, C         int
	program, output []int
}

func (c Computer) combo(operand int) int {
	switch operand {
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	case 7:
		panic("invalid combo operand")
	default:
		return operand
	}
}

func (c *Computer) execute(code, operand int) {
	// Case 3 is only used to loop back to start
	switch code {
	case 0:
		c.A = c.A >> c.combo(operand)
	case 1:
		c.B = c.B ^ operand
	case 2:
		c.B = c.combo(operand) % 8
	case 4:
		c.B = c.B ^ c.C
	case 5:
		c.output = append(c.output, c.combo(operand)%8)
	case 6:
		c.B = c.A >> c.combo(operand)
	case 7:
		c.C = c.A >> c.combo(operand)
	}
}

func (c *Computer) sweep() {
	for i := 0; i < len(c.program); i += 2 {
		code, operand := c.program[i], c.program[i+1]
		c.execute(code, operand)
	}
}

func (c *Computer) run() {
	for c.A > 0 {
		c.sweep()
	}
}

func (c Computer) String() string {
	outStr := make([]string, len(c.output))
	for i, val := range c.output {
		outStr[i] = strconv.Itoa(val)
	}
	return strings.Join(outStr, ",")
}

func Solve() {
	// Initialize manually
	computer := Computer{
		A:       0,
		program: []int{0},
	}

	computer.run()
	fmt.Printf("Part 01: %v\n", computer)
}
