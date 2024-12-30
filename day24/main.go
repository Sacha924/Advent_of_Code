package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Equation struct {
	v1, v2, res, op string
}

// remark we could use byte instead of int for memory optim

func main() {
	knownValues, equations := parseInput("day24/input.txt")
	systemRes := ResolveAll(knownValues, equations)
	re := regexp.MustCompile(`z(\d+)`)
	var result [46]int // 46 because the max value of z is 45 in the input (z45) (starting at 00)
	for k, v := range systemRes {
		matches := re.FindStringSubmatch(k)
		if len(matches) != 0 {
			index, _ := strconv.Atoi(matches[1])
			result[index] = v
		}
	}
	decimalValue := binaryArrayToDecimal(result[:])
	fmt.Println(decimalValue)

}

func parseInput(path string) (map[string]int, []Equation) {
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	isEquation := false
	equations := make([]Equation, 0)
	knownValues := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isEquation = true
			continue
		}
		if isEquation {
			parts := strings.Split(line, " ")
			if len(parts) != 5 {
				panic("unexpected input")
			}
			equations = append(equations, Equation{
				v1:  parts[0],
				op:  parts[1],
				v2:  parts[2],
				res: parts[4],
			})
		} else {
			parts := strings.Split(line, ": ")
			val, _ := strconv.Atoi(parts[1])
			knownValues[parts[0]] = val
		}
	}
	return knownValues, equations
}

func ResolveAll(knownValues map[string]int, equations []Equation) map[string]int {
	for len(equations) > 0 {
		var toDelete []int // index of equations resolved that we can remove from our array of equations
		for i := 0; i < len(equations); i++ {
			equation := equations[i]
			if equation.isResolvable(knownValues) {
				toDelete = append(toDelete, i)

				v1, ok1 := knownValues[equation.v1]
				v2, ok2 := knownValues[equation.v2]
				res, _ := knownValues[equation.res]

				if !ok1 {
					knownValues[equation.v1] = Resolve(v2, res, equation.op)
				} else if !ok2 {
					knownValues[equation.v2] = Resolve(v1, res, equation.op)
				} else {
					knownValues[equation.res] = Resolve(v1, v2, equation.op)
				}
			}
		}
		equations = removeByIndices(equations, toDelete)
	}
	return knownValues
}

func Resolve(a, b int, op string) int {
	switch op {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "XOR":
		return a ^ b
	}
	panic("should not reach here")
}

// isResolvable return bool indicating if we can solve the eq
func (e Equation) isResolvable(knownValues map[string]int) bool {
	count := 0
	if _, ok := knownValues[e.v1]; !ok {
		count++
	}
	if _, ok := knownValues[e.v2]; !ok {
		count++
	}
	if _, ok := knownValues[e.res]; !ok {
		count++
	}
	return count == 1
}

func removeByIndices(equations []Equation, indices []int) []Equation {
	// Sort the indices in descending order to avoid shifting issues
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))

	for _, index := range indices {
		equations = append(equations[:index], equations[index+1:]...)
	}

	return equations
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func binaryArrayToDecimal(arr []int) int {
	decimal := 0
	for i, bit := range arr {
		if bit == 1 {
			decimal += 1 << i // 2^i
		}
	}
	return decimal
}
