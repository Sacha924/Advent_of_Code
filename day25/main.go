package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	locks, keys := buildLockAndKeys("day25/input.txt")
	fmt.Println("matchingPairs:", findMatchingPairs(locks, keys)) // 2854
}

func buildLockAndKeys(path string) (locks [][]int, keys [][]int) {
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var item [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			switch item[0][0] {
			case '#':
				locks = append(locks, buildItemArrNb(item))
			case '.':
				keys = append(keys, buildItemArrNb(item))
			}
			item = [][]rune{} // reset item
		} else {
			item = append(item, []rune(line))
		}
	}
	keys = append(keys, buildItemArrNb(item))
	return
}

// could be optimized by looping over column and stop when it reaches a dot
func buildItemArrNb(item [][]rune) []int {
	itemArrNb := make([]int, len(item[0]))
	for i := range itemArrNb {
		itemArrNb[i] = -1
	}
	for i := range item {
		for j := range item[i] {
			if item[i][j] == '#' {
				itemArrNb[j]++
			}
		}
	}
	return itemArrNb
}

func findMatchingPairs(locks [][]int, keys [][]int) int {
	count := 0

	for _, lock := range locks {
		for _, key := range keys {
			if isMatch(lock, key) {
				count++
			}
		}
	}
	return count
}

func isMatch(lock []int, key []int) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}
