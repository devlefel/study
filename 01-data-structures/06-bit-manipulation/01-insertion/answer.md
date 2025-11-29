# Answer: Insertion

## Problem

You are given two 32-bit numbers, N and M, and two bit positions, i and j. Write a method to insert M into N such that M starts at bit j and ends at bit i.

## Solution Approach

1.  **Clear Bits**: Clear the bits in N from j to i. We need a mask like `11100000011`.
    - Left part: `~0 << (j + 1)` -> `11100000000`
    - Right part: `(1 << i) - 1` -> `00000000011`
    - Mask: `Left | Right`
    - `N_cleared = N & Mask`
2.  **Shift M**: Shift M to the correct position.
    - `M_shifted = M << i`
3.  **Merge**: Combine them.
    - `Result = N_cleared | M_shifted`

### Diagram

```text
N = 10000000000
M = 10011
i = 2, j = 6

1. Mask Creation:
   11111111111
   Clear j through i (6 to 2)
   11110000011

2. Clear N:
   10000000000
 & 11110000011
   -----------
   10000000000 (Bits 2-6 are already 0 here, but usually they wouldn't be)

3. Shift M:
   10011 << 2 = 1001100

4. Merge:
   10000000000
 | 00001001100
   -----------
   10001001100
```

## Code Snippet

```go
func Insertion(n, m int, i, j int) int {
	// Create mask to clear bits i through j in n
	allOnes := ^0 // Sequence of 1s

	// Left part of mask (1s before position j)
	left := allOnes << (j + 1)

	// Right part of mask (1s after position i)
	right := (1 << i) - 1

	// Combine masks
	mask := left | right

	// Clear bits j through i
	nCleared := n & mask

	// Shift m into correct position
	mShifted := m << i

	return nCleared | mShifted
}
```

## Complexity Analysis

- **Time Complexity**: O(1).
- **Space Complexity**: O(1).
