# Advent of Code 2024

My goal for this year's advent of code is to enjoy the puzzles and gain some familiarity with [Go](https://go.dev/). Thus, my process is as follows:
1. Solve the problem on my own first - i.e. without AI support.
2. Look on [Reddit](https://www.reddit.com/r/adventofcode/) for inspiration and alternative solutions.
3. Ask an LLM for code review, feedback and other possible improvements.

## Daily notes

### Day 06
- After learning about *iterators* in Go, I decided to write my own custom iterator for stepping through the map. Very satisfied with the final solution. 
- Happy that I found the word *pose* to represent both position and direction. I find problems easier to solve when they are named properly. 
- For part 02 I used *goroutines* for the first time to solve the problem concurrently. Getting it right took some time, because:
    1. I made unintentional race conditions, reading variables that were written to outside of the goroutine.
    2. I counted obstacles placed in the same position more than once.

Valuable GPT feedback:
- Avoid global locks that can cause contention between goroutines. For counting, flags or other operations that update a single value concurrently, the `iter/atomic` package has lock-free methods that can be used instead of a Mutex to reduce overhead.
- Use an empty struct instead of a boolean to test for presence (`map[Pose]struct{}`). Go does not have native sets, but an empty set occupies **zero bytes** and is a common way of implementing set features.
Go does not have native sets. The Use an empty struct instead of a bool. The map is used like a Set, but Go does not have that natively.
- Avoid deep copies if possible. It is expensive in terms of runtime and memory usage. I could perhaps have modified the `IsLoop(...)` function to take the obstacle position as an argument. However, this does not align easily with the iterator method for this problem. I did not implement this suggestion.
- Consider using a worker pool to limit the number of goroutines. Spawning too many goroutines may lead to memory pressure and bottlenecks. I did not implement this suggestion.

**Inspiration from others**
- Python: Very elegant use of complex numbers to represent both position and direction ([reddit](https://www.reddit.com/r/adventofcode/comments/1h7tovg/comment/m0o44m5))

### Day 05
I am very happy with my solution for today, although the process was a bit convoluted. I spent 2 hours trying to implement what I suddenly realized was and overly complicated and inefficient *sorting algorithm*. I should probably have caught on earlier, considering that I named my variables e.g. `comesBefore` or `comesAfter`. This was a good reminder of thouroughly understanding what type of problem it is before implementing a solution.

Nevertheless, I got to practise some important concepts:
- Closures - I wrote my first proper closure in Go for custom sorting of the list of numbers!
- Generics - On my way to the final solution, I got to practise Go generics.
- Primitive Obsession - My original solution smelled of [primitive obsession](https://wiki.c2.com/?PrimitiveObsession), so I implemented a few value objects.

Valuable ChatGPT feedback:
- Great suggestion by ChatGPT to replace my original `slices.Contains(...)` (O(n)) logic with map lookups (O(1)) since I was already creating maps for the sorting function. This greatly improved readability as well.

### Day 04
My original idea for today was to use a single line regex. However, it turns out that capturing overlapping patterns are not straight forward. For instance, the diagonal regex captures multiple lines into one match. Any other matches found on those lines are not captured. I am content with the scanning solution I came up with instead. Treating the input as a single line simplifies the problem somewhat. I could perhaps have generalized this solution, taking valid steps and words as function arguments.

**Inspiration from others**
- I really liked [this solution for Python on Reddit](https://www.reddit.com/r/adventofcode/comments/1h689qf/comment/m0bw4f7/), using a dictionary (map) and clever loops to iterate over all directions. I did not use it for my solution, but implementing something similar in Go would have been interesting.

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

- Errors
    - To create an error with a message, use `fmt.Errorf(...)`
    - Use `panic(...)` with an error message (string) for situations that won't be handled at runtime

- Enums
    - Not a native feature in Go, but can be implemented using the `iota` keyword - a counter that resets when encountering the `const` keyword. See [yourbasic.org](https://yourbasic.org/golang/iota/) for a practical example.