package main

import (
	"fmt"
	"runtime"
	"time"
)

// SetOfStacks implements a data structure composed of several stacks and should create a new stack once the previous one exceeds capacity.
// SetOfStacks.push() and SetOfStacks.pop() should behave identically to a single stack.
//
// Examples:
// Capacity = 2
// Push(1), Push(2) -> Stack 1: [1, 2]
// Push(3) -> Stack 1: [1, 2], Stack 2: [3]
// Pop() -> Returns 3, Stack 2 removed
type SetOfStacks struct {
	capacity int
	// TODO: Add fields
}

func NewSetOfStacks(capacity int) *SetOfStacks {
	return &SetOfStacks{capacity: capacity}
}

func (s *SetOfStacks) Push(val int) {
	// TODO: Implement
}

func (s *SetOfStacks) Pop() int {
	// TODO: Implement
	return -1
}

// Expected Time Complexity: O(1) for push/pop
// Expected Space Complexity: O(N)
// Thresholds:
//   Time:   Low < 2µs,   Medium < 20µs,   High > 200µs
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
			"SetOfStacks Capacity 2",
			func() []int {
				set := NewSetOfStacks(2)
				set.Push(1)
				set.Push(2)
				set.Push(3)
				set.Push(4)
				set.Push(5)
				val1 := set.Pop()
				val2 := set.Pop()
				val3 := set.Pop()
				return []int{val1, val2, val3}
			},
			[]int{5, 4, 3},
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
