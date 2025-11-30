# Answer: Robot in a Grid

## Problem

Imagine a robot sitting on the upper left corner of a grid with r rows and c columns. The robot can only move in two directions, right and down, but certain cells are "off limits" such that the robot cannot step on them. Design an algorithm to find a path for the robot from the top left to the bottom right.

## Solution Approach

This is similar to finding a path in a maze.
We start at (0,0) and try to reach (r-1, c-1).
Or, we can think backwards: to reach (r-1, c-1), we must have come from (r-2, c-1) or (r-1, c-2).
We can use **DFS (Depth First Search)** with **Memoization** (to avoid re-visiting failed paths).

- If `at destination`, return true.
- If `off limits` or `already visited`, return false.
- If `path exists from right neighbor` OR `path exists from down neighbor`, add current cell to path and return true.
- Else, mark as visited (failed) and return false.

### Diagram

```text
S . .
. X .
. . E

Path: (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2)
```

## Code Snippet

```go
func RobotInAGrid(maze [][]bool) []Point {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return nil
	}
	path := []Point{}
	failedPoints := make(map[Point]bool)
	if getPath(maze, len(maze)-1, len(maze[0])-1, &path, failedPoints) {
		return path
	}
	return nil
}

func getPath(maze [][]bool, row, col int, path *[]Point, failedPoints map[Point]bool) bool {
	if col < 0 || row < 0 || maze[row][col] { // maze[row][col] == true means obstacle
		return false
	}

	p := Point{row, col}
	if failedPoints[p] {
		return false
	}

	isAtOrigin := (row == 0) && (col == 0)

	if isAtOrigin || getPath(maze, row, col-1, path, failedPoints) || getPath(maze, row-1, col, path, failedPoints) {
		*path = append(*path, p)
		return true
	}

	failedPoints[p] = true
	return false
}
```

## Complexity Analysis

- **Time Complexity**: O(R\*C).
- **Space Complexity**: O(R\*C).
