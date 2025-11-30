# Answer: Magic Index

## Problem

A magic index in an array `A[0...n-1]` is defined to be an index such that `A[i] = i`. Given a sorted array of distinct integers, write a method to find a magic index, if one exists, in array A.

## Solution Approach

Since the array is sorted, we can use a modification of **Binary Search**.

- Check `mid`.
- If `A[mid] == mid`, return `mid`.
- If `A[mid] > mid`:
  - Since elements are distinct and sorted, for all `j > mid`, `A[j] >= A[mid] + (j - mid) > mid + (j - mid) = j`.
  - So `A[j] > j` for all `j > mid`.
  - The magic index must be on the left side.
- If `A[mid] < mid`:
  - Similarly, the magic index must be on the right side.

### Diagram

```text
Index: 0   1   2   3   4
Value: -10 -5  2   20  30

Mid = 2. A[2] = 2. Magic Index Found!

Index: 0   1   2   3   4
Value: -10 -5  1   20  30

Mid = 2. A[2] = 1. 1 < 2.
Search Right: [3, 4]
Mid = 3. A[3] = 20. 20 > 3.
Search Left (of [3,4]): Empty.
Return -1.
```

## Code Snippet

```go
func MagicIndex(arr []int) int {
	return search(arr, 0, len(arr)-1)
}

func search(arr []int, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if arr[mid] == mid {
		return mid
	} else if arr[mid] > mid {
		return search(arr, start, mid-1)
	} else {
		return search(arr, mid+1, end)
	}
}
```

## Complexity Analysis

- **Time Complexity**: O(log N).
- **Space Complexity**: O(log N) (Recursion).
