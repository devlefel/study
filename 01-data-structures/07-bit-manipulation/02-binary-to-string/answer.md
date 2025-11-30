# Answer: Binary to String

## Problem

Given a real number between 0 and 1 (e.g., 0.72) that is passed in as a double, print the binary representation. If the number cannot be represented accurately in binary with at most 32 characters, print "ERROR".

## Solution Approach

To convert a decimal fraction to binary:

1.  Multiply by 2.
2.  If result >= 1, append "1" and subtract 1.
3.  If result < 1, append "0".
4.  Repeat until result is 0 or length > 32.

### Diagram

```text
Num = 0.625

1. 0.625 * 2 = 1.25
   >= 1? Yes. Append "1".
   Num = 1.25 - 1 = 0.25
   Binary: 0.1

2. 0.25 * 2 = 0.5
   >= 1? No. Append "0".
   Num = 0.5
   Binary: 0.10

3. 0.5 * 2 = 1.0
   >= 1? Yes. Append "1".
   Num = 1.0 - 1 = 0.0
   Binary: 0.101

4. Num is 0. Stop.
Result: 0.101
```

## Code Snippet

```go
import "strings"

func BinaryToString(num float64) string {
	if num >= 1 || num <= 0 {
		return "ERROR"
	}

	var sb strings.Builder
	sb.WriteString("0.")

	for num > 0 {
		if sb.Len() >= 32 {
			return "ERROR"
		}

		r := num * 2
		if r >= 1 {
			sb.WriteString("1")
			num = r - 1
		} else {
			sb.WriteString("0")
			num = r
		}
	}
	return sb.String()
}
```

## Complexity Analysis

- **Time Complexity**: O(1) (Max 32 iterations).
- **Space Complexity**: O(1) (Max 32 chars).
