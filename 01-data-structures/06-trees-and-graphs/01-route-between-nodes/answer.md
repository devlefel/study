# Answer: Route Between Nodes

## Problem

Given a directed graph and two nodes (S and E), design an algorithm to find out whether there is a route from S to E.

## Solution Approach

This is a classic Graph Traversal problem.

1.  **BFS (Breadth-First Search)**: Best for finding the shortest path. Explores neighbors first.
2.  **DFS (Depth-First Search)**: Simpler to implement (recursion), but might go deep down a wrong path.

We use **BFS** because it finds the shortest path and avoids getting stuck in infinite loops (if cycles exist) more easily (though both need a `visited` set).

### Diagram

```text
Graph:
A -> B -> C
|
v
D -> E

Route A -> C?
Queue: [A]
Pop A. Neighbors: B, D. Queue: [B, D]
Pop B. Neighbors: C. Found C! Return True.

Route D -> A?
Queue: [D]
Pop D. Neighbors: E. Queue: [E]
Pop E. Neighbors: None. Queue: []
Empty -> Return False.
```

## Code Snippet

```go
func RouteBetweenNodes(g *Graph, start, end int) bool {
	if start == end {
		return true
	}

	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // Dequeue

		for _, neighbor := range g.Nodes[node] {
			if neighbor == end {
				return true
			}
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return false
}
```

## Complexity Analysis

- **Time Complexity**: O(V + E), where V is vertices and E is edges.
- **Space Complexity**: O(V) for the queue and visited set.
