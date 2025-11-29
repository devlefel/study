package main

import (
	"fmt"
	"runtime"
	"time"
)

// IsUnique determines if a string has all unique characters.
// Implement an algorithm to determine if a string has all unique characters.
// Challenge: What if you cannot use additional data structures?
//
// Examples:
// Input: "abcde" -> Output: true
// Input: "hello" -> Output: false
// Input: "apple" -> Output: false
func IsUnique(str string) bool {
	// TODO: Implement this function
	return false
}

// Expected Time Complexity: O(N)
// Expected Space Complexity: O(1) if charset is fixed, else O(min(c, n))
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB
func main() {
	// Test Cases
	testCases := []struct {
		input    string
		expected bool
	}{
		{"abcde", true},
		{"hello", false},
		{"apple", false},
		{"kite", true},
		{"padle", true},
	}

	for _, tc := range testCases {
		result := IsUnique(tc.input)
		status := "FAIL"
		if result == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: %s (Expected: %v, Got: %v)\n", status, tc.input, tc.expected, result)
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	largeInput := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	IsUnique(largeInput)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Input Length: %d\n", len(largeInput))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
