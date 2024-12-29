package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	matrix, x, y, dirX, dirY := readFile("day06/input.txt")
	visited := make(map[string]bool)
	makeGuardMove(matrix, x, y, dirX, dirY, visited)
	fmt.Printf("Day 6 part 2: %d\n", day6Part2(matrix, x, y, dirX, dirY)) // 1729 IT TAKES 4 SEC
}

func readFile(filename string) ([][]int, int, int, int, int) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, 0, 0, 0, 0
	}
	defer f.Close()
	matrix := make([][]int, 0)
	var guardX, guardY int
	var dirX, dirY int = -1, 0 // Initial direction is up
	scanner := bufio.NewScanner(f)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		row := make([]int, 0)
		for col, r := range line {
			switch r {
			case '.':
				row = append(row, 0)
			case '#':
				row = append(row, 1)
			case '^':
				row = append(row, 0)
				guardX = rowCount
				guardY = col
				dirX, dirY = -1, 0 // Facing up
			default:
				panic("Invalid character in input")
			}
		}
		matrix = append(matrix, row)
		rowCount++
	}
	return matrix, guardX, guardY, dirX, dirY
}

func makeGuardMove(matrix [][]int, x, y, dirX, dirY int, visited map[string]bool) int {
	for {
		key := fmt.Sprintf("%d,%d,%d,%d", x, y, dirX, dirY)
		if visited[key] {
			// Cycle detected; stop to avoid infinite loop
			return 1
		}
		visited[key] = true

		// Mark the current position as visited
		matrix[x][y] = 2

		// Calculate next position
		nextX := x + dirX
		nextY := y + dirY

		// Check if next position is within bounds
		if nextX < 0 || nextX >= len(matrix) || nextY < 0 || nextY >= len(matrix[0]) {
			// Guard leaves the mapped area
			return 0
		}

		if matrix[nextX][nextY] == 1 {
			// Obstacle detected; turn right 90 degrees
			dirX, dirY = rotateRight(dirX, dirY)
			continue
		} else {
			// Move forward
			x = nextX
			y = nextY
		}
	}
}

func rotateRight(dirX, dirY int) (int, int) {
	// Rotate direction vector 90 degrees to the right
	if dirX == -1 && dirY == 0 { // Up -> Right
		return 0, 1
	} else if dirX == 0 && dirY == 1 { // Right -> Down
		return 1, 0
	} else if dirX == 1 && dirY == 0 { // Down -> Left
		return 0, -1
	} else if dirX == 0 && dirY == -1 { // Left -> Up
		return -1, 0
	} else {
		panic("Invalid direction")
	}
}

// // Idea for part 2
// // Running the code once takes 20ms
// // Every position where I can try to put a wall to get a loop is all the position where the guard walks on (aka. the 4977 cases)
// // I can try to brute force and expect the code to take at most 5k * 10ms = 50s
func getCandidateWallPosition(matrix [][]int, guardStartX, guardStartY int) [][2]int {
	res := make([][2]int, 0)
	for i, row := range matrix {
		for j, cell := range row {
			if cell == 2 {
				res = append(res, [2]int{i, j})
			}
		}
	}
	return res
}

func day6Part2(matrix [][]int, guardStartX, guardStartY, dirX, dirY int) int {
	candidates := getCandidateWallPosition(matrix, guardStartX, guardStartY)
	res := 0
	for _, candidate := range candidates {
		wallX, wallY := candidate[0], candidate[1]
		matrix[wallX][wallY] = 1
		visited := make(map[string]bool)
		res += makeGuardMoveAndDetectCycle(matrix, guardStartX, guardStartY, dirX, dirY, visited)
		matrix[wallX][wallY] = 0 // remove the wall
	}
	return res
}

func makeGuardMoveAndDetectCycle(matrix [][]int, x, y, dirX, dirY int, visited map[string]bool) int {
	for {
		key := fmt.Sprintf("%d,%d,%d,%d", x, y, dirX, dirY)
		if visited[key] {
			// Cycle detected; stop to avoid infinite loop
			return 1
		}
		visited[key] = true
		nextX := x + dirX
		nextY := y + dirY

		// Check if next position is within bounds
		if nextX < 0 || nextX >= len(matrix) || nextY < 0 || nextY >= len(matrix[0]) {
			// Guard leaves the mapped area
			return 0
		}

		if matrix[nextX][nextY] == 1 {
			// Obstacle detected; turn right 90 degrees
			dirX, dirY = rotateRight(dirX, dirY)
			continue
		} else {
			x = nextX
			y = nextY
		}
	}
}
