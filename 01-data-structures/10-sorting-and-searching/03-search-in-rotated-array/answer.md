# Answer: Search in Rotated Array

## Problem

Given a sorted array of n integers that has been rotated an unknown number of times, write code to find an element in the array. You may assume that the array was originally sorted in increasing order.

## Solution Approach

We can use a modified **Binary Search**.

1.  Find `mid`.
2.  If `A[mid] == target`, return `mid`.
3.  Check which half is sorted:
    - If `A[left] < A[mid]`, then the **left** half is sorted.
      - If `A[left] <= target < A[mid]`, search left.
      - Else, search right.
    - If `A[left] > A[mid]`, then the **right** half is sorted.
      - If `A[mid] < target <= A[right]`, search right.
      - Else, search left.
    - If `A[left] == A[mid]`, we have duplicates (if allowed). If duplicates allowed, we might need to search both sides (O(N)). Assuming distinct for O(log N).

### Diagram

```text
Array: [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14]
Target: 5

Left=0 (15), Right=11 (14), Mid=5 (1).
Left (15) > Mid (1) -> Right half is sorted.
Is 5 in [1, 14]? Yes.
Search Right: Left=6 (3), Right=11 (14).

Mid=8 (5).
Found 5! Return 8.
```

## Code Snippet

```go
func SearchInRotatedArray(arr []int, target int) int {
	return search(arr, 0, len(arr)-1, target)
}

func search(a []int, left, right, x int) int {
	if right < left {
		return -1
	}

	mid := (left + right) / 2
	if a[mid] == x {
		return mid
	}

	if a[left] < a[mid] { // Left is normally ordered
		if x >= a[left] && x < a[mid] {
			return search(a, left, mid-1, x)
		} else {
			return search(a, mid+1, right, x)
		}
	} else if a[left] > a[mid] { // Right is normally ordered
		if x > a[mid] && x <= a[right] {
			return search(a, mid+1, right, x)
		} else {
			return search(a, left, mid-1, x)
		}
	} else { // Left == Mid (Duplicates)
		if a[mid] != a[right] { // If right is different, search it
			return search(a, mid+1, right, x)
		} else { // Else we have to search both
			result := search(a, left, mid-1, x)
			if result == -1 {
				return search(a, mid+1, right, x)
			}
			return result
		}
	}
}
```

## Complexity Analysis

- **Time Complexity**: O(log N) if distinct elements. O(N) with duplicates.
- **Space Complexity**: O(log N) (Recursion).
