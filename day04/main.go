package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	A = 65
	M = 77
	S = 83
	X = 88
)

func main() {
	words := fileToWords("day04/input.txt")
	fmt.Printf("result of day4 : %d \n", findAllXmas(words))          // 2551
	fmt.Printf("result of day4 part 2 : %d \n", findXmasPart2(words)) // 2551
}

func fileToWords(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()
	words := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

func findAllXmas(words []string) int {
	counter := 0
	dirs := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	for x, word := range words {
		for y, char := range word {
			if char == X {
				for _, dir := range dirs {
					counter += findXmasHelper(words, x, y, dir[0], dir[1])
				}
			}
		}
	}
	return counter
}

// Once we start searching in a dir, we continue in that dir
// No need of recursion or weird stuff in fact we can just try to advance in one dir to find the word
func findXmasHelper(words []string, x, y, dirX, dirY int) int {
	limitX, limitY := len(words), len(words[0])
	charToFind := getNextCharToFind(words[x][y])
	x += dirX
	y += dirY
	for x < limitX && x >= 0 && y < limitY && y >= 0 {
		if words[x][y] == charToFind {
			if charToFind == S {
				return 1 // We found XMAS
			}
			charToFind = getNextCharToFind(words[x][y])
			x += dirX
			y += dirY
			continue
		} else {
			return 0
		}
	}
	return 0
}

func getNextCharToFind(curr uint8) uint8 {
	switch curr {
	case X:
		return M
	case M:
		return A
	case A:
		return S
	default:
		panic(fmt.Sprintf("should not happened :%d", curr))
	}
}

func findXmasPart2(words []string) int {
	counter := 0
	rows, cols := len(words), len(words[0])
	for x, word := range words {
		for y, char := range word {
			if char == A {
				if 0 < x && x < rows-1 && 0 < y && y < cols-1 {
					counter += isNeighborMatching(words, x, y)
				}
			}
		}
	}
	return counter
}

// to find a double MAS, we need to check the 4 diagonal from 'A', and check that :
// There is 2S and 2M
// the 2S (or 2M) is not on the same diagonal (i.e check S1i != S2i and S1j != S2j)
func isNeighborMatching(words []string, x, y int) int {
	topLeft, topRight := int(words[x-1][y-1]), int(words[x-1][y+1])
	bottomLeft, bottomRight := int(words[x+1][y-1]), int(words[x+1][y+1])

	if (topLeft+topRight+bottomLeft+bottomRight)/4 == 80 && topLeft+bottomRight == 160 {
		return 1
	}
	return 0
}
