package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

/*
My main Idea for this challenge :
I need to keep track of the number I saw, and I need to make a mapping of the numbers TO the number that should not be before
It's important to consider in this "75|53", 53 is not "a number that should be after 75, but more "a number that shoud not be before"
Because numbers can just be absent of the list of updates.
EveryTime I find a number X, I need to check that any number already saw is not in the list associated with the numbers that should not be before X
*/

func main() {
	fmt.Printf("result of day5 : %d \n", day5("day5/input.txt", false))       // 7024
	fmt.Printf("result of day5 part 2 : %d \n", day5("day5/input.txt", true)) //
}

// rulesMapping give for a number, the list of number that should not be before
func day5(path string, isPart2 bool) int {
	f, _ := os.Open(path)
	defer f.Close()
	rulesMapping := make(map[int][]int)
	res := 0
	scanner := bufio.NewScanner(f)

	// first phase of scan is scanning rules, in the form of xx|yy
	// second phase of scan is scanning a list of word separated by spaces, we can switch to a more appropriate scanning function
	isScanningRules := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if isScanningRules && line == "" {
			isScanningRules = false
			continue
		}
		if isScanningRules {
			values := strings.Split(line, "|")
			if len(values) == 2 {
				v0, _ := strconv.Atoi(values[0])
				v1, _ := strconv.Atoi(values[1])
				rulesMapping[v0] = append(rulesMapping[v0], v1)
			}
		} else {
			values := strings.Split(line, ",")
			seen := make([]int, 0)
			isValidSequence := true
			for _, v := range values {
				v, _ := strconv.Atoi(v)
				shouldNotBeFoundInSeen := rulesMapping[v]
				for _, seenValue := range seen {
					if slices.Contains(shouldNotBeFoundInSeen, seenValue) {
						isValidSequence = false
						break
					}
				}
				if !isValidSequence {
					break
				}
				seen = append(seen, v)
			}
			if !isPart2 && isValidSequence {
				middleValue, _ := strconv.Atoi(values[len(values)/2])
				res += middleValue
			}
			if isPart2 && !isValidSequence {
				valuesInt := make([]int, len(values))
				for i, v := range values {
					val, _ := strconv.Atoi(strings.TrimSpace(v))
					valuesInt[i] = val
				}
				sortedValues, err := topologicalSort(valuesInt, rulesMapping)
				if err != nil {
					fmt.Printf("Error sorting update: %v\n", err)
					continue
				}
				middleValue := sortedValues[len(sortedValues)/2]
				res += middleValue
			}
		}
	}
	return res
}

// topologicalSort sorts the given sequence of pages according to the rules
func topologicalSort(sequence []int, rulesMapping map[int][]int) ([]int, error) {
	// Build graph
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	nodes := make(map[int]bool)

	// Initialize inDegree and nodes
	for _, node := range sequence {
		nodes[node] = true
		inDegree[node] = 0
	}

	// Build edges and compute in-degrees
	for after, befores := range rulesMapping {
		if !nodes[after] {
			continue
		}
		for _, before := range befores {
			if !nodes[before] {
				continue
			}
			graph[before] = append(graph[before], after)
			inDegree[after]++
		}
	}

	// Kahn's Algorithm
	queue := []int{}
	for node, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, node)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(sorted) != len(nodes) {
		return nil, fmt.Errorf("cycle detected in the graph")
	}

	return sorted, nil
}

// Let's have fun and use the famous bogosort !!!! spoiler : it didn't work it takes too much time
func transformIntoValidSeq(values []string, rulesMapping map[int][]int) {
	for {
		shuffle(values)
		// same code as in day 5 before to find if a sequence is valid
		seen := make([]int, 0)
		isValidSequence := true
		for _, v := range values {
			v, _ := strconv.Atoi(v)
			shouldNotBeFoundInSeen := rulesMapping[v]
			for _, seenValue := range seen {
				if slices.Contains(shouldNotBeFoundInSeen, seenValue) {
					isValidSequence = false
					break
				}
			}
			if !isValidSequence {
				break
			}
			seen = append(seen, v)
		}
		if isValidSequence {
			break
		}
	}
}

func shuffle(arr []string) {
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		j := rand.Intn(len(arr))
		arr[i], arr[j] = arr[j], arr[i] // Swap elements randomly
	}
}
