# Linked Lists

A linked list is a data structure that represents a sequence of nodes. In a singly linked list, each node points to the next node. In a doubly linked list, each node points to both the next and the previous node.

## Key Properties

- **Dynamic Size**: Unlike arrays, linked lists can grow and shrink dynamically.
- **Efficient Insertion/Deletion**: Inserting or deleting a node is $O(1)$ if you have a reference to the node (and its predecessor for singly linked lists).
- **Sequential Access**: Accessing the $k$-th element takes $O(k)$ time.

## Go Implementation

Go doesn't have a built-in Linked List class like Java's `LinkedList`, but it has the `container/list` package for doubly linked lists. However, in interviews, you are usually expected to define your own Node struct.

```go
type Node struct {
    Data int
    Next *Node
}
```

## The "Runner" Technique (Two Pointers)

A common technique is to use two pointers iterating at different speeds.

- **Finding the Middle**: Fast pointer moves 2 steps, slow pointer moves 1 step. When fast reaches the end, slow is at the middle.
- **Cycle Detection**: If fast and slow collide, there is a cycle.

## Recursive Problems

Many linked list problems can be solved recursively.

- $O(N)$ space for the stack.
- Useful for printing in reverse or processing from the end.
