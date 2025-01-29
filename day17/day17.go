package day17

import (
	"fmt"
	"strconv"
	"strings"
)

type Computer struct {
	A, B, C, ip  int
	program, out []int
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

func (c *Computer) execute(opcode, operand int) {
	switch opcode {
	case 0:
		c.adv(operand)
	case 1:
		c.bxl(operand)
	case 2:
		c.bst(operand)
	case 3:
		c.jnz(operand)
	case 4:
		c.bxc()
	case 5:
		c.ut(operand)
	case 6:
		c.bdv(operand)
	case 7:
		c.cdv(operand)
	}
}

func (c Computer) dv(operand int) int {
	pow := c.combo(operand) - 1
	if pow < 0 {
		return c.A / (2 >> -pow)
	}
	return c.A / (2 << pow)
}

func (c *Computer) adv(operand int) {
	c.A = c.dv(operand)
}

func (c *Computer) bxl(operand int) {
	c.B = c.B ^ operand
}

func (c *Computer) bst(operand int) {
	c.B = c.combo(operand) % 8
}

func (c *Computer) jnz(operand int) {
	if c.A != 0 && c.ip != operand {
		c.ip = operand - 2
	}
}

func (c *Computer) bxc() {
	c.B = c.B ^ c.C
}

func (c *Computer) ut(operand int) {
	c.out = append(c.out, c.combo(operand)%8)
}

func (c *Computer) bdv(operand int) {
	c.B = c.dv(operand)
}

func (c *Computer) cdv(operand int) {
	c.C = c.dv(operand)
}

func (c *Computer) run() {
	for c.ip < len(c.program) {
		opcode := c.program[c.ip]
		operand := c.program[c.ip+1]
		c.execute(opcode, operand)
		c.ip += 2
	}
}

func (c Computer) output() string {
	outStr := make([]string, len(c.out))
	for i, val := range c.out {
		outStr[i] = strconv.Itoa(val)
	}
	return strings.Join(outStr, ",")
}

func Solve() {
	// Initialize manually
	computer := Computer{
		A:       0,
		B:       0,
		C:       0,
		ip:      0,
		program: []int{0},
	}

	computer.run()
	fmt.Printf("Part 01: %v\n", computer.output())
}
