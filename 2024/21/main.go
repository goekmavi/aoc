package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	I int
	J int
}

type Code struct {
	codes [][]string
}

var numericPad = map[string]node{
	"7": {
		I: 0,
		J: 0,
	},
	"8": {
		I: 0,
		J: 1,
	},
	"9": {
		I: 0,
		J: 2,
	},
	"4": {
		I: 1,
		J: 0,
	},
	"5": {
		I: 1,
		J: 1,
	},
	"6": {
		I: 1,
		J: 2,
	},
	"1": {
		I: 2,
		J: 0,
	},
	"2": {
		I: 2,
		J: 1,
	},
	"3": {
		I: 2,
		J: 2,
	},
	"_X": {
		I: 3,
		J: 0,
	},
	"0": {
		I: 3,
		J: 1,
	},
	"A": {
		I: 3,
		J: 2,
	},
}

var directionalPad = map[string]node{
	"_X": {
		I: 0,
		J: 0,
	},
	"^": {
		I: 0,
		J: 1,
	},
	"A": {
		I: 0,
		J: 2,
	},
	"<": {
		I: 1,
		J: 0,
	},
	"v": {
		I: 1,
		J: 1,
	},
	">": {
		I: 1,
		J: 2,
	},
}

var codes = [][]string{}

func numericPart(code []string) int {
	_str := strings.Join(code, "")
	_temp := strings.Replace(_str, "A", "", -1)

	nmbr, _ := strconv.Atoi(_temp)

	return nmbr
}

func updateFinalCodes(finalCodes map[int]Code) [][]string {
	length := len(finalCodes)
	tempCodes := [][]string{}

	tempCodes = append(tempCodes, finalCodes[0].codes...)

	for i := 1; i < length; i++ {
		_temp := [][]string{}
		for _, v := range finalCodes[i].codes {
			for _, t := range tempCodes {
				_temp = append(_temp, append(t, v...))
			}
		}
		tempCodes = _temp
	}

	return tempCodes
}

func calcPath(code []string, pad map[string]node) []string {
	fmt.Println(code)
	finalCodes := map[int]Code{}

	startNode := pad["A"]

	for round, c := range code {
		tempCode1 := []string{}
		tempCode2 := []string{}

		destinationNode := pad[c]

		diffI := destinationNode.I - startNode.I
		diffJ := destinationNode.J - startNode.J

		if diffJ > 0 { // >
			for range diffJ {
				tempCode1 = append(tempCode1, ">")
			}
		} else if diffJ < 0 { // <
			for range diffJ * -1 {
				tempCode1 = append(tempCode1, "<")
			}
		}
		if diffI > 0 { // v
			for range diffI {
				tempCode1 = append(tempCode1, "v")
			}
		} else if diffI < 0 { // ^
			for range diffI * -1 {
				tempCode1 = append(tempCode1, "^")
			}
		}

		if diffI > 0 { // v
			for range diffI {
				tempCode2 = append(tempCode2, "v")
			}
		} else if diffI < 0 { // ^
			for range diffI * -1 {
				tempCode2 = append(tempCode2, "^")
			}
		}
		if diffJ > 0 { // >
			for range diffJ {
				tempCode2 = append(tempCode2, ">")
			}
		} else if diffJ < 0 { // <
			for range diffJ * -1 {
				tempCode2 = append(tempCode2, "<")
			}
		}

		tempCode1 = append(tempCode1, "A") // press
		tempCode2 = append(tempCode2, "A") // press

		finalCodes[round] = Code{
			codes: [][]string{tempCode1, tempCode2},
		}

		startNode = destinationNode
	}

	_codes := updateFinalCodes(finalCodes)

	min := 0

	for i := range _codes {
		if len(_codes[min]) > len(_codes[i]) {
			min = i
		}
	}

	return _codes[min]
}

func shortestPath(code []string) int {
	fmt.Println("code", code)
	robot1 := calcPath(code, numericPad)
	robot2 := calcPath(robot1, directionalPad)
	robot3 := calcPath(robot2, directionalPad)

	fmt.Println("r1", len(robot1), robot1)
	fmt.Println("r2", len(robot2), robot2)
	fmt.Println("r3", len(robot3), robot3)

	return len(robot2)
}

func calcComplexity() int {
	sum := 0

	for _, code := range codes {
		shortest_path := shortestPath(code)
		numeric_part := numericPart(code)

		sum += (shortest_path * numeric_part)
	}

	return sum
}

func main() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := []string{}

		for _, v := range line {
			lineArr = append(lineArr, string(v))
		}

		codes = append(codes, lineArr)
	}

	complexity := calcComplexity()
	fmt.Println("complexity", complexity)
}
