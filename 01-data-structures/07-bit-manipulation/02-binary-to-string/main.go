package main

import (
	"fmt"
	"runtime"
	"time"
)

// BinaryToString prints the binary representation of a real number between 0 and 1 (e.g., 0.72).
// If the number cannot be represented accurately in binary with at most 32 characters, print "ERROR".
//
// Examples:
// Input: 0.625 -> Output: "0.101"
// Input: 0.1 -> Output: "ERROR" (infinite repeating)
func BinaryToString(num float64) string {
	// TODO: Implement this function
	return ""
}

// Expected Time Complexity: O(N) (where N is the number of bits, max 32)
// Expected Space Complexity: O(N)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 5µs,   High > 10µs
//   Memory: Low < 1KB,   Medium < 1KB,   High > 1KB

func main() {
	// Test Cases
	testCases := []struct {
		input    float64
		expected string
	}{
		{0.75, "0.11"},
		{0.625, "0.101"},
		{0.72, "ERROR"},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		result := BinaryToString(tc.input)
		
		status := "FAIL"
		if result == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: %v -> '%s' (Expected: '%s')\n", status, tc.input, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
