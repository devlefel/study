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
	stack := NewStackMin()
	stack.Push(5)
	stack.Push(6)
	stack.Push(3)
	stack.Push(7)
	
	// Min should be 3
	min1 := stack.Min()
	stack.Pop() // Pop 7
	min2 := stack.Min() // Still 3
	stack.Pop() // Pop 3
	min3 := stack.Min() // Now 5
	
	status := "FAIL"
	if min1 == 3 && min2 == 3 && min3 == 5 {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Mins: %d, %d, %d)\n", status, min1, min2, min3)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	largeStack := NewStackMin()
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	for i := 1000; i > 0; i-- {
		largeStack.Push(i) // Pushing decreasing values updates min often
	}
	largeStack.Min()
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Operations: 1000 Push + 1 Min\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
