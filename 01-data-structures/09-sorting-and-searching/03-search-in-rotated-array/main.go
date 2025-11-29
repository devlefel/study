package main

import (
	"fmt"
	"runtime"
	"time"
)

// SearchInRotatedArray searches for a target in a sorted array that has been rotated.
//
// Examples:
// Input: [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], Target: 5 -> Output: 8
func SearchInRotatedArray(arr []int, target int) int {
	return search(arr, 0, len(arr)-1, target)
}

func search(a []int, left, right, x int) int {
	if right < left {
		return -1
	}
	
	mid := (left + right) / 2
	if a[mid] == x {
		return mid
	}
	
	if a[left] < a[mid] { // Left is normally ordered
		if x >= a[left] && x < a[mid] {
			return search(a, left, mid-1, x)
		} else {
			return search(a, mid+1, right, x)
		}
	} else if a[left] > a[mid] { // Right is normally ordered
		if x > a[mid] && x <= a[right] {
			return search(a, mid+1, right, x)
		} else {
			return search(a, left, mid-1, x)
		}
	} else { // Left == Mid (Duplicates)
		if a[mid] != a[right] { // If right is different, search it
			return search(a, mid+1, right, x)
		} else { // Else we have to search both
			result := search(a, left, mid-1, x)
			if result == -1 {
				return search(a, mid+1, right, x)
			}
			return result
		}
	}
}

// Expected Time Complexity: O(log N)
// Expected Space Complexity: O(log N) (Recursion) or O(1) (Iterative)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 1KB,   High > 1KB

func main() {
	arr := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	target := 5
	expected := 8
	
	result := SearchInRotatedArray(arr, target)
	
	status := "FAIL"
	if result == expected {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Index: %d)\n", status, result)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	// Create large rotated array
	count := 1000
	largeInput := make([]int, count)
	// Rotate at 500
	for i := 0; i < count; i++ {
		val := i + 500
		if val >= count {
			val -= count
		}
		largeInput[i] = val
	}
	// Target is 499 (which is at index 999)
	targetVal := 499
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	SearchInRotatedArray(largeInput, targetVal)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Input Size: 1000\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
