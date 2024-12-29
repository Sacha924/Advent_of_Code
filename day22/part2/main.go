package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// over the different buyers, we store the seq and the sum of bananas it can afford (if two sequences appear in the same buyer prices, we only add the best of the two)
var seqToBananas = map[string]int{}

func main() {
	numbers := parseInput("day22/input.txt")
	start := time.Now()
	for _, number := range numbers {
		SimulateBuyerPlay(number)
	}
	fmt.Printf("Time elapsed: %s\n", time.Since(start))
	res := 0
	var bestSeq string
	for k, v := range seqToBananas {
		if v > res {
			res = v
			bestSeq = k
		}
	}
	fmt.Printf("res %d, seqToBananas: %s\n", res, bestSeq)
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

func SimulateBuyerPlay(number int) {
	var seqToBananasBuyer = map[string]int{}
	currentSeq := make([]string, 0, 5)
	previousBananas := number % 10 // previous numbers of bananas according to the previous secretNumber%10
	var currentSeqString string

	for range 2000 {
		number = generateNextSecret(number)
		bananas := number % 10
		diffStr := strconv.Itoa(bananas - previousBananas)
		currentSeq = append(currentSeq, diffStr)
		if len(currentSeq) == 5 {
			currentSeq = currentSeq[1:]
		}
		currentSeqString = strings.Join(currentSeq, "")
		if len(currentSeq) == 4 && seqToBananasBuyer[currentSeqString] == 0 {
			seqToBananasBuyer[currentSeqString] = bananas
		}
		previousBananas = bananas

	}
	for k, v := range seqToBananasBuyer {
		seqToBananas[k] += v
	}
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
