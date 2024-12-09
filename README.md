# Advent of Code 2024

## Daily notes

### Day 04
My original idea for today was to use a single line regex. However, it turns out that capturing overlapping patterns are not straight forward. For instance, the diagonal regex captures multiple lines into one match. Any other matches found on those lines are not captured.

### Day 03
Learned how to work with regular expressions in Go. For the second part, I could not find a regular expression to extract the valid multiplication elements without using lookahead features that are not supported by Go natively. Still happy with the two-step solution.

### Day 02
Solved today with TDD and SOLID principles in mind. Though it took some time to realise, I am happy that I could extend upon the logic from the first part when solving the second.

**Inspiration from others**
- Elegant way of checking if the levels are only increasing or decreasing: `(r[0]-r[1])*(r[i]-r[i+1]) > 0`. If the result is negative, the trajectory has changed. If the result is zero, there is no increase or decrease ([source](https://github.com/mnml/aoc/blob/5e49f2c1b4839d4a115131ac21bf845caf700ccd/2024/02/1.go#L35)).

### Day 01
Trying out Go for the first time, after completing [A Tour of Go](https://go.dev/tour/welcome/1) a month back. The main challenge for today was getting familiar with syntax and file I/O operations. ChatGPT provided useful feedback on testing and writing more idiomatic Go.

## Feedback and tips to remember
- Regex
    - Regular expressions in Go are implemented using a library (RE2) that is designed for linear runtime and safety. Thus, features like lookarounds are not supported ([source](https://www.honeybadger.io/blog/a-definitive-guide-to-regular-expressions-in-go/)).
    - `(?s)` - Single-line pattern matching flag. Treat the input as a single line, regardless of line breaks.
	- `(?:)` - Non selecting group. Useful for grouping tokens, without catching them.
	- `.*?`- Lazy quantifiers. By default, regex quantifiers are *greedy*, meaning they will match as much as they can. Making them lazy is useful, when we want the pattern to stop matching at specified tokens.
    
- Slice manipulation
    - On making deep copies. The `copy(dst, src)` function copies only what there is room for in `dst`, meaning size must be preallocated. Instead, one can append the `src` elements to a *nil* slice using the ellipsis `...` operator: `dst := append([]int{}, src... )`
    
- Testing
    - Use table-driven tests with `t.Run(...)` to create subtests
    - Use non-fatal assertions with `t.Errorf(...)` instead of `t.Fatalf(...)` to ensure all test scenarios are run, even if one fails