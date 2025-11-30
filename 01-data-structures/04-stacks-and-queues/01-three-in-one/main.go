package main

import (
	"fmt"
	"runtime"
	"time"
)

// ThreeInOne describes how you could use a single array to implement three stacks.
// Implement a data structure that uses a single array to implement three stacks.
//
// Examples:
// stack.Push(0, 10) -> Stack 0 has [10]
// stack.Push(1, 20) -> Stack 1 has [20]
// stack.Push(0, 11) -> Stack 0 has [10, 11]
// stack.Pop(0) -> Returns 11, Stack 0 has [10]
type ThreeStacks struct {
	// TODO: Add fields
}

func NewThreeStacks(stackSize int) *ThreeStacks {
	// TODO: Implement
	return &ThreeStacks{}
}

func (ts *ThreeStacks) Push(stackNum int, value int) error {
	// TODO: Implement
	return nil
}

func (ts *ThreeStacks) Pop(stackNum int) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (ts *ThreeStacks) Peek(stackNum int) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (ts *ThreeStacks) IsEmpty(stackNum int) bool {
	// TODO: Implement
	return true
}

// Expected Time Complexity: O(1) for push/pop
// Expected Space Complexity: O(N) for the array
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB

func main() {
	// Test Cases
	type testCase struct {
		name     string
		ops      func() ([]int, error)
		expected []int
	}

	testCases := []testCase{
		{
			"Push/Pop Multi Stacks",
			func() ([]int, error) {
				stack := NewThreeStacks(3)
				stack.Push(0, 10)
				stack.Push(0, 11)
				stack.Push(1, 20)
				val1, _ := stack.Pop(0)
				val2, _ := stack.Pop(1)
				return []int{val1, val2}, nil
			},
			[]int{11, 20},
		},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		result, _ := tc.ops()
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
