# Answer: Sorted Merge

## Problem

You are given two sorted arrays, A and B, where A has a large enough buffer at the end to hold B. Write a method to merge B into A in sorted order.

## Solution Approach

If we merge from the front, we would need to shift elements in A to make room for elements from B, which is O(N^2).
Instead, we should **merge from the back**.

1.  Compare the last element of A (index `countA - 1`) and the last element of B (index `countB - 1`).
2.  Place the larger element at the end of A's buffer (index `countA + countB - 1`).
3.  Decrement indices and repeat.
4.  If B runs out, we are done (A is already in place).
5.  If A runs out, copy remaining B to A.

### Diagram

```text
A = [1, 3, 5, _, _, _]
B = [2, 4, 6]

Compare 5 (A) and 6 (B). 6 > 5.
A = [1, 3, 5, _, _, 6]
B used: 6

Compare 5 (A) and 4 (B). 5 > 4.
A = [1, 3, 5, _, 5, 6]
A used: 5

Compare 3 (A) and 4 (B). 4 > 3.
A = [1, 3, 5, 4, 5, 6]
B used: 4

... and so on.
```

## Code Snippet

```go
func SortedMerge(a []int, b []int, countA int, countB int) {
	indexA := countA - 1
	indexB := countB - 1
	indexMerged := countA + countB - 1

	for indexB >= 0 {
		if indexA >= 0 && a[indexA] > b[indexB] {
			a[indexMerged] = a[indexA]
			indexA--
		} else {
			a[indexMerged] = b[indexB]
			indexB--
		}
		indexMerged--
	}
}
```

## Complexity Analysis

- **Time Complexity**: O(A + B).
- **Space Complexity**: O(1) (In-place).
