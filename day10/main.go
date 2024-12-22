package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	startPos, gameBoard := readAndParseInput("./day10/input.txt")
	res := 0
	for _, pos := range startPos {
		res += dfs(gameBoard, pos[0], pos[1], 1, make(map[[2]int]bool))
	}
	res2 := 0
	for _, pos := range startPos {
		res2 += dfs2(gameBoard, pos[0], pos[1], 1)
	}
	fmt.Println(res)
	fmt.Println(res2)
}

func readAndParseInput(path string) ([][2]int, [][]int) {
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	board := make([][]int, 0)
	zeroPos := make([][2]int, 0)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		numbersStr := strings.Split(line, "")
		for col, numberStr := range numbersStr {
			number, _ := strconv.Atoi(numberStr)
			row = append(row, number)
			if number == 0 {
				zeroPos = append(zeroPos, [2]int{rowCount, col})
			}
		}
		board = append(board, row)
		rowCount++
	}
	return zeroPos, board
}

// visited keeps track of the head (9) already visited
func dfs(board [][]int, row, col, target int, visited map[[2]int]bool) int {
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		newX, newY := row+dir[0], col+dir[1]
		if newX >= 0 && newX < len(board) && newY >= 0 && newY < len(board[newX]) && board[newX][newY] == target {
			if target == 9 {
				visited[[2]int{newX, newY}] = true
			} else {
				dfs(board, newX, newY, target+1, visited)
			}
		}
	}
	return len(visited)
}

func dfs2(board [][]int, row int, col, target int) int {
	res := 0
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range dirs {
		newX, newY := row+dir[0], col+dir[1]
		if newX >= 0 && newX < len(board) && newY >= 0 && newY < len(board[newX]) && board[newX][newY] == target {
			if target == 9 {
				res += 1
			}
			res += dfs2(board, newX, newY, target+1)
		}
	}
	return res
}
