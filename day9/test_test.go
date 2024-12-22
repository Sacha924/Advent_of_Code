package main

import (
	"fmt"
	"testing"
	"time"
)

func TestLoopEfficiency(t *testing.T) {
	// 100 char
	longString := "azertyuiop"
	i := 0
	for i < 28 {
		longString += longString
		i++
	}
	fmt.Println(len(longString))
	// 10M char in longstring

	start := time.Now()
	for a := 0; a < len(longString); a++ {

	}
	fmt.Println("for loop", time.Since(start))
	start = time.Now()
	for range longString {

	}
	fmt.Println("range loop", time.Since(start))
}
