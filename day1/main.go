package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	col1, col2 := readAndProcessFile("day1/input.txt", "   ")
	sort.Ints(col1)
	sort.Ints(col2)
	fmt.Printf("Result for day1 : %d", compare(col1, col2)) // 2756096
}

func readAndProcessFile(path, sep string) ([]int, []int) {
	col1, col2 := make([]int, 1000), make([]int, 1000) // input file contains 1k lines
	f, _ := os.Open(path)
	defer func() { _ = f.Close() }()

	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		lineStr := string(line)
		if len(line) > 0 {
			values := strings.Split(lineStr, sep)
			if len(values) == 2 {
				val1, err1 := strconv.Atoi(values[0])
				if err1 == nil {
					col1 = append(col1, val1)
				}
				val2, err2 := strconv.Atoi(values[1])
				if err2 == nil {
					col2 = append(col2, val2)
				}
			}
		}
		if err != nil {
			break
		}
	}

	return col1, col2
}

func compare(arr1, arr2 []int) int {
	diff := 0
	for i := 0; i < len(arr1); i++ {
		diff += abs(arr1[i] - arr2[i])
	}
	return diff
}

func abs(x int) int { // working with math.Abs is annoying bc it's using float values
	if x < 0 {
		return -x
	}
	return x
}
