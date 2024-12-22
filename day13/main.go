package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// The day13 starts by resolving a system:
// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400
// It can be described as :
// 94a + 22b = 8400
// 34a + 67b = 5400
// from what I remember in high school :
// I can eliminate one of the 2 values, either a or b, let's say I eliminate a.
// Then I know the value of b and i determine the value of a and b, and the tocal cost knowing total cost is
// 3A + B
// note that if a and b are not integers, it means the solution is not acceptable

type System [2][3]int

func main() {
	// equations is an array of size 2 containing array of size 3 where first col is a, second col is b, thrid col is result.
	// first array is eq 1 and 2nd array is eq2
	equations := parseInput("./day13/input.txt")
	solutions := resolveEquations(equations)
	fmt.Println(costToWinEachPrize(solutions))
}

func parseInput(path string) [][2][3]int {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var equations [][2][3]int
	reButtonA := regexp.MustCompile(`Button A: X\+([0-9]+), Y\+([0-9]+)`) // rq : [0-9] or \d is equivalent
	reButtonB := regexp.MustCompile(`Button B: X\+([0-9]+), Y\+([0-9]+)`)
	rePrize := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)
	lineCount := 0

	var eq1, eq2 [3]int

	for scanner.Scan() {
		line := scanner.Text()
		var matches []string
		switch lineCount % 4 {
		case 0:
			matches = reButtonA.FindStringSubmatch(line)
			a, b := matchesToInt(matches)
			eq1[0] = a
			eq2[0] = b
		case 1:
			matches = reButtonB.FindStringSubmatch(line)
			a, b := matchesToInt(matches)
			eq1[1] = a
			eq2[1] = b
		case 2:
			matches = rePrize.FindStringSubmatch(line)
			targetX, targetY := matchesToInt(matches)
			eq1[2] = targetX // + 10000000000000         JUST UNCOMMENT THESE TWO LINES TO GET RESULT OF DAY 2
			eq2[2] = targetY // + 10000000000000
			equations = append(equations, [2][3]int{eq1, eq2})
		}
		lineCount++
	}
	return equations
}

func matchesToInt(matches []string) (int, int) {
	// if len(matches) > 0 { NO ERROR HANDLING, EXPECTING input is always correct (which is the case in AOC)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return x, y
}

// resolve Plan :
// STEP 1
// mult L1 per b coeff from L2
// mult L2 par b coeff from L1
// STEP 2
// do L1 - L2, it will remove the b, i.e  someInt * A = someInt
// STEP 3
// we can then determine A
// and then B
func resolveEquations(equations [][2][3]int) [][2]int {
	solutions := make([][2]int, 0)
	for _, system := range equations {
		var a, b int
		eq1, eq2 := system[0], system[1]

		coefBL1 := eq1[1]
		coefBL2 := eq2[1]
		for i := 0; i < 3; i++ { // STEP 1
			eq1[i] *= coefBL2
			eq2[i] *= coefBL1
		}

		for i := 0; i < 3; i++ { // STEP 2
			eq1[i] -= eq2[i]
		}

		// STEP 3
		if eq1[2]%eq1[0] == 0 { // We have something like 4a = 15, we verify that the solution is an int
			a = eq1[2] / eq1[0]
		} else {
			continue
		}
		if (eq2[2]-a*eq2[0])%eq2[1] == 0 { // same for b
			b = (eq2[2] - a*eq2[0]) / eq2[1]
		} else {
			continue
		}
		solutions = append(solutions, [2]int{a, b})
	}
	return solutions
}

func costToWinEachPrize(solutions [][2]int) int {
	cost := 0
	for _, solution := range solutions {
		cost += solution[0]*3 + solution[1]
	}
	return cost
}
