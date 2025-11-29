# Answer: Prime Naive

## Problem

Implement a function that checks if a number is prime.

## Solution Approach

A naive approach is to iterate from 2 to n-1.
A better approach is to iterate from 2 to sqrt(n).

- If n is divisible by any number up to sqrt(n), it is not prime.
- We only need to check up to sqrt(n) because if n = a \* b, then one of a or b must be <= sqrt(n).

### Diagram

```text
IsPrime(17)
Check 2: 17 % 2 != 0
Check 3: 17 % 3 != 0
Check 4: 17 % 4 != 0
Stop (4 * 4 > 17)
Result: True
```

## Code Snippet

```go
import "math"

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
```

## Complexity Analysis

- **Time Complexity**: O(sqrt(N)).
- **Space Complexity**: O(1).
