# Answer: URLify

## Problem

Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.

## Solution Approach

In languages with mutable strings (like C++ or Java char arrays), we do this **in-place** by working backwards.
In Go, strings are immutable, so we typically use a `strings.Builder` or a byte slice.

**Two-Pass Approach (In-Place Logic)**:

1.  **Count Spaces**: Count the number of spaces in the first `trueLength` characters.
2.  **Calculate New Length**: `newLength = trueLength + spaces * 2`.
3.  **Reverse Edit**: Start from the end of the string.
    - If char is a space, write '0', '2', '%' and move index back by 3.
    - If char is not a space, copy it and move index back by 1.

### Diagram

```text
Input:  "Mr John Smith    " (Length 13, Buffer 17)
                      ^ True Length
Pointer P1 at 13 (h)
Pointer P2 at 17 (end)

Move 'h' to P2. P1--, P2--
...
When P1 hits ' ' (space):
  Write '0' at P2, P2--
  Write '2' at P2, P2--
  Write '%' at P2, P2--
  P1--
```

## Code Snippet

```go
import "strings"

func URLify(str string, length int) string {
	str = str[:length] // Trim to true length
	var sb strings.Builder
	for _, char := range str {
		if char == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}
```

## Complexity Analysis

- **Time Complexity**: O(N), where N is the true length of the string.
- **Space Complexity**: O(N) in Go (creating new string). O(1) in C++ (in-place).
