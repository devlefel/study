# Answer: Triple Step

## Problem

A child is running up a staircase with n steps and can hop either 1 step, 2 steps, or 3 steps at a time. Implement a method to count how many possible ways the child can run up the stairs.

## Solution Approach

This is a classic Dynamic Programming problem.
Let `f(n)` be the number of ways to reach step `n`.
To reach step `n`, the child could have come from:

- Step `n-1` (hopped 1 step)
- Step `n-2` (hopped 2 steps)
- Step `n-3` (hopped 3 steps)
  Therefore, `f(n) = f(n-1) + f(n-2) + f(n-3)`.
  Base cases:
- `f(0) = 1` (One way to be at the bottom: do nothing)
- `f(1) = 1`
- `f(2) = 2` (1+1, 2)

We should use **Memoization** to cache results and avoid re-calculating `f(n)` multiple times.

### Diagram

```text
f(3) = f(2) + f(1) + f(0)
     = 2 + 1 + 1 = 4

f(4) = f(3) + f(2) + f(1)
     = 4 + 2 + 1 = 7
```

## Code Snippet

```go
func TripleStep(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return countWays(n, memo)
}

func countWays(n int, memo []int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if memo[n] > -1 {
		return memo[n]
	}

	memo[n] = countWays(n-1, memo) + countWays(n-2, memo) + countWays(n-3, memo)
	return memo[n]
}
```

## Complexity Analysis

- **Time Complexity**: O(N).
- **Space Complexity**: O(N) (Recursion stack + memo array).
