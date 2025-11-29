# Answer: Minimal Tree

## Problem

Given a sorted (increasing order) array with unique integer elements, write an algorithm to create a binary search tree with minimal height.

## Solution Approach

To minimize height, we need to balance the tree. The root should be the **middle** element of the array.

1.  **Find Middle**: `mid = (start + end) / 2`.
2.  **Create Root**: `node = new Node(arr[mid])`.
3.  **Recurse Left**: `node.Left = createMinimalBST(arr, start, mid - 1)`.
4.  **Recurse Right**: `node.Right = createMinimalBST(arr, mid + 1, end)`.

### Diagram

```text
Array: [1, 2, 3, 4, 5, 6, 7]

1. Mid is 4. Root = 4.
   Left Array: [1, 2, 3]
   Right Array: [5, 6, 7]

2. Left Mid is 2. Left Child = 2.
   Left of 2: [1] -> Child 1
   Right of 2: [3] -> Child 3

3. Right Mid is 6. Right Child = 6.
   Left of 6: [5] -> Child 5
   Right of 6: [7] -> Child 7

      4
    /   \
   2     6
  / \   / \
 1   3 5   7
```

## Code Snippet

```go
func MinimalTree(arr []int) *Node {
	return createMinimalBST(arr, 0, len(arr)-1)
}

func createMinimalBST(arr []int, start, end int) *Node {
	if end < start {
		return nil
	}
	mid := (start + end) / 2
	node := &Node{Data: arr[mid]}
	node.Left = createMinimalBST(arr, start, mid-1)
	node.Right = createMinimalBST(arr, mid+1, end)
	return node
}
```

## Complexity Analysis

- **Time Complexity**: O(N).
- **Space Complexity**: O(N) (Recursion stack O(log N) + Result O(N)).
