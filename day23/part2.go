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
	adjencyList, nodeSet := builfAdjListAndNodeSet("day23/input.txt")
	fmt.Println(findLargestNetwork(adjencyList, nodeSet))
	fmt.Println("elapsed:", time.Since(start))
}

func builfAdjListAndNodeSet(filename string) (map[string][]string, sets.Set[string]) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	adjList := make(map[string][]string)
	var nodeSet sets.Set[string]

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		adjList[parts[0]] = append(adjList[parts[0]], parts[1])
		adjList[parts[1]] = append(adjList[parts[1]], parts[0])
		nodeSet.Add(parts[0])
		nodeSet.Add(parts[1])
	}
	return adjList, nodeSet
}

func findLargestNetwork(adjList map[string][]string, nodeSet sets.Set[string]) string {
	networks := make([][]string, 0, 10000)
	for nodeToInclude, _ := range nodeSet {
		for _, network := range networks { // for each network we check if we can add the node or not, each nodes in the network need to have the nodeToInclude in it's agency list
			canBeAdded := true
			for _, node := range network {
				if !find(adjList[node], nodeToInclude) {
					canBeAdded = false
					break
				}
			}
			if canBeAdded {
				newNetwork := append([]string{}, network...)
				newNetwork = append(newNetwork, nodeToInclude)
				networks = append(networks, newNetwork)

			}
		}
		networks = append(networks, []string{nodeToInclude})
	}
	largestArr := findLargestArray(networks)
	sort.Strings(largestArr)
	return strings.Join(largestArr, ",")
}

func find(arr []string, v string) bool {
	for _, str := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func findLargestArray(arrays [][]string) []string {
	largestArray := arrays[0]
	for _, array := range arrays {
		if len(array) > len(largestArray) { // Compare lengths
			largestArray = array // Update if a larger array is found
		}
	}
	return largestArray
}
