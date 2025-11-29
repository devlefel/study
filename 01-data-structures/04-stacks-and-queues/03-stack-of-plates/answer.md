# Answer: Stack of Plates

## Problem

Imagine a (literal) stack of plates. If the stack gets too high, it might topple. Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold. Implement a data structure `SetOfStacks` that mimics this. `push()` and `pop()` should behave identically to a single stack (that is, `pop()` should return the same values as it would if there were just a single stack).

## Solution Approach

We need a list of stacks.

- `Push`: Check if the last stack is full.
  - If full, create a new stack and push to it.
  - If not full, push to the last stack.
- `Pop`: Check if the last stack is empty.
  - If empty, remove it and pop from the _new_ last stack.
  - If not empty, pop from it.

### Diagram

```text
Capacity = 2

Push 1, 2:
Stack 1: [1, 2] (Full)

Push 3:
Stack 1: [1, 2]
Stack 2: [3]

Push 4, 5:
Stack 1: [1, 2]
Stack 2: [3, 4] (Full)
Stack 3: [5]

Pop(): Returns 5. Stack 3 empty -> Delete Stack 3.
Pop(): Returns 4.
```

## Code Snippet

```go
type SetOfStacks struct {
	stacks     []*Stack
	capacity   int
}

func (s *SetOfStacks) Push(val int) {
	last := s.getLastStack()
	if last != nil && !last.IsFull() {
		last.Push(val)
	} else {
		newStack := NewStack(s.capacity)
		newStack.Push(val)
		s.stacks = append(s.stacks, newStack)
	}
}

func (s *SetOfStacks) Pop() int {
	last := s.getLastStack()
	if last == nil {
		return -1
	}
	val := last.Pop()
	if last.IsEmpty() {
		s.stacks = s.stacks[:len(s.stacks)-1] // Remove empty stack
	}
	return val
}
```

## Complexity Analysis

- **Time Complexity**: O(1).
- **Space Complexity**: O(N).
