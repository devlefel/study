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
	// Test Cases
	type testCase struct {
		name     string
		setup    func() [][]bool
		verifier func([]Point) bool
	}

	testCases := []testCase{
		{
			"3x3 Grid with Obstacle",
			func() [][]bool {
				return [][]bool{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				}
			},
			func(path []Point) bool {
				if path != nil && len(path) > 0 && path[0] == (Point{0, 0}) && path[len(path)-1] == (Point{2, 2}) {
					return true
				}
				return false
			},
		},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		maze := tc.setup()
		path := RobotInAGrid(maze)
		
		status := "FAIL"
		if tc.verifier(path) {
			status = "PASS"
		}
		fmt.Printf("%s: '%s' -> Path Length %d\n", status, tc.name, len(path))
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
