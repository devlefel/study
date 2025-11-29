# Answer: Remove Dups

## Problem

Write a code to remove duplicates from an unsorted linked list.

## Solution Approach

1.  **Hash Table (Buffer)**: Keep a set of values we've seen. Iterate through the list.
    - If `node.Data` is in the set, remove the node (change `prev.Next` to `node.Next`).
    - Else, add `node.Data` to the set.
2.  **No Buffer (Runner)**: If no extra space is allowed, use two pointers.
    - `Current` iterates normally.
    - `Runner` checks all subsequent nodes for duplicates of `Current`.
    - Time Complexity becomes O(N^2).

### Diagram

```text
List: 1 -> 2 -> 1 -> 3

Step 1: Seen {1}
  Curr: 1. Next is 2.

Step 2: Seen {1, 2}
  Curr: 2. Next is 1.

Step 3: Seen {1, 2}. 1 is in Set!
  DELETE 1.
  Prev(2) -> Next(3).
  List: 1 -> 2 -> 3
```

## Code Snippet

```go
func RemoveDups(head *Node) {
	if head == nil {
		return
	}
	seen := make(map[int]bool)
	previous := head
	seen[head.Data] = true
	current := head.Next

	for current != nil {
		if seen[current.Data] {
			previous.Next = current.Next // Skip duplicate
		} else {
			seen[current.Data] = true
			previous = current
		}
		current = current.Next
	}
}
```

## Complexity Analysis

- **Time Complexity**: O(N) with buffer. O(N^2) without.
- **Space Complexity**: O(N) with buffer. O(1) without.
