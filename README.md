# Advent of Code 2024

## Daily notes

### Day 01
Trying out Go for the first time, after completing [A Tour of Go](https://go.dev/tour/welcome/1) a month back. The main challenge for today was getting familiar with syntax and file I/O operations. ChatGPT provided useful feedback on testing and writing more idiomatic Go.

### Day 02
Solved today with TDD and SOLID principles in mind. Though it took some time to realise, I am happy that I could extend upon the logic from the first part when solving the second.

**Inspiration from others**
- Elegant way of checking if the levels are only increasing or decreasing: `(r[0]-r[1])*(r[i]-r[i+1]) > 0`. If the result is negative, the trajectory has changed. If the result is zero, there is no increase or decrease ([source](https://github.com/mnml/aoc/blob/5e49f2c1b4839d4a115131ac21bf845caf700ccd/2024/02/1.go#L35)).

## Feedback and tips to remember
- Slice manipulation
    - On making deep copies. The `copy(dst, src)` function copies only what there is room for in `dst`, meaning size must be preallocated. Instead, one can append the `src` elements to a *nil* slice using the ellipsis `...` operator: `dst := append([]int{}, src... )`
    
- Testing
    - Use table-driven tests with `t.Run(...)` to create subtests
    - Use non-fatal assertions with `t.Errorf(...)` instead of `t.Fatalf(...)` to ensure all test scenarios are run, even if one fails