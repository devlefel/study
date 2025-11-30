package main

import (
	"fmt"
	"runtime"
	"time"
)

// URLify replaces all spaces in a string with '%20'.
// Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.
//
// Examples:
// Input: "Mr John Smith    ", 13 -> Output: "Mr%20John%20Smith"
// Input: "Hello World  ", 11 -> Output: "Hello%20World"
// Input: "a b c    ", 5 -> Output: "a%20b%20c"
//
// Note: In Go, strings are immutable, so we'll use a rune slice to simulate in-place modification if desired,
// or just return a new string. For this exercise, returning a new string is fine, but try to do it efficiently.
func URLify(str string, trueLength int) string {
	// TODO: Implement this function
	return ""
}

// Expected Time Complexity: O(N)
// Expected Space Complexity: O(N) (if creating new string) or O(1) (in-place)
// Thresholds:
//   Time:   Low < 2µs,   Medium < 20µs,   High > 200µs
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB

func main() {
	// Test Cases
	testCases := []struct {
		input    string
		length   int
		expected string
	}{
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
		{"Hello World  ", 11, "Hello%20World"},
		{"   ", 1, "%20"},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	for _, tc := range testCases {
		// Convert string to rune slice for in-place modification simulation if needed
		// But for Go strings are immutable, so we usually return a new string
		result := URLify(tc.input, tc.length)
		status := "FAIL"
		if result == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: '%s' -> '%s' (Expected: '%s')\n", status, tc.input, result, tc.expected)
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	largeInput := "Mr John Smith                                     "
	trueLength := 13
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	URLify(largeInput, trueLength)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
```
