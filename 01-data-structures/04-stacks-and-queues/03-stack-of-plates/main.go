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
	// Capacity 2
	set := NewSetOfStacks(2)
	
	set.Push(1)
	set.Push(2)
	set.Push(3) // Should start new stack
	set.Push(4)
	set.Push(5) // Should start 3rd stack
	
	// Pop should return 5
	val1 := set.Pop()
	// Pop should return 4
	val2 := set.Pop()
	// Pop should return 3 (from 2nd stack)
	val3 := set.Pop()
	
	status := "FAIL"
	if val1 == 5 && val2 == 4 && val3 == 3 {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Got %d, %d, %d)\n", status, val1, val2, val3)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	largeSet := NewSetOfStacks(10) // Small capacity to force many stacks
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	for i := 0; i < 1000; i++ {
		largeSet.Push(i)
	}
	for i := 0; i < 1000; i++ {
		largeSet.Pop()
	}
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Operations: 1000 Push + 1000 Pop\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
