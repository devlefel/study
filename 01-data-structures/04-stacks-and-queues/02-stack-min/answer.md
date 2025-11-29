# Answer: Stack Min

## Problem

How would you design a stack which, in addition to push and pop, has a function `min` which returns the minimum element? Push, pop and min should all operate in O(1) time.

## Solution Approach

We cannot iterate through the stack to find the min (that would be O(N)).
We need to store the minimum state _at each level_.

1.  **Node with Min**: Each node in the stack stores `val` and `minSoFar`.
    - `minSoFar = min(val, nextNode.minSoFar)`
2.  **Separate Min Stack**: Keep a second stack that only stores the minimums.
    - Push: If `val <= minStack.Peek()`, push to `minStack`.
    - Pop: If `val == minStack.Peek()`, pop from `minStack`.
    - _Pros_: More space efficient if data is not sorted (fewer duplicates in min stack).

### Diagram

```text
Main Stack: [5, 6, 3, 7]
Min Stack:  [5, 3]

Push 5: MinStack [5]
Push 6: 6 > 5. MinStack [5]
Push 3: 3 < 5. MinStack [5, 3]
Push 7: 7 > 3. MinStack [5, 3]

Pop 7: Val 7 != 3. MinStack [5, 3]
Pop 3: Val 3 == 3. MinStack [5]
```

## Code Snippet

```go
type Node struct {
	value int
	min   int // Local min
}

type Stack struct {
	data []Node
}

func (s *Stack) Push(val int) {
	newMin := val
	if len(s.data) > 0 {
		currentMin := s.data[len(s.data)-1].min
		if currentMin < newMin {
			newMin = currentMin
		}
	}
	s.data = append(s.data, Node{value: val, min: newMin})
}

func (s *Stack) Min() int {
	if len(s.data) == 0 {
		return -1 // Error
	}
	return s.data[len(s.data)-1].min
}
```

## Complexity Analysis

- **Time Complexity**: O(1).
- **Space Complexity**: O(N).
