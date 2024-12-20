package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nodeMap = map[string]Node{}
var coordinates = map[string]Node{}
var strArr = [][]string{}
var iLength, jLength = 71, 71
var startX, startY = 0, 0
var endX, endY = iLength - 1, jLength - 1

type Node struct {
	X     int
	Y     int
	Value string
}

func createFile(arr [][]string) {
	file, err := os.Create("./newInput.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	for i, v := range arr {
		line := strings.Join(v, "")

		_, err := file.WriteString(line)

		if err != nil {
			panic(err)
		}

		if i < len(arr)-1 {
			_, err := file.WriteString("\n")

			if err != nil {
				panic(err)
			}
		}

		fmt.Println(line)
	}

	fmt.Println("New file is created succesfully!") // use code from day 16
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
		data := strings.Split(line, `,`)

		xStr := data[0]
		yStr := data[1]

		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)

		coordinates[yStr+":"+xStr] = Node{
			X:     x,
			Y:     y,
			Value: "#",
		}
	}

	for i := 0; i < iLength; i++ {
		_line := []string{}

		for j := 0; j < jLength; j++ {
			_str := strconv.Itoa(i) + ":" + strconv.Itoa(j)

			_, _vExists := coordinates[_str]

			if _vExists {
				nodeMap[_str] = Node{
					X:     j,
					Y:     i,
					Value: "#",
				}

				_line = append(_line, "#")
			} else {
				nodeMap[_str] = Node{
					X:     j,
					Y:     i,
					Value: ".",
				}

				_line = append(_line, ".")
			}
		}

		strArr = append(strArr, _line)
	}

	strArr[startY][startX] = "S"
	strArr[endY][endX] = "E"

	createFile(strArr)
}
