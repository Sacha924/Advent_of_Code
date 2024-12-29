package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("result of day2 : %d \n", readAndProcessFile("day02/input.txt"))
}

func readAndProcessFile(path string) int {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		values := stringSliceToInt(strings.Fields(line))
		res += isSafeReport(values)
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
	return res
}

func stringSliceToInt(slice []string) []int {
	res := make([]int, len(slice))
	for i, v := range slice {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		res[i] = intVal
	}
	return res
}

// return 1 if safe else 0
func isSafeReport(values []int) int {
	if len(values) < 2 {
		return 1
	}
	if values[0]-values[1] == 0 {
		return 0
	}
	asc := values[0]-values[1] < 0

	// stillvalid is used for the 2nd part of exercise day2
	stillValid := true
	for i := 0; i < len(values)-1; i++ {
		diff := values[i] - values[i+1]
		if (asc && diff <= -1 && diff >= -3) || (!asc && diff >= 1 && diff <= 3) {
			continue
		} else if stillValid {
			stillValid = false
			continue
		}
		return 0
	}
	return 1
}

// INTERESTING : https://stackoverflow.com/questions/76811065/why-does-issorted-in-the-standard-library-iterate-over-the-slice-in-reverse
