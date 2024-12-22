package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Ideas :
map of antenna type to all the corresponding positions of these antenna types (string to [2]int)
for each pair of an antenna type, compute the possible antinodes position within the bounds
add this antinode positions into a set and do the same for every antenna type.
*/

type computeAntiNodeFunc func([2]int, [2]int, int, int) [][2]int

func main() {
	antennaTypeToPos, rows, cols := readAndProcess("day8/input.txt")
	fmt.Printf("Day 8 result : %d\n", computeAllAntiNodes(antennaTypeToPos, rows, cols, computeAntiNode))
	fmt.Printf("Day 8 result Part 2 : %d\n", computeAllAntiNodes(antennaTypeToPos, rows, cols, computeAntiNodePart2))
}

func readAndProcess(filename string) (map[rune][][2]int, int, int) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	antennaTypeToPos := make(map[rune][][2]int)
	rows := 0
	cols := 0
	for scanner.Scan() {
		line := scanner.Text()
		if rows == 0 {
			cols = len(line) // just store once the length of the row for later (matrix dim)
		}
		for index, char := range line {
			if char != '.' {
				antennaTypeToPos[char] = append(antennaTypeToPos[char], [2]int{rows, index})
			}
		}
		rows++
	}
	return antennaTypeToPos, rows, cols
}

func computeAllAntiNodes(m map[rune][][2]int, rows, cols int, computeAntiNode computeAntiNodeFunc) int {
	antiNodePos := make(map[[2]int]bool, 0) // Set counting the distinct antinodes pos

	// for every antenna type
	// test every pair of nodes (so node 1 with all the nodes, node2 with all the nodes except node1, etc)
	for _, nodes := range m {
		for i := 0; i < len(nodes)-1; i++ {
			for j := i + 1; j < len(nodes); j++ {
				for _, confirmedAntinodePos := range computeAntiNode(nodes[i], nodes[j], rows, cols) {
					antiNodePos[confirmedAntinodePos] = true
				}
			}
		}
	}
	return len(antiNodePos)
}

func computeAntiNode(node1, node2 [2]int, rows, cols int) [][2]int {
	var antiNodePos [][2]int
	diffRow, diffCol := node2[0]-node1[0], node2[1]-node1[1]
	// antinodes are created at node1 + (diffRow, diffCol) and  node2 + (diffRow, diffCol) if they are in the matrix
	if node1[0]-diffRow >= 0 && node1[1]-diffCol >= 0 && node1[0]-diffRow < rows && node1[1]-diffCol < cols {
		antiNodePos = append(antiNodePos, [2]int{node1[0] - diffRow, node1[1] - diffCol})
	}
	if node2[0]+diffRow >= 0 && node2[1]+diffCol >= 0 && node2[0]+diffRow < rows && node2[1]+diffCol < cols {
		antiNodePos = append(antiNodePos, [2]int{node2[0] + diffRow, node2[1] + diffCol})
	}
	return antiNodePos
}

func computeAntiNodePart2(node1, node2 [2]int, rows, cols int) [][2]int {
	antiNodePos := [][2]int{node1, node2}
	diffRow, diffCol := node2[0]-node1[0], node2[1]-node1[1]
	// antinodes are created at node1 + (diffRow, diffCol) and  node2 + (diffRow, diffCol) if they are in the matrix
	x, y := node1[0], node1[1]
	for x-diffRow >= 0 && y-diffCol >= 0 && x-diffRow < rows && y-diffCol < cols {
		x -= diffRow
		y -= diffCol
		antiNodePos = append(antiNodePos, [2]int{x, y})
	}
	x, y = node2[0], node2[1]
	for x+diffRow >= 0 && y+diffCol >= 0 && x+diffRow < rows && y+diffCol < cols {
		x += diffRow
		y += diffCol
		antiNodePos = append(antiNodePos, [2]int{x, y})
	}
	return antiNodePos
}
