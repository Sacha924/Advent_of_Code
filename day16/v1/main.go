package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Coord struct {
	x, y int
}

type Direction [2]int

var up = Direction{-1, 0}
var down = Direction{1, 0}
var left = Direction{0, -1}
var right = Direction{0, 1}
var dirs = []Direction{up, left, down, right}

//	type LabCellData struct {
//		Position      Coord
//		PreviousCells map[Coord]bool
//		cost          int
//		previousDir   Direction
//	}
type LabCellData struct {
	Position     Coord
	LastPosition Coord
	cost         int
	previousDir  Direction
}

type LabCellDataHeap []LabCellData

func (h LabCellDataHeap) Len() int           { return len(h) }
func (h LabCellDataHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h LabCellDataHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LabCellDataHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(LabCellData))
}

func (h *LabCellDataHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// test my own heap
//func main() {
//	// Create a priority queue (min-heap)
//	labHeap := &LabCellDataHeap{}
//	heap.Init(labHeap)
//
//	// Push some cells with different costs
//	heap.Push(labHeap, LabCellData{Position: Coord{1, 1}, cost: 10})
//	heap.Push(labHeap, LabCellData{Position: Coord{2, 2}, cost: 5})
//	heap.Push(labHeap, LabCellData{Position: Coord{3, 3}, cost: 15})
//
//	// Pop from the heap and print the results
//	for labHeap.Len() > 0 {
//		cell := heap.Pop(labHeap).(LabCellData)
//		fmt.Println("Popped cell:", cell)
//	}
//}

func main() {
	// Open the file
	file, err := os.Open("./day16/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	labyrinth := make([][]rune, 0)
	for scanner.Scan() {
		labyrinth = append(labyrinth, []rune(scanner.Text()))
	}
	start, end := findStartEnd(labyrinth)
	fmt.Println(start, end)

	var sum = findCheapestPath(labyrinth, start, end)

	fmt.Printf("Sum: %d\n", sum)
}

func findStartEnd(labyrinth [][]rune) (start, end Coord) { // I update the code from V2 that is doing the opposite becuase they are reading coordinate in a different way using an array, not a matrix like me
	for x, line := range labyrinth {
		for y, char := range line {
			if char == 'S' {
				start = Coord{x, y}
			} else if char == 'E' {
				end = Coord{x, y}
			}
		}
	}
	return start, end
}

func findCheapestPath(labyrinth [][]rune, start, end Coord) int {
	labHeap := &LabCellDataHeap{}
	heap.Init(labHeap)

	costMatrix := initMatrix(labyrinth, 999999) // keep minimal cost for current pos
	// Start from the initial position, with "no" previous positions.
	heap.Push(labHeap, LabCellData{
		Position:     start,
		LastPosition: Coord{-1, -1}, // No previous position
		cost:         0,
		previousDir:  right, // Initial direction is East
	})

	for labHeap.Len() > 0 {
		current := heap.Pop(labHeap).(LabCellData)

		for _, dir := range dirs {
			neiCoord := Coord{
				x: current.Position.x + dir[0],
				y: current.Position.y + dir[1],
			}

			if neiCoord != current.LastPosition && labyrinth[neiCoord.x][neiCoord.y] != '#' {
				neiCost := current.cost
				if current.previousDir != dir {
					neiCost += 1000
				}
				neiCost += 1

				if costMatrix[neiCoord.x][neiCoord.y] >= neiCost { // allow for equal if two paths cost the same to reach the pos
					labHeap.Push(LabCellData{
						Position:     neiCoord,
						LastPosition: current.Position,
						cost:         neiCost,
						previousDir:  dir,
					})
					costMatrix[neiCoord.x][neiCoord.y] = neiCost
				}

				// If we've reached the end position, return the cost
				if neiCoord == end {
					//PrintMatrix(costMatrix) // debug (can compare to v2 to see the diff
					return neiCost
				}
			}
		}
	}
	return -1 // No path found
}

func initMatrix(labyrinth [][]rune, initialValue int) [][]int {
	rows := len(labyrinth)
	cols := len(labyrinth[0])
	matrix := make([][]int, rows)

	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = initialValue
		}
	}

	return matrix
}

func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, col := range row {
			// %6d force l'affichage à 6 caractères de large, ajoutant des espaces à gauche si nécessaire
			fmt.Printf("%6d ", col)
		}
		fmt.Println()
	}
}
