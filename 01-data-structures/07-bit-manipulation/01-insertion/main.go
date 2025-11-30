package main

import (
	"fmt"
	"runtime"
	"time"
)

// Insertion inserts M into N such that M starts at bit j and ends at bit i.
// You can assume that the bits j through i have enough space to fit all of M.
// That is, if M = 10011, you can assume that there are at least 5 bits between j and i.
// You would not, for example, have j = 3 and i = 2, because M could not fully fit between bit 3 and bit 2.
//
// Examples:
// Input: N = 10000000000, M = 10011, i = 2, j = 6
// Output: N = 10001001100
func Insertion(n int, m int, i int, j int) int {
	// TODO: Implement this function
	return 0
}

// Expected Time Complexity: O(1) (assuming 32 or 64 bit integers)
// Expected Space Complexity: O(1)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 1µs,   High > 1µs (Bitwise is extremely fast)
//   Memory: Low < 1KB,   Medium < 1KB,   High > 1KB

func main() {
	// Test Cases
	testCases := []struct {
		n, m, i, j int
		expected   int
	}{
		{1024, 19, 2, 6, 1100},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		result := Insertion(tc.n, tc.m, tc.i, tc.j)
		
		status := "FAIL"
		if result == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: N=%b, M=%b, i=%d, j=%d -> %b (Expected: %b)\n", status, tc.n, tc.m, tc.i, tc.j, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
