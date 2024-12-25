package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type RobotStat struct {
	row    int
	col    int
	speedX int
	speedY int
}

const (
	width      = 101
	height     = 103
	nbOfRounds = 100
)

func main() {
	robots := getRobots("day14/input.txt")
	//start := time.Now()
	//fmt.Println(Play(robots, nbOfRounds))
	//fmt.Println("Time elapsed: ", time.Since(start))

	i := 1
	for i < 104 {
		fmt.Printf("round %d:\n", i)
		PlayAndPrint(robots)
		i++
	}
	//there is a cycle so for me everytime a number modulo 103 is equal to 42 the values are align vertically, and every time a number modulo 101 (width) the points are align horizontaly)
	//h = mod(42,103)
	//w = mod(99,101)
	//
	//both are true when (we use chinese reminder theorem to solve a system of congruence)
	// res = 8179
}

func getRobots(filename string) []RobotStat {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	robots := make([]RobotStat, 0)
	//p=2,4 v=2,-3
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		speedX, _ := strconv.Atoi(matches[3])
		speedY, _ := strconv.Atoi(matches[4])
		if len(matches) == 5 {
			robots = append(robots, RobotStat{
				row:    x,
				col:    y,
				speedX: speedX,
				speedY: speedY,
			})
		}
	}
	return robots
}

func Play(robots []RobotStat, nbOfRounds int) int {
	TL, TR, BL, BR := 0, 0, 0, 0 // quadrants
	halfWidth := width / 2
	halfHeight := height / 2
	for _, robot := range robots {
		finalX, finalY := (robot.row+robot.speedX*nbOfRounds)%width, (robot.col+robot.speedY*nbOfRounds)%height
		if finalX < 0 {
			finalX = width + finalX
		}
		if finalY < 0 {
			finalY = height + finalY
		}
		if finalX == halfWidth || finalY == halfHeight {
			continue // Robots that are exactly in the middle (horizontally or vertically) don't count as being in any quadrant
		}
		if finalX < halfWidth && finalY < halfHeight { // Top-Left quadrant
			TL++
		} else if finalX >= halfWidth && finalY < halfHeight { // Top-Right quadrant
			TR++
		} else if finalX < halfWidth && finalY >= halfHeight { // Bottom-Left quadrant
			BL++
		} else { // Bottom-Right quadrant
			BR++
		}
	}
	return TL * TR * BL * BR
}

// at the begining I add no idea how I can do that without doing it manually
// saw this video https://www.youtube.com/watch?v=hhRC8XrXY1o&ab_channel=bmenrigh
// I could do an algo that check for clusters, but the way it's solved using mod and
// https://en.wikipedia.org/wiki/Chinese_remainder_theorem is really interesting and I want to give it a try
// let's start by displaying the board game at every turn and find one cluster horiz and one vert

func PlayAndPrint(robots []RobotStat) {
	for i, robot := range robots {
		newRow, newCol := (robot.row+robot.speedX)%width, (robot.col+robot.speedY)%height
		if newRow < 0 {
			newRow = width + newRow
		}
		if newCol < 0 {
			newCol = height + newCol
		}
		robots[i].row, robots[i].col = newRow, newCol
	}
	PrintBoard(robots)
}

func PrintBoard(robots []RobotStat) {
	var matrix [width][height]string
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			matrix[i][j] = ". "
		}
	}

	// Place robots in the matrix
	for _, robot := range robots {
		matrix[robot.row][robot.col] = "# "
	}

	var output string
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			output += matrix[i][j]
		}
		output += "\n"
	}
	fmt.Println(output)
	//textView.SetText(output)
}
