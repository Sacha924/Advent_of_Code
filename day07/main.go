package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type EqualCheckFunc func(int, []int) bool

func main() {
	start := time.Now()
	fmt.Printf("Day 7 result with canBeEqual: %d\n", readAndProcess("day07/input.txt", canBeEqual)) // 3351424677624
	fmt.Println(time.Since(start))                                                                  // 21ms, bitmask tech blazingly fast
	start = time.Now()
	fmt.Printf("Day 7 result with canBeEqual2: %d\n", readAndProcess("day07/input.txt", func(expected int, numbers []int) bool {
		return canBeEqual2(expected, numbers, numbers[0], 1)
	}))                            // 204976636995111%
	fmt.Println(time.Since(start)) // 1.1s not that bad
}

func readAndProcess(filename string, equalCheck EqualCheckFunc) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		valuesStr := strings.Split(line, " ")
		var expected int
		numbers := make([]int, 0)
		for i, v := range valuesStr {
			if i == 0 {
				expected, _ = strconv.Atoi(strings.Trim(v, ":"))
			}
			valueInt, _ := strconv.Atoi(v)
			numbers = append(numbers, valueInt)
		}
		if equalCheck(expected, numbers) {
			res += expected
		}
	}

	return res
}

// return 1 if we can find a pattern of operators such that the equation can be equal
// we have 2 ^ len(op) - 1 combination of operators
// we can represent the combination as 2^1 * 2^0 * 2^0 ... n-1 times, the bit is either 1 or 0
// Let's say 1 represent mult and 0 addition
func canBeEqual(expected int, numbers []int) bool {
	n := len(numbers)
	// Total combinations of operators is 2^(n-1)
	totalCombinations := 1 << (n - 1)

	for i := 0; i < totalCombinations; i++ {
		result := numbers[0]
		for j := 0; j < n-1; j++ {
			if (i>>j)&1 == 1 {
				// If the j-th bit is 1, multiply
				result *= numbers[j+1]
			} else {
				// If the j-th bit is 0, add
				result += numbers[j+1]
			}
		}
		if result == expected {
			return true
		}
	}
	return false
}

func canBeEqual2(expected int, numbers []int, current int, index int) bool {
	if index == len(numbers) {
		return current == expected
	}
	if canBeEqual2(expected, numbers, current+numbers[index], index+1) {
		return true
	}
	if canBeEqual2(expected, numbers, current*numbers[index], index+1) {
		return true
	}
	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, numbers[index]))
	if canBeEqual2(expected, numbers, concatenated, index+1) {
		return true
	}
	return false
}
