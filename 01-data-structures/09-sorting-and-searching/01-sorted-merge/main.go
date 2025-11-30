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
	// Test Cases
	type testCase struct {
		name     string
		setup    func() ([]int, []int, int, int)
		expected []int
	}

	testCases := []testCase{
		{
			"Merge with Buffer",
			func() ([]int, []int, int, int) {
				a := []int{1, 3, 5, 0, 0, 0}
				b := []int{2, 4, 6}
				return a, b, 3, 3
			},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		a, b, countA, countB := tc.setup()
		SortedMerge(a, b, countA, countB)
		
		status := "FAIL"
		match := true
		if len(a) != len(tc.expected) {
			match = false
		} else {
			for i := range a {
				if a[i] != tc.expected[i] {
					match = false
					break
				}
			}
		}
		
		if match {
			status = "PASS"
		}
		fmt.Printf("%s: '%s' -> '%v' (Expected: '%v')\n", status, tc.name, a, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
