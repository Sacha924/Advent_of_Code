package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	numbers := parseInput("day22/input.txt")
	res := 0
	start := time.Now()
	for _, number := range numbers {
		res += generate2000th(number)
	}
	fmt.Printf("Time elapsed: %s\n", time.Since(start))
	fmt.Println(res)
}

func parseInput(filename string) []int {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	numbers := make([]int, 0)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	return numbers
}

func generate2000th(number int) int {
	for range 2000 {
		number = generateNextSecret(number)
	}
	return number
}

func generateNextSecret(number int) int {
	// step 1
	tmp := number * 64
	number = mix(number, tmp)
	number = prune(number)

	// step 2
	tmp = number / 32
	number = mix(number, tmp)
	number = prune(number)

	// step 3
	tmp = number * 2048
	number = mix(number, tmp)
	number = prune(number)

	return number
}

func mix(number int, tmp int) int {
	return number ^ tmp
}

func prune(number int) int {
	return number % 16777216
}
