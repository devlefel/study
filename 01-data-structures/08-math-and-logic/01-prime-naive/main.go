package main

import (
	"fmt"
	"runtime"
	"time"
)

// IsPrime checks if a number is prime.
// Implement a function that checks if a number is prime.
//
// Examples:
// Input: 5 -> Output: true
// Input: 4 -> Output: false
// Input: 1 -> Output: false
func IsPrime(n int) bool {
	// TODO: Implement this function
	return false
}

// Expected Time Complexity: O(sqrt(N))
// Expected Space Complexity: O(1)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 1KB,   High > 1KB

func main() {
	// Test Cases
	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{17, true},
		{18, false},
		{1, false},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		result := IsPrime(tc.input)
		
		status := "FAIL"
		if result == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: %d -> %v (Expected: %v)\n", status, tc.input, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
