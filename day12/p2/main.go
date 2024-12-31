package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ x, y int }

func parseInput() (map[P]uint8, int) {
	scanner := bufio.NewScanner(os.Stdin)
	m := map[P]uint8{}
	var i int
	for i = 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			m[P{i, j}] = uint8(c)
		}
	}
	return m, i
}

func main() { // go run day12/p2/main.go < day12/input.txt
	m, size := parseInput()
	price := 0
	visited := map[P]struct{}{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if _, ok := visited[P{i, j}]; ok {
				continue
			}
			area, fence := explore(m, P{i, j}, visited)
			price += area * fence
		}
	}
	fmt.Println(price)
}

func explore(m map[P]uint8, p P, visited map[P]struct{}) (int, int) {
	visited[p] = struct{}{}

	id := m[p]

	area, corners := 0, 0

	var c P
	toDo := []P{p}
	for len(toDo) > 0 {
		c, toDo = toDo[0], toDo[1:]

		area++

		for _, d := range []P{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := P{c.x + d.x, c.y + d.y}
			val, ok := m[next]
			if ok && val == id {
				if _, ok := visited[next]; !ok {
					toDo = append(toDo, next)
					visited[next] = struct{}{}
				}
			}
		}

		for _, d := range []P{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}} {
			// convex corner
			if m[P{c.x + d.x, c.y}] != id &&
				m[P{c.x, c.y + d.y}] != id {
				corners++
			}
			// concave corner
			if m[P{c.x + d.x, c.y}] == id &&
				m[P{c.x, c.y + d.y}] == id &&
				m[P{c.x + d.x, c.y + d.y}] != id {
				corners++
			}
		}

	}

	return area, corners
}
