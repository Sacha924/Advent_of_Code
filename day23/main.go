package main

import (
	"advent_of_code/pkg/sets"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	adjencyList := builfAdjList("day23/input.txt")
	fmt.Println(CountSetSize3(adjencyList, 't'))
	fmt.Println("elapsed:", time.Since(start))
}

func builfAdjList(filename string) map[string][]string {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	adjList := make(map[string][]string)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		adjList[parts[0]] = append(adjList[parts[0]], parts[1])
		adjList[parts[1]] = append(adjList[parts[1]], parts[0])
	}
	return adjList
}

func CountSetSize3(adjList map[string][]string, letter uint8) int {
	set := sets.Set[string]{}
	for curr, connectedToCurr := range adjList {
		if curr[0] == letter { // example : curr is 'ta'
			valueSet := sets.Set[string]{}
			for _, val := range connectedToCurr {
				valueSet.Add(val)
			}
			for _, val := range connectedToCurr { // looking at the connected node to 'ta'
				for _, val2 := range adjList[val] {
					if valueSet.Contains(val2) {
						size3 := []string{curr, val, val2}
						sort.Strings(size3)
						set.Add(strings.Join(size3, ""))
					}
				}
			}
		}
	}
	return len(set)
}
