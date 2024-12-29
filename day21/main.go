package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Coord struct {
	x, y int
}

func (c Coord) isEqual(c2 Coord) bool {
	return c.x == c2.x && c.y == c2.y
}

const (
	left  = "<"
	right = ">"
	up    = "^"
	down  = "v"
)

var numericPad = map[rune]Coord{
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	'0': {3, 1},
	'A': {3, 2},
}
var unauthorizedNumericPadPos = Coord{3, 0}

var movesPad = map[rune]Coord{
	'^': {0, 1},
	'A': {0, 2},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}
var unauthorizedMovePadPos = Coord{0, 0}

func main() {
	res := 0
	start := time.Now()
	for _, code := range parseInput("day21/input.txt") {
		transfo := inputToOutputMoves(code, numericPad, unauthorizedNumericPadPos, false)
		newTransfo := make([]string, 0)
		for range 2 {
			for _, t := range transfo {
				newTransfo = append(newTransfo, keepShortestTransfos(inputToOutputMoves(t, movesPad, unauthorizedMovePadPos, true))...)
			}
			transfo = newTransfo
			newTransfo = make([]string, 0)
		}
		codeNumValue, _ := strconv.Atoi(code[0:3])
		res += findMinStringLength(transfo) * codeNumValue
	}
	fmt.Printf("Time elapsed: %s\n", time.Since(start))
	fmt.Println(res)
}

func parseInput(filename string) []string {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	codes := make([]string, 0)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}
	return codes
}

type memoObject struct {
	currentPos Coord
	nextPos    Coord
}

var memo = make(map[memoObject][]string)

// either we use numeric pad for the first code, or movesPad
func inputToOutputMoves(code string, pad map[rune]Coord, illegalPos Coord, useMemo bool) []string {
	currentPos := pad['A'] //  start pos is A
	outputs := make([]string, 1)
	for _, char := range code {
		charPos := pad[char]
		var movesCandidate []string
		if val, ok := memo[memoObject{currentPos: currentPos, nextPos: charPos}]; ok && useMemo {
			movesCandidate = val
		} else {
			diffX, diffY := charPos.x-currentPos.x, charPos.y-currentPos.y
			movesCandidate = distanceToMovesInput(diffX, diffY, illegalPos, currentPos)
			memo[memoObject{currentPos: currentPos, nextPos: charPos}] = movesCandidate
		}
		if len(movesCandidate) == 1 {
			for i := range outputs {
				outputs[i] += movesCandidate[0]
			}
		} else if len(movesCandidate) == 2 {
			newOutputs := make([]string, 0, len(outputs)*2)
			for _, output := range outputs {
				newOutputs = append(newOutputs, output+movesCandidate[0])
				newOutputs = append(newOutputs, output+movesCandidate[1])
			}
			outputs = newOutputs
		} else {
			panic("impossible")
		}

		currentPos = charPos
	}
	return outputs
}

func distanceToMovesInput(x, y int, illegalPos, currentPos Coord) []string { // e.g we need to 2 times on the right, this func will return >>
	var output string
	if x/-1 > 0 {
		for range abs(x) {
			output += up
		}
	} else if x > 0 {
		for range x {
			output += down
		}
	}
	if y/-1 > 0 {
		for range abs(y) {
			output += left
		}
	} else if y > 0 {
		for range y {
			output += right
		}
	}
	// we need to return two ouputs because there is two potential best path. if we do for example ^^>, the best path action can be >^^ or ^^>
	// remark : no need to do that when len(output) == 1 or output and reverse(output)
	if len(output) <= 1 || output == Reverse(output) {
		return []string{output + "A"}
	}
	return areMovesAllowed([]string{output, Reverse(output)}, currentPos, illegalPos)
}

func areMovesAllowed(moves []string, currentMatrixPos, illegalPos Coord) []string {
	outputs := make([]string, 0)
	for _, move := range moves {
		isLegal := true
		virtualPos := currentMatrixPos
		for _, char := range move {
			switch char {
			case '>':
				virtualPos.y += 1
			case '<':
				virtualPos.y -= 1
			case 'v':
				virtualPos.x += 1
			case '^':
				virtualPos.x -= 1
			}
			if virtualPos.isEqual(illegalPos) {
				isLegal = false
				break
			}
		}
		if isLegal {
			outputs = append(outputs, move+"A")
		}
	}
	return outputs
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findMinStringLength(strings []string) int {
	if len(strings) == 0 {
		return 0
	}
	minStrLen := len(strings[0])
	for _, str := range strings[1:] {
		if len(str) < minStrLen {
			minStrLen = len(str)
		}
	}
	return minStrLen
}

func keepShortestTransfos(strings []string) []string {
	if len(strings) == 0 {
		return nil
	}
	minLen := findMinStringLength(strings)
	var result []string
	for _, str := range strings {
		if len(str) == minLen {
			result = append(result, str)
		}
	}

	return result
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
