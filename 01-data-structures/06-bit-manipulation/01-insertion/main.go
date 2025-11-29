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
	// N = 10000000000 (1024), M = 10011 (19), i = 2, j = 6
	// Expected: 10001001100 (1100)
	N := 1024
	M := 19
	i := 2
	j := 6
	
	result := Insertion(N, M, i, j)
	expected := 1100
	
	status := "FAIL"
	if result == expected {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Got %b, Expected %b)\n", status, result, expected)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	for k := 0; k < 1000000; k++ {
		Insertion(N, M, i, j)
	}
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Operations: 1,000,000 Insertions\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
