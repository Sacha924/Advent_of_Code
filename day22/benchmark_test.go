package main

import (
	"strconv"
	"testing"
)

// Method 1: Using modulo 10 to get the last digit
func lastDigitModulo(num int) int {
	return num % 10
}

// Method 2: Converting to string and accessing the last character
func lastDigitString(num int) int {
	numStr := strconv.Itoa(num)                    // Convert to string
	lastChar := numStr[len(numStr)-1]              // Access the last character
	lastDigit, _ := strconv.Atoi(string(lastChar)) // Convert back to int
	return lastDigit
}

// Benchmark for modulo approach
func BenchmarkLastDigitModulo(b *testing.B) {
	num := 123456789
	for i := 0; i < b.N; i++ {
		_ = lastDigitModulo(num)
	}
}

// Benchmark for string conversion approach
func BenchmarkLastDigitString(b *testing.B) {
	num := 123456789
	for i := 0; i < b.N; i++ {
		_ = lastDigitString(num)
	}
}
