# Answer: List of Depths

## Problem

Given a binary tree, design an algorithm which creates a linked list of all the nodes at each depth (e.g., if you have a tree with depth D, you'll have D linked lists).

## Solution Approach

We can use **BFS (Breadth-First Search)** or **DFS (Depth-First Search)**.

1.  **BFS (Level Order Traversal)**:
    - Start with `currentLevel` list containing root.
    - While `currentLevel` is not empty:
      - Add `currentLevel` to results.
      - Create `parents` = `currentLevel`.
      - Create `currentLevel` = new list.
      - For each parent, add left and right children to `currentLevel`.
2.  **DFS (Pre-order)**:
    - Pass `level` + 1 to recursive calls.
    - If `results.size() == level`, create new list.
    - Add node to `results[level]`.

### Diagram

```text
      1
    /   \
   2     3
  /
 4

Level 0: [1]
Level 1: [2, 3]
Level 2: [4]

Result: List<LinkedList> = { {1}, {2, 3}, {4} }
```

## Code Snippet

```go
func ListOfDepths(root *Node) []*list.List {
	var result []*list.List
	if root == nil {
		return result
	}

	current := list.New()
	current.PushBack(root)

	for current.Len() > 0 {
		result = append(result, current) // Add previous level
		parents := current
		current = list.New()

		for e := parents.Front(); e != nil; e = e.Next() {
			node := e.Value.(*Node)
			if node.Left != nil {
				current.PushBack(node.Left)
			}
			if node.Right != nil {
				current.PushBack(node.Right)
			}
		}
	}
	return result
}
```

## Complexity Analysis

- **Time Complexity**: O(N).
- **Space Complexity**: O(N) (to store the lists).
