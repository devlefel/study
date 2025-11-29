# Answer: Check Permutation

## Problem

Given two strings, write a method to decide if one is a permutation of the other.

## Solution Approach

Two strings are permutations if they have the same characters with the same frequency.

1.  **Check Length**: If lengths differ, they cannot be permutations.
2.  **Character Count**: Use an array (for ASCII) or Map (for Unicode) to count char frequencies in the first string.
3.  **Decrement**: Iterate through the second string, decrementing the counts in the array.
4.  **Verify**: If any count drops below zero, return `false`.

### Diagram

```text
String 1: "dog"
[d:1, o:1, g:1]  <-- Frequency Map

String 2: "god"
'g': [d:1, o:1, g:0]
'o': [d:1, o:0, g:0]
'd': [d:0, o:0, g:0]

All counts 0 -> True
```

## Code Snippet

```go
func CheckPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)
	for _, char := range s1 {
		counts[char]++
	}

	for _, char := range s2 {
		counts[char]--
		if counts[char] < 0 {
			return false
		}
	}
	return true
}
```

## Complexity Analysis

- **Time Complexity**: O(N), where N is the length of the strings.
- **Space Complexity**: O(1) if charset is fixed (e.g., 128 for ASCII). O(N) if using a map for arbitrary Unicode.
