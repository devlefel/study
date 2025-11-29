# Answer: Is Unique

## Problem

Implement an algorithm to determine if a string has all unique characters.

## Solution Approach

The most efficient way to solve this (assuming ASCII character set) is to use a boolean array (or bit vector) to track characters we've seen so far.

1.  **Check Length**: If the string length > 128 (for standard ASCII), it must have duplicates (Pigeonhole Principle). Return `false`.
2.  **Boolean Array**: Create a boolean array of size 128.
3.  **Iterate**: Loop through the string.
    - Get the ASCII value of the char.
    - If `seen[char]` is already `true`, return `false`.
    - Set `seen[char]` to `true`.
4.  **Return True**: If loop finishes without collision, return `true`.

![Is Unique Diagram](../../../02-system-design/images/is_unique.png)

## Code Snippet

```go
func IsUnique(str string) bool {
	if len(str) > 128 {
		return false
	}

	seen := make([]bool, 128)
	for _, char := range str {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
```

## Complexity Analysis

- **Time Complexity**: O(N), where N is the length of the string. Actually O(1) because the loop never runs more than 128 times.
- **Space Complexity**: O(1) for the fixed-size boolean array.
