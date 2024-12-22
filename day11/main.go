package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// PART 1 / NON OPTIMIZED CODE
//func main() {
//	input := "4 4841539 66 5279 49207 134 609568 0"
//	arr := make([]int, 0)
//	for _, nbStr := range strings.Split(input, " ") {
//		nb, _ := strconv.Atoi(nbStr)
//		arr = append(arr, nb)
//	}
//
//	for i := 0; i < 25; i++ {
//		arr = applyRule(arr)
//	}
//	fmt.Println(len(arr))
//}
//
//func applyRule(arr []int) []int {
//	res := make([]int, 0)
//	for _, stone := range arr {
//		if stone == 0 {
//			res = append(res, 1)
//		} else if nbOfDigit := getNbOfDigit(stone); nbOfDigit%2 == 0 {
//			lnb, rnb := SplitEvenDigit(stone, nbOfDigit)
//			res = append(res, lnb, rnb)
//		} else {
//			res = append(res, stone*2024)
//		}
//	}
//	return res
//}
//
//func getNbOfDigit(nb int) int {
//	count := 1
//	for nb >= 10 {
//		nb /= 10
//		count++
//	}
//	return count
//}
//
//func SplitEvenDigit(nb, length int) (int, int) {
//	arr := numberToArray(nb)
//	l, r := 0, 0
//	i := 0
//	for i < length/2 {
//		l = l*10 + arr[i]
//		i++
//	}
//	for i < length {
//		r = r*10 + arr[i]
//		i++
//	}
//	return l, r
//}
//
//func numberToArray(n int) []int {
//	var result []int
//	str := strconv.Itoa(n)
//	for _, char := range str {
//		digit, _ := strconv.Atoi(string(char))
//		result = append(result, digit)
//	}
//	return result
//}

// PART 2 FIRST TRY
//func main() {
//	input := "4 4841539 66 5279 49207 134 609568 0"
//	arr := make([]int, 0)
//	for _, nbStr := range strings.Split(input, " ") {
//		nb, _ := strconv.Atoi(nbStr)
//		arr = append(arr, nb)
//	}
//
//	start := time.Now()
//	for i := 0; i < 42; i++ {
//		fmt.Print("round ", i)
//		arr = applyRule(arr)
//	}
//	fmt.Println(len(arr))
//	fmt.Println(time.Since(start))
// for 38 iterations
// 4.65s
// 4.20s by adding capacity of the arr to prevent arr resizing every time
// removing the old logique to know if a number is even, and the logic to split the number in two
// by directly using len of the number converted in string and slice splitting
// 1.483396554s

// bumping to 45 iterations
// 8.140721614s
// we still have an issue, the code is taking too long as the string becomes bigger and bigger...

// we can't optimize a process that is still inefficient in terms of datastructure, we need to keep track of the stone
// without using an array to store everything, we could try with a hashmap
// let's notice that the order of the stones didn't matter
//}

//func applyRule(arr []int) []int {
//	res := make([]int, 0, len(arr))
//	for _, stone := range arr {
//		if stone == 0 {
//			res = append(res, 1)
//		} else if nbString := strconv.Itoa(stone); len(nbString)%2 == 0 {
//			l, r := nbString[0:len(nbString)/2], nbString[len(nbString)/2:]
//			nbL, _ := strconv.Atoi(l)
//			nbR, _ := strconv.Atoi(r)
//			res = append(res, nbL, nbR)
//		} else {
//			res = append(res, stone*2024)
//		}
//	}
//	return res
//}

func main() {
	input := "4 4841539 66 5279 49207 134 609568 0"
	hash := make(map[int]int, 0)
	for _, nbStr := range strings.Split(input, " ") {
		nb, _ := strconv.Atoi(nbStr)
		hash[nb] += 1
	}

	start := time.Now()
	for i := 0; i < 75; i++ {
		hash = applyRule(hash)
	}
	res := 0
	for _, v := range hash {
		res += v
	}
	fmt.Println(res)
	fmt.Println(time.Since(start))
	// now 42 iterations takes 3.243193ms instead of the previous 8s !!!
}

func applyRule(hash map[int]int) map[int]int {
	newState := make(map[int]int, 0)
	for stone, nb := range hash {
		if stone == 0 {
			// convert all '0' stones into '1' stones
			newState[1] += nb
		} else if nbString := strconv.Itoa(stone); len(nbString)%2 == 0 {
			l, r := nbString[0:len(nbString)/2], nbString[len(nbString)/2:]
			nbL, _ := strconv.Atoi(l)
			nbR, _ := strconv.Atoi(r)
			newState[nbL] += nb
			newState[nbR] += nb

		} else {
			newState[stone*2024] += nb
		}
	}
	return newState
}
