# Advent of Code 2024

My goal for this year's advent of code is to enjoy the puzzles and gain some familiarity with [Go](https://go.dev/). Thus, my process is as follows:
1. Solve the problem on my own first - i.e. without AI support.
2. Look on [Reddit](https://www.reddit.com/r/adventofcode/) for inspiration and alternative solutions.
3. Ask an LLM for code review, feedback and other possible improvements.

## Daily notes

### Day 16
The first part use Dijkstra (BFS with min priority queue) to explore paths from S to E that minimize the score. The second part use recursive DFS on the paths explored by Dijkstra to return tiles on a path which minimize the score.

Happy with my final solution, although it took some time to arrive at. I spent a lot of time making things more complex than they had to be, trying to optimize too soon.

*Other solutions*
- [This Python solution](https://www.reddit.com/r/adventofcode/comments/1hfboft/comment/m2bcfmq) solves both parts in the same loop by storing the entire path traversed so far for each step in Dijkstra. The paths that reach the end are added to a set, such that only unique tiles remain. I consider it elegant in terms of being easy to reason about, but perhaps not so memory efficient.
- [This Rust solution](https://www.reddit.com/r/adventofcode/comments/1hfboft/comment/m2bk6i4) uses the same ideas as mine. However, it uses a bucket queue instead of a min priority queue (heap) to optimize further. Could be interesting to read more about this.

### Day 15
I feel like all the practise with recursive solutions are paying off! Very happy with my intuition and solution for today. Visualizing the recursion as a tree really helps in designing the algorithm.

For the first part, I wrote a recursive function that applied the move immediately if it was valid. For the second part, I had to verify if the entire chain could be moved before doing so. I initially solved this by modifying the recursive function to return `nil` if the move was invalid or a function to apply the move otherwise. This way I could defer the actual movement execution. While perhaps clever, it was not very readable or intuitive. 

I refactored to a more readable approach, consisting of three steps:
1. Recursively link the tiles that will be moved together.
2. Iterate over the linked tiles to see if any would move into a wall.
3. If not, iterate over the linked tiles again to move them.

Although a few more steps are required, this approach is much easier to reason about. It is still more than fast enough.

### Day 14
Go implements the modulo (%) operator different than Python. In Python, the remainder is chosen to be the smallest *positive* number `r` that satisfies `a = qb + r, q=a/b, abs(r) < abs(n)`. In Go, the remainder is chosen to be the smallest number with the same sign as the dividend `a`. To get Python-like behaviour in Go, we can add the divisor to the result and do another modulo to always make it positive: `((a % b) + b) % b`.

### Day 13
The wording of this problem lead me to believe that this was a textbook dynamic programming problem. I solved part 01 that way, but quickly caught on that this was indeed a system of linear equations that could be solved using linear algebra. All the claw machines in the problem input had buttons that were linearly independent, meaning that there will be only one possible solution for how many times to press each button. However, if we consider a hypothetical case where the buttons had been linearly dependent, there could still have been a unique optimal solution to the problem.

Consider the following hypothetical problems:

1. A=[1, 1], B=[2, 2] and T=[5, 5]. Even though A and B are linearly dependent, the optimal solution is pressing B 2 times and A 1 time for a cost of 5 tokens.
2. A=[7 7], B=[2, 2] and T=[20, 20]. In this case, pressing A is more cost efficient than B, and the optimal solution is pressing A 2 times and B 3 times for a cost of 9 tokens.
3. A=[4, 4], B=[3, 3] and T=[14, 14]. Here pressing A 2 times and B 2 times give the optimal solution of 8 tokens.

Note that it is not a matter of simply pressing the most cost efficient button as much as possible without exceeding the target. If we had done that for the third problem above, we would have pressed B 4 times to end up at (12, 12). There is no way to reach the target from (12, 12) without backtracking to pressing B 2 times, followed by A 2 times.

Turns out we can use Linear Diophantine Equations and The Euclidian Algorithm to find the (A, B) pair which minimize the cost mathematically - i.e. we do not have to iterate through all valid (A, B) pairs to find the minimal cost.

I made a [post on Reddit](https://www.reddit.com/r/adventofcode/comments/1i20wpg) to explain the solution in more detail.

Apart from this, I got to review both dynamic programming and some linear algebra, which felt good.

### Day 12
Part 01 was a pretty straight forward flood fill algorithm. Instead of recursion, I solved it iteratively using a queue. Initially I used counters to have the algorithm return the area and perimeter directly, but I refactored it to return the points within a region instead for the second part. I quickly understood that it would be easier to count the number of corners instead of sides, however finding the exact method of counting took some time. I experimented with operations on the points accross the perimeter, but this turned out to be complex. My final solution is inspired by [this comment on Reddit](https://www.reddit.com/r/adventofcode/comments/1hcdnk0/comment/m1nio0w) although the actual implementation is a little different. Consider each point in the region. It is an outside corner if the point to the N & E (or E & S, S & W, W & N) is not within the region. It is an inside corner if those points are within the region and the diagonal between them are not - i.e. N & E inside, but NE not inside.

*Valuable feedback:*
- I have to remember that I can iterate over key, value pairs in maps without calling `maps.All(...)`.

*Other solutions*
- Blown away by [this solution on Reddit](https://www.reddit.com/r/adventofcode/comments/1hcdnk0/comment/m1pb0m7) that uses a kernel / convolution matrix to perform edge detection mathematically. Ingenious. 
- I appreciate the solutions that use math and complex numbers to navigate the matrix. This is helpful, as multiplying by 1j is the equivalent of rotating 90 degrees clockwise. This is another way of spotting corners.

### Day 11
Another recursive DFS algorithm, but this time with memoization. The idea is to count the number of leaf nodes after branching N times from the root node (stone). I used a closure to capture the cache variable and improve readability. My final solution is inspired by [this comment on Reddit](https://www.reddit.com/r/adventofcode/comments/1hbm0al/comment/m1i36gs).

It took a while to complete today's problem, because I tried to over-optimize the caching. I wanted to store the number of stones following each blink in a long chain, e.g.: `stone: [countBlink0, countBlink1, countBlink2, ...]`. However, this was harder to reason about and turned out more complex than I imagined to begin with. Thus, I ended up keeping it simple while perhaps doing a bit more computation.

*Valuable feedback:*
- The `sync.Map` type can be used for concurrent memoization
- The splitting can be done without converting to strings to make it more efficient. Two methods:
    1. Use `% 10` to get the last digit and append it to a list of digits, then divide by 10 until the input integer is 0. No need for type conversion, but there is some overhead in performance and memory due to slice manipulation (ChatGPT suggestion).
    2. Use logarithms to find the number of digits, then divison to get the first part and modulus for the second ([Reddit comment](https://www.reddit.com/r/adventofcode/comments/1hbm0al/comment/m1hez5b)). The only overhead comes from float conversion.

*Other solutions*
- This [proof-of-concept on Reddit](https://www.reddit.com/r/adventofcode/comments/1hbm0al/comment/m1jblcs/) uses a *state transition matrix* to apply all the blinks in one go. The possible values have to be pre-computed, but it is a very elegant solution. The matrix at point *(i, j)* describes how many stones transition from number *i* to *j* in one step. Raising it to the power of X blinks yields a matrix that describes how many stones transition from *i* to *j* in X blinks. Then taking the dot product between the initial state vector (`[1 for s in stones, else 0]`) and this matrix you get the end state in one step. Very clever!

### Day 10
Straight forward recursive depth first search algorithm for both parts of today's problem.

*Valuable feedback:*
- There is no need for passing pointers to maps, as maps are referenced types.
- Prefer methods over free functions when logic is tightly coupled to the type.
- Use *state* or *context* struct variables instead of passing multiple separate arguments to the same function. This improves readability.
- Try to use more domain specific language for readability. 

### Day 09
My final solution for part 02 is inspired by [this comment on Reddit](https://www.reddit.com/r/adventofcode/comments/1ha27bo/comment/m15wwre/). I was on the right track - parsing the input to memory blocks and comparing their sizes - but I got stuck in the complexity of actually moving them around. Because we only care about the final checksum, there is no need to move the memory blocks. Setting the index/position they would have had if they were moved is enough. We can then use sum of a divergent sequence to compute how each memory block contributes to the final checksum.

### Day 08
One of my fastest solve times so far. My solution in O(n**2) time. The code is not as modular as I like, but it is simple to understand and works.

### Day 07
Recursive depth first search algorithm came easy for today's problem. Made a silly mistake that had me debugging part 02 for a long time. If the target was reached early, I did not consider the remaining chain. This resulted in two of the input rows to return `true` too soon.

*Valuable feedback*
- String splitting with `strings.Fields()` if the separator is whitespace.

*Other*:
- Compare files in VSCode by right clicking on a file to "Select for Compare", and another to "Compare with Selected" ([source](https://vscode.one/diff-vscode/)).

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
    - You can append to the beginning of a slice with: `slice = append([]int{value}, slice...)`
    
- Testing
    - Use table-driven tests with `t.Run(...)` to create subtests
    - Use non-fatal assertions with `t.Errorf(...)` instead of `t.Fatalf(...)` to ensure all test scenarios are run, even if one fails

- Errors
    - To create an error with a message, use `fmt.Errorf(...)`
    - Use `panic(...)` with an error message (string) for situations that won't be handled at runtime

- Enums
    - Not a native feature in Go, but can be implemented using the `iota` keyword - a counter that resets when encountering the `const` keyword. See [yourbasic.org](https://yourbasic.org/golang/iota/) for a practical example.

- Global variables
    - Global variables can be declared with the `var name = ` syntax. 

- Type conversion
    - Iterating over a string, yields a bite or rune for each position. To get the integer value of a byte/rune, subtract `'0'` to offset the ASCII conversion (`int(rune - '0')`)
    
- Composition over inheritance
    - Go does not have inheritance; it favours *composition*. You can *embed* a type within a struct by using the type as a field without a name. The embedded type's methods are *promoted* to the outer struct, making them callable directly. This allows us to implement interfaces for a new struct by simply embedding a type that already satisfies that interface ([source](https://go.dev/doc/effective_go#embedding)).