# Stacks and Queues

## Stacks

A stack is a **LIFO** (Last-In, First-Out) data structure.

- **Push**: Add an item to the top. $O(1)$.
- **Pop**: Remove the top item. $O(1)$.
- **Peek**: Return the top item without removing it. $O(1)$.
- **IsEmpty**: Return true if the stack is empty. $O(1)$.

**Uses**: Recursive algorithms (call stack), backtracking (maze solving), undo mechanisms.

## Queues

A queue is a **FIFO** (First-In, First-Out) data structure.

- **Add (Enqueue)**: Add an item to the end. $O(1)$.
- **Remove (Dequeue)**: Remove the first item. $O(1)$.
- **Peek**: Return the top item. $O(1)$.
- **IsEmpty**: Return true if the queue is empty. $O(1)$.

**Uses**: Breadth-First Search (BFS), print queues, processing tasks in order.

## Go Implementation

Go uses slices for stacks and queues, but for queues, removing from the front of a slice is $O(N)$ if you shift elements. A linked list or a circular buffer is better for a true $O(1)$ queue.

**Stack (Slice)**

```go
stack := []int{}
stack = append(stack, 1) // Push
val := stack[len(stack)-1] // Peek
stack = stack[:len(stack)-1] // Pop
```

**Queue (Slice - Simple but inefficient dequeue)**

```go
queue := []int{}
queue = append(queue, 1) // Enqueue
val := queue[0] // Peek
queue = queue[1:] // Dequeue (O(N) shift)
```
