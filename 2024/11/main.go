package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func blink(stones []int, amount int) int {
	for i := 0; i < amount; i++ {
		for j := 0; j < len(stones); j++ {
			currentNr := stones[j]
			currentNrAsString := strconv.Itoa(currentNr)

			if currentNr == 0 {
				stones[j] = 1
			} else if len(currentNrAsString)%2 == 0 {
				leftPart := ""
				rightPart := ""

				onlyZerosOnLeftSide := true
				onlyZerosOnRightSide := true

				for k := 0; k < len(currentNrAsString); k++ {
					if k < len(currentNrAsString)/2 {
						leftPart += string(currentNrAsString[k])
					} else {
						rightPart += string(currentNrAsString[k])
					}
				}

				for _, v := range leftPart {
					if string(v) != "0" {
						onlyZerosOnLeftSide = false
						break
					}
				}

				for _, v := range rightPart {
					if string(v) != "0" {
						onlyZerosOnRightSide = false
						break
					}
				}

				if onlyZerosOnLeftSide {
					leftPart = "0"
				}

				if onlyZerosOnRightSide {
					rightPart = "0"
				}

				leftNr, errLeftNr := strconv.Atoi(leftPart)
				rightNr, errRightNr := strconv.Atoi(rightPart)

				if errLeftNr != nil && errRightNr != nil {
					panic(errLeftNr)
				}

				tempSlice := make([]int, len(stones))
				copy(tempSlice, stones)

				tempSliceTwo := make([]int, len(stones))
				copy(tempSliceTwo, stones)

				_first := append(tempSlice[:j], []int{leftNr, rightNr}...)
				_second := tempSliceTwo[j+1:]

				stones = append(_first, _second...)

				j = j + 1
			} else {
				stones[j] *= 2024
			}
		}
	}

	return len(stones)
}

func main() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	stones := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		temp := ""

		for i, v := range line {
			if string(v) != " " {
				temp += string(v)
			} else {
				nr, errNr := strconv.Atoi(temp)

				if errNr != nil {
					panic(errNr)
				}

				stones = append(stones, nr)
				temp = ""
			}

			if i == len(line)-1 {
				nr, errNr := strconv.Atoi(temp)

				if errNr != nil {
					panic(errNr)
				}

				stones = append(stones, nr)
				temp = ""
			}
		}
	}

	fmt.Println(blink(stones, 25))
	//fmt.Println(blink(stones, 3))
}
