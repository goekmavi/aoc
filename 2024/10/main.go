package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var positions = map[string]int{}

func calcTrailHead(currentNr int, mapArray [][]int, i int, j int, fromDirection string, startI int, startJ int) bool {
	if currentNr == 0 {
		positions[strconv.Itoa(startI)+":"+strconv.Itoa(startJ)+" "+strconv.Itoa(i)+":"+strconv.Itoa(j)]++

		return true
	}

	hasTop := false
	hasDown := false
	hasLeft := false
	hasRight := false

	if i-1 >= 0 && fromDirection != "top" { // top
		if mapArray[i-1][j] == currentNr-1 {
			hasTop = calcTrailHead(mapArray[i-1][j], mapArray, i-1, j, "bottom", startI, startJ)
		}
	}

	if i+1 < len(mapArray) && fromDirection != "down" { // down
		if mapArray[i+1][j] == currentNr-1 {
			hasDown = calcTrailHead(mapArray[i+1][j], mapArray, i+1, j, "top", startI, startJ)
		}
	}

	if j-1 >= 0 && fromDirection != "left" { // left
		if mapArray[i][j-1] == currentNr-1 {
			hasLeft = calcTrailHead(mapArray[i][j-1], mapArray, i, j-1, "right", startI, startJ)
		}
	}

	if j+1 < len(mapArray[0]) && fromDirection != "right" { // right
		if mapArray[i][j+1] == currentNr-1 {
			hasRight = calcTrailHead(mapArray[i][j+1], mapArray, i, j+1, "left", startI, startJ)
		}
	}

	if !hasTop && !hasDown && !hasLeft && !hasRight {
		return false
	}

	return true
}

func calcTrailsHeadScore(mapArray [][]int) (int, int) {
	for i, line := range mapArray {
		for j, nr := range line {
			if nr == 9 {
				if i-1 >= 0 { // top
					if mapArray[i-1][j] == 8 {
						calcTrailHead(mapArray[i-1][j], mapArray, i-1, j, "bottom", i, j)
					}
				}

				if i+1 < len(mapArray) { // down
					if mapArray[i+1][j] == 8 {
						calcTrailHead(mapArray[i+1][j], mapArray, i+1, j, "top", i, j)
					}
				}

				if j-1 >= 0 { // left
					if mapArray[i][j-1] == 8 {
						calcTrailHead(mapArray[i][j-1], mapArray, i, j-1, "right", i, j)
					}
				}

				if j+1 < len(mapArray[0]) { // right
					if mapArray[i][j+1] == 8 {
						calcTrailHead(mapArray[i][j+1], mapArray, i, j+1, "left", i, j)
					}
				}
			}
		}
	}

	rating := 0

	for _, v := range positions {
		rating += v
	}

	return len(positions), rating
}

func main() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	mapArray := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		tempArray := []int{}

		for _, v := range line {
			tempArray = append(tempArray, int(v-'0'))
		}

		mapArray = append(mapArray, tempArray)
	}

	score, rating := calcTrailsHeadScore(mapArray)

	fmt.Println("score:", score)
	fmt.Println("rating:", rating)
}
