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
	// Test Cases
	testCases := []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5, 8},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		result := SearchInRotatedArray(tc.input, tc.target)
		
		status := "FAIL"
		if result == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: Target %d -> Index %d (Expected: %d)\n", status, tc.target, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
