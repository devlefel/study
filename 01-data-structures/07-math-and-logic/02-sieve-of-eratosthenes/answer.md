# Answer: Sieve of Eratosthenes

## Problem

Generate a list of primes up to a given number `max`.

## Solution Approach

The Sieve of Eratosthenes is a highly efficient way to generate a list of primes.

1.  Create a boolean array `isPrime` of size `max + 1`, initialized to all true.
2.  Set `isPrime[0]` and `isPrime[1]` to false.
3.  Iterate from `p = 2` to `sqrt(max)`.
    - If `isPrime[p]` is true:
      - Mark all multiples of `p` (starting from `p*p`) as false.
      - `isPrime[i] = false` for `i = p*p, p*p + p, ...`

### Diagram

```text
Max = 10
Array: [T, T, T, T, T, T, T, T, T, T, T] (Indices 0-10)
0, 1 -> False

p = 2 (Prime)
Mark multiples: 4, 6, 8, 10 -> False
Array: [F, F, T, T, F, T, F, T, F, T, F]

p = 3 (Prime)
Mark multiples: 9 -> False (6 already false)
Array: [F, F, T, T, F, T, F, T, F, F, F]

p = 4 (Not Prime, skip)

Stop (sqrt(10) approx 3.16)

Primes: 2, 3, 5, 7
```

## Code Snippet

```go
import "math"

func SieveOfEratosthenes(max int) []bool {
	if max < 0 {
		return nil
	}

	// Initialize all to true
	flags := make([]bool, max+1)
	for i := range flags {
		flags[i] = true
	}

	flags[0] = false
	flags[1] = false

	limit := int(math.Sqrt(float64(max)))
	for prime := 2; prime <= limit; prime++ {
		if flags[prime] {
			// Cross off multiples starting at prime * prime
			for i := prime * prime; i <= max; i += prime {
				flags[i] = false
			}
		}
	}

	return flags
}
```

## Complexity Analysis

- **Time Complexity**: O(N log log N).
- **Space Complexity**: O(N).
