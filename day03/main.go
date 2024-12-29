package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// PART 1
//func main() {
//	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
//
//	data, err := os.ReadFile("day03/input.txt")
//
//	if err != nil {
//		panic(err)
//	}
//	arr := re.FindAllStringSubmatch(string(data), -1)
//	res := 0
//	for _, match := range arr {
//		re2 := regexp.MustCompile(`-?[0-9]+`)
//		numbers := re2.FindAllString(match[0], -1)
//		num1, _ := strconv.Atoi(numbers[0])
//		num2, _ := strconv.Atoi(numbers[1])
//		res += num1 * num2
//	}
//	fmt.Printf("result of day2 : %d \n", res)
//}

// PART 2
func main() {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|don't|do`) // don't before do, order matter

	data, err := os.ReadFile("day03/input.txt")

	if err != nil {
		panic(err)
	}
	arr := re.FindAllStringSubmatch(string(data), -1)

	res := 0
	do := true
	for _, match := range arr {
		if match[0] == "do" {
			do = true
		} else if match[0] == "don't" {
			do = false
		} else {
			if do {
				re2 := regexp.MustCompile(`-?[0-9]+`)
				numbers := re2.FindAllString(match[0], -1)
				num1, _ := strconv.Atoi(numbers[0])
				num2, _ := strconv.Atoi(numbers[1])
				res += num1 * num2
			}
		}
	}
	fmt.Printf("result of day2 : %d \n", res)
}
