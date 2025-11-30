# Trees and Graphs

## Trees

A tree is a data structure composed of nodes.

- **Root**: The top node.
- **Child**: A node directly connected to another node when moving away from the root.
- **Leaf**: A node with no children.

### Binary Trees

Each node has at most two children.

- **Binary Search Tree (BST)**: For every node, all left descendants <= n < all right descendants.
- **Balanced Tree**: A tree where the height is $O(\log N)$. (e.g., AVL, Red-Black).
- **Complete Binary Tree**: Every level is fully filled except perhaps the last.
- **Full Binary Tree**: Every node has either 0 or 2 children.

### Tree Traversal

- **In-Order**: Left -> Current -> Right. (Sorted order for BST).
- **Pre-Order**: Current -> Left -> Right.
- **Post-Order**: Left -> Right -> Current.

## Graphs

A tree is actually a type of graph (connected, acyclic). A graph consists of nodes (vertices) and edges.

- **Directed vs Undirected**: One-way vs two-way streets.
- **Connected**: A path exists between every pair of vertices.
- **Acyclic**: No cycles.

### Graph Search

- **Depth-First Search (DFS)**: Explore deep before wide. Uses a stack (or recursion). Good for visiting every node.
- **Breadth-First Search (BFS)**: Explore neighbors before neighbors' neighbors. Uses a queue. Good for finding shortest path.

## Go Implementation

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type GraphNode struct {
    Val       int
    Neighbors []*GraphNode
}
```
