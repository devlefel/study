# Answer: Return Kth to Last

## Problem

Implement an algorithm to find the kth to last element of a singly linked list.

## Solution Approach

Use the **Runner Technique** (Two Pointers).

1.  **P1 and P2**: Start both pointers at the head.
2.  **Move P1**: Move P1 `k` steps forward.
    - If P1 hits `nil` before `k` steps, the list is too short.
3.  **Move Both**: Move P1 and P2 at the same speed until P1 hits the end (`nil`).
4.  **Result**: P2 is now at the `k`th node from the end. Return `P2.Data`.

### Diagram

```text
List: 1 -> 2 -> 3 -> 4 -> 5, K=2

Start:
P1, P2
v
1 -> 2 -> 3 -> 4 -> 5

Move P1 (2 steps):
P2        P1
v         v
1 -> 2 -> 3 -> 4 -> 5

Move Both until P1 is nil:
          P2        P1
          v         v
1 -> 2 -> 3 -> 4 -> 5  (P1 is at 5)
               P2        P1
               v         v
1 -> 2 -> 3 -> 4 -> 5 -> nil

Result: P2.Data is 4. Correct.
```

## Code Snippet

```go
func ReturnKthToLast(head *Node, k int) int {
	p1 := head
	p2 := head

	// Move p1 k steps ahead
	for i := 0; i < k; i++ {
		if p1 == nil {
			return 0 // List too short
		}
		p1 = p1.Next
	}

	// Move both until p1 hits end
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2.Data
}
```

## Complexity Analysis

- **Time Complexity**: O(N).
- **Space Complexity**: O(1).
