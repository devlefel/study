package main

import (
	"fmt"
	"runtime"
	"time"
)

// MagicIndex finds a magic index in a sorted array. A magic index is an index such that A[i] = i.
// Given a sorted array of distinct integers, write a method to find a magic index, if one exists.
//
// Examples:
// Input: [-10, -5, 2, 20] -> Output: 2 (A[2] = 2)
// Input: [0, 5, 10] -> Output: 0 (A[0] = 0)
// Input: [10, 11, 12] -> Output: -1 (None)
func MagicIndex(arr []int) int {
	// TODO: Implement this function (Hint: Binary Search)
	return -1
}

// Expected Time Complexity: O(log N)
// Expected Space Complexity: O(log N) (Recursion)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 1KB,   High > 1KB

func main() {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{-10, -5, 2, 20}, 2},
		{[]int{0, 5, 10}, 0},
		{[]int{10, 11, 12}, -1},
	}

	for _, test := range tests {
		result := MagicIndex(test.input)
		if result == test.expected {
			fmt.Printf("PASS: MagicIndex(%v) = %d\n", test.input, result)
		} else {
			fmt.Printf("FAIL: MagicIndex(%v) = %d, expected %d\n", test.input, result, test.expected)
		}
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	// Create array where A[i] = i only at 999
	largeInput := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		largeInput[i] = i - 1 // A[i] = i-1, so A[i] < i
	}
	largeInput[999] = 999 // Magic index
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	MagicIndex(largeInput)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Input Size: 1000\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
