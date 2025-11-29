package main

import (
	"fmt"
	"runtime"
	"time"
)

// SieveOfEratosthenes generates a list of booleans where index i is true if i is prime.
// Implement the Sieve of Eratosthenes to find all primes up to max.
//
// Examples:
// Input: 5 -> Output: [false, false, true, true, false, true] (0, 1, 2, 3, 4, 5)
func SieveOfEratosthenes(max int) []bool {
	// TODO: Implement this function
	return nil
}

// Expected Time Complexity: O(N log log N)
// Expected Space Complexity: O(N)
// Thresholds:
//   Time:   Low < 1ms,   Medium < 10ms,   High > 100ms
//   Memory: Low < 1MB,   Medium < 10MB,   High > 100MB

func main() {
	max := 10
	primes := SieveOfEratosthenes(max)
	
	// Primes up to 10: 2, 3, 5, 7
	expectedPrimes := []int{2, 3, 5, 7}
	
	status := "PASS"
	if primes == nil {
		status = "FAIL"
		fmt.Println("FAIL: Returned nil")
	} else {
		for _, p := range expectedPrimes {
			if !primes[p] {
				status = "FAIL"
				break
			}
		}
		if status == "PASS" {
			if primes[4] || primes[6] || primes[8] || primes[9] {
				status = "FAIL"
			}
		}
	}
	fmt.Printf("Test Case 1: %s (Primes up to 10)\n", status)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	// Find primes up to 1,000,000
	SieveOfEratosthenes(1000000)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Operation: Sieve up to 1,000,000\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
