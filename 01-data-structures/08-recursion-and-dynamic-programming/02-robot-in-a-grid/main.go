package main

import (
	"fmt"
	"runtime"
	"time"
)

type Point struct {
	Row, Col int
}

// RobotInAGrid finds a path for a robot from the top left (0,0) to the bottom right (r-1, c-1).
// The robot can only move down or right. Some cells are "off limits" (represented by true in the maze).
// Return a list of points representing the path.
//
// Examples:
// Input: 2x2 grid, no obstacles -> Output: [(0,0), (0,1), (1,1)] or [(0,0), (1,0), (1,1)]
func RobotInAGrid(maze [][]bool) []Point {
	// TODO: Implement this function
	return nil
}

// Expected Time Complexity: O(R*C)
// Expected Space Complexity: O(R*C)
// Thresholds:
//   Time:   Low < 10µs,   Medium < 100µs,   High > 1ms
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB

func main() {
	// 3x3 Grid
	// F F F
	// F T F  (Obstacle at 1,1)
	// F F F
	maze := [][]bool{
		{false, false, false},
		{false, true, false},
		{false, false, false},
	}
	
	path := RobotInAGrid(maze)
	
	status := "FAIL"
	if path != nil && len(path) > 0 && path[0] == (Point{0, 0}) && path[len(path)-1] == (Point{2, 2}) {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Path Length: %d)\n", status, len(path))

	// Profiling
	fmt.Println("\n--- Profiling ---")
	// Create 100x100 grid
	largeMaze := make([][]bool, 100)
	for i := range largeMaze {
		largeMaze[i] = make([]bool, 100)
	}
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	RobotInAGrid(largeMaze)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Grid: 100x100\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
