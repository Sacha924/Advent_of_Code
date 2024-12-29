package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x, y int
}
type Direction = Coord

var up = Direction{-1, 0}
var down = Direction{1, 0}
var left = Direction{0, -1}
var right = Direction{0, 1}
var dirs = []Direction{up, left, down, right}

func main() {
	lab := parseInput("day20/input.txt")
	nodesFromTrack := getNodesFromTrack(lab) // return the nodes we are traversing to go grom E to S (only one path possible according to the problem statement)
	fmt.Println(countShortCut(lab, nodesFromTrack))
}

func parseInput(filename string) [][]rune {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lab := make([][]rune, 0)
	for scanner.Scan() {
		lab = append(lab, []rune(scanner.Text()))
	}
	return lab
}

func findPos(lab [][]rune, char rune) Coord {
	for i := 0; i < len(lab); i++ {
		for j := 0; j < len(lab[i]); j++ {
			if lab[i][j] == char {
				return Coord{i, j}
			}
		}
	}
	return Coord{}
}

func getNodesFromTrack(lab [][]rune) []Coord {
	start, end := findPos(lab, 'S'), findPos(lab, 'E')
	nodes := make([]Coord, 0)
	previous := Coord{-1, -1}
	curr := start
	for curr != end {
		nodes = append(nodes, Coord{curr.x, curr.y})
		for _, dir := range dirs { // only one possible position each round (if it's a wall or the previous pos, it's not a valid position)
			newX, newY := curr.x+dir.x, curr.y+dir.y
			if (newX == previous.x && newY == previous.y) || lab[newX][newY] == '#' {
				continue
			}
			previous = curr
			curr = Coord{newX, newY}
			break
		}
	}
	nodes = append(nodes, end)
	return nodes
}

func countShortCut(lab [][]rune, nodesFromTrack []Coord) int {
	// shortcut reduce this cost by the difference in index of the second node - first node
	// valid shortcut if dist between two nodes is 2 vertically or horizontally
	res := 0
	for i := 0; i < len(nodesFromTrack); i++ {
		for j := i + 102; j < len(nodesFromTrack); j++ { // i + 100 + 2 because we want a shortcut that make us win at least 100 picoseconds, which mean the two cases must be spaced by at least 100 +2 because doing the move while cheating cost 2 picoseconds
			// PART 1
			//if isShortcutable(lab, nodesFromTrack[i], nodesFromTrack[j]) {
			//	res++
			//}
			// PART 2
			ok, nbOfMoveInShortcut := isShortcutablePart2(lab, nodesFromTrack[i], nodesFromTrack[j])
			// if the dist between two node is more than 100 but less than 100 + nb of moves to make the shortcut,
			// we are not saving at least 100 picoseconds, so we need to check :
			if ok && j-i-nbOfMoveInShortcut >= 100 {
				res++
			}
		}
	}
	return res
}

func isShortcutable(lab [][]rune, node1, node2 Coord) bool {
	if (node2.x == node1.x-2 && node2.y == node1.y) || (node2.x == node1.x+2 && node2.y == node1.y) || (node2.x == node1.x && node2.y == node1.y-2) || (node2.x == node1.x && node2.y == node1.y+2) {
		if lab[(node2.x+node1.x)/2][(node2.y+node1.y)/2] == '#' { // wall between the two nodes
			return true
		}
	}
	return false
}

func isShortcutablePart2(lab [][]rune, node1, node2 Coord) (bool, int) {
	dist := abs(node1.x-node2.x) + abs(node1.y-node2.y)
	if dist <= 20 {
		return true, dist
	}
	return false, -1
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
