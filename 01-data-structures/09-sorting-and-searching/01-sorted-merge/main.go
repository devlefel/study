package main

import (
	"fmt"
	"runtime"
	"time"
)

// SortedMerge merges two sorted arrays, A and B, where A has a large enough buffer at the end to hold B.
// Write a method to merge B into A in sorted order.
//
// Examples:
// Input: A=[1, 3, 5, 0, 0, 0], B=[2, 4, 6] -> Output: A=[1, 2, 3, 4, 5, 6]
func SortedMerge(a []int, b []int, countA int, countB int) {
	indexA := countA - 1
	indexB := countB - 1
	indexMerged := countA + countB - 1
	
	for indexB >= 0 {
		if indexA >= 0 && a[indexA] > b[indexB] {
			a[indexMerged] = a[indexA]
			indexA--
		} else {
			a[indexMerged] = b[indexB]
			indexB--
		}
		indexMerged--
	}
}

// Expected Time Complexity: O(A + B)
// Expected Space Complexity: O(1) (In-place)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 1KB,   High > 1KB

func main() {
	// A has buffer
	a := []int{1, 3, 5, 0, 0, 0}
	b := []int{2, 4, 6}
	
	SortedMerge(a, b, 3, 3)
	
	expected := []int{1, 2, 3, 4, 5, 6}
	match := true
	for i := 0; i < 6; i++ {
		if a[i] != expected[i] {
			match = false
			break
		}
	}
	
	status := "FAIL"
	if match {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Result: %v)\n", status, a)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	// Create large arrays
	count := 1000
	largeA := make([]int, count*2)
	largeB := make([]int, count)
	for i := 0; i < count; i++ {
		largeA[i] = i * 2
		largeB[i] = i*2 + 1
	}
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	SortedMerge(largeA, largeB, count, count)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Merge Size: 1000 + 1000\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
