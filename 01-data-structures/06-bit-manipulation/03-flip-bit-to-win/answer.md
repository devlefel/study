# Answer: Flip Bit to Win

## Problem

You have an integer and you can flip exactly one bit from a 0 to a 1. Write code to find the length of the longest sequence of 1s you could create.

## Solution Approach

We can iterate through the bits of the number. We need to keep track of:

1.  `currentLength`: Length of 1s sequence currently being counted.
2.  `previousLength`: Length of the _previous_ 1s sequence (separated by a single 0).
3.  `maxLength`: Max length found so far.

When we see a 1:

- Increment `currentLength`.

When we see a 0:

- Update `maxLength` with `previousLength + 1 + currentLength`.
- If the _next_ bit is also 0 (meaning we have `00`), then `previousLength` becomes 0 (cannot bridge two 0s).
- Else (we have `101`), `previousLength` becomes `currentLength`.
- Reset `currentLength` to 0.

### Diagram

```text
Input: 11011101111 (1775)

Bit 0 (1): Cur=1, Prev=0, Max=1
Bit 1 (1): Cur=2, Prev=0, Max=2
Bit 2 (1): Cur=3, Prev=0, Max=3
Bit 3 (1): Cur=4, Prev=0, Max=4
Bit 4 (0): Found 0.
           Max = max(4, 0 + 1 + 4) = 5.
           Prev = 4. Cur = 0.
Bit 5 (1): Cur=1, Prev=4, Max=5
Bit 6 (1): Cur=2, Prev=4, Max=5
Bit 7 (1): Cur=3, Prev=4, Max=5
           Current potential = 4 (prev) + 1 (flip) + 3 (cur) = 8.
           Max = 8.
...
```

## Code Snippet

```go
func FlipBitToWin(a int) int {
	if ^a == 0 { // All 1s
		return 64 // Or 32 depending on int size
	}

	currentLen := 0
	previousLen := 0
	maxLen := 1 // Can always have at least 1 by flipping a 0

	for a != 0 {
		if (a & 1) == 1 {
			currentLen++
		} else {
			// If next bit is 1, we can merge.
			// Actually simpler logic:
			// prevLen becomes currentLen (if we flip this 0)
			// But if the *next* bit is 0, we can't merge, so prev should be 0.
			// We can check this by (a & 2).
			if (a & 2) == 0 {
				previousLen = 0
			} else {
				previousLen = currentLen
			}
			currentLen = 0
		}
		maxLen = max(maxLen, previousLen + currentLen + 1)
		a >>= 1
	}
	return maxLen
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
```

## Complexity Analysis

- **Time Complexity**: O(b), where b is the number of bits (32 or 64). Effectively O(1).
- **Space Complexity**: O(1).
