package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

var memo = make(map[string]int)

func main() {
	elems, pattern := parseInput("day19/input.txt") // elem is a stripe (containing one or few letters), pattern is a towel (containing many letters)
	res := 0
	resPart2 := 0
	for _, towel := range pattern {
		filteredElems := reduceElemsSet(towel, elems)
		if isTowelBuildable(elems, towel) {
			res++
		}
		resPart2 += getEveryBuildPerTowel(towel, filteredElems)
	}
	fmt.Println(res)
	fmt.Println(resPart2)
}

func parseInput(path string) ([]string, []string) {
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	elems := []string{}
	pattern := []string{}
	isElem := true
	for scanner.Scan() {
		if isElem {
			elems = append(elems, strings.Split(scanner.Text(), ", ")...)
			isElem = false
		} else {
			line := strings.Trim(strings.TrimSpace(scanner.Text()), "\n")
			if line == "" {
				continue
			}
			pattern = append(pattern, line)
		}
	}
	return elems, pattern
}

func isTowelBuildable(elems []string, pattern string) bool {
	var pq PriorityQueue
	pq.Push(
		&Item{
			subPattern: pattern,
			priority:   len(pattern),
			index:      0,
		},
	)
	heap.Init(&pq)
	for pq.Len() > 0 {
		item := pq.Pop().(*Item)
		for _, elem := range elems {
			if len(elem) > len(item.subPattern) {
				continue
			}
			if elem == item.subPattern[:len(elem)] {
				if len(elem) == len(item.subPattern) {
					return true // we made the towel
				}
				// else there is remaining stripes to made
				remainingStripes := item.subPattern[len(elem):]
				pq.Push(
					&Item{
						subPattern: remainingStripes,
						priority:   len(remainingStripes),
					})
			}
		}
	}
	return false
}

func getEveryBuildPerTowel(remainingStr string, elems []string) int {
	if val, ok := memo[remainingStr]; ok {
		return val
	}
	if len(remainingStr) == 0 {
		return 1
	}
	for _, elem := range elems {
		if len(elem) > len(remainingStr) {
			continue
		}
		if elem == remainingStr[:len(elem)] {
			memo[remainingStr] += getEveryBuildPerTowel(remainingStr[len(elem):], elems)
		}
	}
	return memo[remainingStr]
}

func reduceElemsSet(pattern string, elems []string) []string {
	// if a towel didn't have the sequence of char inside an elem, it's not usefull to keep this elem as it will not be used to make any combinations to build the towel
	var filteredElems []string
	for _, elem := range elems {
		if strings.Contains(pattern, elem) {
			filteredElems = append(filteredElems, elem)
		}
	}
	return filteredElems
}

//func getEveryBuildPerTowel(elems []string, pattern string) int { // TOO SLOW
//	// as we need to iterate over all towels, we need to reduce the set of valid towels
//	// also a priority queue is not useful, we can use a stack or a queue or whatever :)
//	sum := 0
//	q := make([]string, 50)
//	q = append(q, pattern)
//	filteredElems := reduceElemsSet(pattern, elems)
//	for len(q) > 0 {
//		subPattern := q[0]
//		q = q[1:]
//		for _, elem := range filteredElems {
//			if len(elem) > len(subPattern) {
//				continue
//			}
//			if elem == subPattern[:len(elem)] {
//				if len(elem) == len(subPattern) {
//					sum++
//				} else {
//					q = append(q, subPattern[len(elem):])
//				}
//			}
//		}
//	}
//	return sum
//}

//// Interesting bug : using dequeue function i did in "isTowerBuildable": subPattern, remainingStripes := dequeue(remainingStripes)
//// It leads to a bug due to the shadowing of remainingStripes, it was stuck on the same pattern
//func dequeue(queue []string) (string, []string) {
//	element := queue[0] // The first element is the one to be dequeued.
//	if len(queue) == 1 {
//		var tmp = []string{}
//		return element, tmp
//	}
//	return element, queue[1:] // Slice off the element once it is dequeued.
//}
