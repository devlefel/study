package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// StackMin is a stack that, in addition to push and pop, has a function min which returns the minimum element.
// Push, pop and min should all operate in O(1) time.
//
// Examples:
// Push(5), Min() -> 5
// Push(6), Min() -> 5
// Push(3), Min() -> 3
// Pop() -> 3, Min() -> 5
type StackMin struct {
	// TODO: Add fields
}

func NewStackMin() *StackMin {
	return &StackMin{}
}

func (s *StackMin) Push(val int) {
	// TODO: Implement
}

func (s *StackMin) Pop() int {
	// TODO: Implement
	return -1
}

func (s *StackMin) Min() int {
	// TODO: Implement
	return math.MaxInt64
}

// Expected Time Complexity: O(1) for push, pop, and min
// Expected Space Complexity: O(N)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB

func main() {
	// Test Cases
	type testCase struct {
		name     string
		ops      func() []int
		expected []int
	}

	testCases := []testCase{
		{
			"Push/Pop/Min Sequence",
			func() []int {
				stack := NewStackMin()
				stack.Push(5)
				stack.Push(6)
				stack.Push(3)
				stack.Push(7)
				m1 := stack.Min()
				stack.Pop()
				m2 := stack.Min()
				stack.Pop()
				m3 := stack.Min()
				return []int{m1, m2, m3}
			},
			[]int{3, 3, 5},
		},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		result := tc.ops()
		status := "FAIL"
		if len(result) == len(tc.expected) {
			match := true
			for i := range result {
				if result[i] != tc.expected[i] {
					match = false
					break
				}
			}
			if match {
				status = "PASS"
			}
		}
		fmt.Printf("%s: '%s' -> '%v' (Expected: '%v')\n", status, tc.name, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
