# Answer: Delete Middle Node

## Problem

Delete a node in the middle (i.e., any node but the first and last node, not necessarily the exact middle) of a singly linked list, given only access to that node.

## Solution Approach

Since we don't have access to the `head`, we cannot traverse to the previous node to change its `Next` pointer.
Instead, we copy the data from the _next_ node to the _current_ node, and then delete the _next_ node.

1.  **Copy Data**: `node.Data = node.Next.Data`
2.  **Skip Next**: `node.Next = node.Next.Next`

### Diagram

```text
List: A -> B -> C -> D
Delete Node B.

1. Copy C's data to B.
   A -> C(was B) -> C -> D

2. Skip C.
   A -> C(was B) -> D
```

## Code Snippet

```go
func DeleteMiddleNode(node *Node) bool {
	if node == nil || node.Next == nil {
		return false // Cannot delete last node with this method
	}

	next := node.Next
	node.Data = next.Data
	node.Next = next.Next
	return true
}
```

## Complexity Analysis

- **Time Complexity**: O(1).
- **Space Complexity**: O(1).
