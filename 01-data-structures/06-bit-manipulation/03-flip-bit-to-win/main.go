package main

import (
	"fmt"
	"runtime"
	"time"
)

// FlipBitToWin finds the length of the longest sequence of 1s you could create by flipping exactly one bit from a 0 to a 1.
//
// Examples:
// Input: 1775 (11011101111) -> Output: 8
func FlipBitToWin(num int) int {
	// TODO: Implement this function
	return 0
}

// Expected Time Complexity: O(b) where b is bits in integer (32/64) -> O(1)
// Expected Space Complexity: O(1)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 5µs,   High > 10µs
//   Memory: Low < 1KB,   Medium < 1KB,   High > 1KB

func main() {
	// Test Cases
	// 1775 (11011101111) -> Flip 0 at index 3 -> 11011111111 -> Length 8
	num := 1775
	result := FlipBitToWin(num)
	expected := 8
	
	status := "FAIL"
	if result == expected {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Got %d, Expected %d)\n", status, result, expected)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	for k := 0; k < 1000000; k++ {
		FlipBitToWin(1775)
	}
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Operations: 1,000,000 Flips\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
