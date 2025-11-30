# Answer: Three in One

## Problem

Describe how you could use a single array to implement three stacks.

## Solution Approach

1.  **Fixed Division**: Divide the array into three equal parts.
    - Stack 1: `[0, n/3)`
    - Stack 2: `[n/3, 2n/3)`
    - Stack 3: `[2n/3, n)`
    - _Pros_: Simple.
    - _Cons_: Inefficient if one stack is full and others are empty.
2.  **Flexible Division**: Store data and "previous index" in the array, or shift elements when a stack grows. (Complex).
    - We will implement **Fixed Division** for simplicity.

### Diagram

```text
Array: [ S1, S1, S1, S2, S2, S2, S3, S3, S3 ]
Indices: 0   1   2   3   4   5   6   7   8

Stack 1 Top: 1 (Value at index 1)
Stack 2 Top: 3 (Value at index 3)
Stack 3 Top: -1 (Empty)
```

## Code Snippet

```go
type FixedMultiStack struct {
	numberOfStacks int
	stackCapacity  int
	values         []int
	sizes          []int
}

func NewFixedMultiStack(stackSize int) *FixedMultiStack {
	return &FixedMultiStack{
		numberOfStacks: 3,
		stackCapacity:  stackSize,
		values:         make([]int, stackSize*3),
		sizes:          make([]int, 3),
	}
}

func (s *FixedMultiStack) Push(stackNum int, value int) {
	if s.IsFull(stackNum) {
		return // Error handling needed
	}
	s.sizes[stackNum]++
	s.values[s.indexOfTop(stackNum)] = value
}

func (s *FixedMultiStack) indexOfTop(stackNum int) int {
	offset := stackNum * s.stackCapacity
	size := s.sizes[stackNum]
	return offset + size - 1
}
```

## Complexity Analysis

- **Time Complexity**: O(1) for Push/Pop.
- **Space Complexity**: O(N) for the array.
