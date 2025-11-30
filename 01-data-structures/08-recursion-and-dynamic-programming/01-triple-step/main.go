package main

import (
	"fmt"
	"runtime"
	"time"
)

// TripleStep counts the number of ways a child can run up n steps, hopping 1, 2, or 3 steps at a time.
// Implement a method to count how many possible ways the child can run up the stairs.
//
// Examples:
// Input: 1 -> Output: 1 (1)
// Input: 2 -> Output: 2 (1+1, 2)
// Input: 3 -> Output: 4 (1+1+1, 1+2, 2+1, 3)
func TripleStep(n int) int {
	// TODO: Implement this function (Hint: Use Memoization)
	return 0
}

// Expected Time Complexity: O(N) (with memoization)
// Expected Space Complexity: O(N)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB

func main() {
	// Test Cases
	testCases := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 7}, // 1+3, 3+1, 2+2, 2+1+1, 1+2+1, 1+1+2, 1+1+1+1
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		result := TripleStep(tc.input)
		
		status := "FAIL"
		if result == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: %d -> %d (Expected: %d)\n", status, tc.input, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
