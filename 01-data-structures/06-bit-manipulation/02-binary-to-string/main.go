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
	// 0.75 -> "0.11"
	// 0.625 -> "0.101"
	// 0.72 -> "ERROR" (cannot be represented accurately in 32 bits)
	
	res1 := BinaryToString(0.75)
	res2 := BinaryToString(0.625)
	res3 := BinaryToString(0.72)
	
	status := "FAIL"
	if res1 == "0.11" && res2 == "0.101" && res3 == "ERROR" {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (0.75->%s, 0.625->%s, 0.72->%s)\n", status, res1, res2, res3)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	for k := 0; k < 100000; k++ {
		BinaryToString(0.125)
	}
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Operations: 100,000 Conversions\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
