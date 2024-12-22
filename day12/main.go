package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// idea :
// first step : dfs on the map, keep track of all visited pos
// for each pos that are not still visited, start a dfs and explore all the pos touching the zone and that have
// the same letter
// at the end of this step we have a list of list of position, where each position in a list have the same letter
// SECOND STEP :
// for each array of positions, compute the perimeter.
// I think for that we can create an adjency list and then compute perimeter: a case count as 4 - nb of neighbor
// then we multiply the perimeter by the length of the current list of node to have his cost

func main() {
	board := readInput("./day12/input.txt")
	res := 0
	visited := make(map[[2]int]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if !visited[[2]int{i, j}] {
				parcel := dfs(board, i, j, visited, board[i][j])
				res += computeParcelScore(parcel)
			}
		}
	}
	fmt.Println(res)
}

func readInput(path string) [][]string {
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	board := make([][]string, 0)
	for scanner.Scan() {
		board = append(board, strings.Split(scanner.Text(), ""))
	}
	return board
}

// target is the letter of the current parcel we are exploring
func dfs(board [][]string, i, j int, visited map[[2]int]bool, target string) [][2]int {
	res := make([][2]int, 0)
	visited[[2]int{i, j}] = true
	res = append(res, [2]int{i, j})
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		newI, newJ := i+dir[0], j+dir[1]
		if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[i]) && !visited[[2]int{newI, newJ}] && board[newI][newJ] == target {
			res = append(res, dfs(board, newI, newJ, visited, target)...)
		}
	}
	return res
}

func computeParcelScore(parcel [][2]int) int {
	area := len(parcel)
	perimeter := computeParcelPerimeter(parcel)
	return area * perimeter
}

//func computeParcelScore2(parcel [][2]int) int {
//	area := len(parcel)
//	sides := computeSides(parcel)
//	return area * sides
//}

// perimeter for a case is 4 - number of neighbor
func computeParcelPerimeter(parcel [][2]int) int {
	res := 0
	lookupTable := make(map[[2]int]bool, 0)
	for _, pos := range parcel {
		lookupTable[pos] = true
	}

	for _, pos := range parcel {
		neiCount := 0
		for _, nei := range getNei(pos) {
			if lookupTable[nei] {
				neiCount++
			}
		}
		res += 4 - neiCount
	}
	return res
}

func getNei(pos [2]int) [][2]int {
	row, col := pos[0], pos[1]
	neighbors := [][2]int{
		{row - 1, col},
		{row + 1, col},
		{row, col - 1},
		{row, col + 1},
	}
	return neighbors
}

//func computeSides(parcel [][2]int) int {
//
//}
