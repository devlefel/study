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
	tests := []struct {
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

	for _, test := range tests {
		result := IsPrime(test.input)
		if result == test.expected {
			fmt.Printf("PASS: IsPrime(%d) = %v\n", test.input, result)
		} else {
			fmt.Printf("FAIL: IsPrime(%d) = %v, expected %v\n", test.input, result, test.expected)
		}
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	// Check a large prime
	IsPrime(1000000007)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Operation: IsPrime(1000000007)\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
