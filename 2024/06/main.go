package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func initMap() ([][]string, struct {
	X         int
	Y         int
	Direction string
}) {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := [][]string{}
	indexLine := 0

	startPos := struct {
		X         int
		Y         int
		Direction string
	}{
		X:         0,
		Y:         0,
		Direction: "",
	}

	for scanner.Scan() {
		line := scanner.Text()
		tempArr := []string{}

		for i, v := range line {
			currentItem := string(v)

			if currentItem == "^" || currentItem == "v" || currentItem == ">" || currentItem == "<" {
				startPos.X = i
				startPos.Y = indexLine

				switch currentItem {
				case "^":
					startPos.Direction = "up"
				case "v":
					startPos.Direction = "down"
				case ">":
					startPos.Direction = "right"
				case "<":
					startPos.Direction = "left"
				}
			}

			tempArr = append(tempArr, currentItem)
		}

		arr = append(arr, tempArr)
		indexLine++
	}

	return arr, startPos
}

func exists(visited map[string]bool, pos struct {
	X         int
	Y         int
	Direction string
}) bool {
	_, contains := visited[strconv.Itoa(pos.X)+":"+strconv.Itoa(pos.Y)+":"+pos.Direction]

	if contains {
		return true
	} else {
		return false
	}
}

func isLoop(arr [][]string, pos struct {
	X         int
	Y         int
	Direction string
}) bool {
	isDone := false

	visitedArr := map[string]bool{
		strconv.Itoa(pos.X) + ":" + strconv.Itoa(pos.Y) + ":" + pos.Direction: true,
	}

	for !isDone {
		if pos.Direction == "up" {
			if pos.Y-1 < 0 {
				isDone = true
			} else {
				if arr[pos.Y-1][pos.X] == "#" {
					pos.Direction = "right"
				} else {
					pos.Y = pos.Y - 1

					if exists(visitedArr, pos) {
						return true
					}

					visitedArr[strconv.Itoa(pos.X)+":"+strconv.Itoa(pos.Y)+":"+pos.Direction] = true
				}
			}
		} else if pos.Direction == "down" {
			if pos.Y+1 > len(arr)-1 {
				isDone = true
			} else {
				if arr[pos.Y+1][pos.X] == "#" {
					pos.Direction = "left"
				} else {
					pos.Y = pos.Y + 1

					if exists(visitedArr, pos) {
						return true
					}

					visitedArr[strconv.Itoa(pos.X)+":"+strconv.Itoa(pos.Y)+":"+pos.Direction] = true
				}
			}
		} else if pos.Direction == "right" {
			if pos.X+1 > len(arr[0])-1 {
				isDone = true
			} else {
				if arr[pos.Y][pos.X+1] == "#" {
					pos.Direction = "down"
				} else {
					pos.X = pos.X + 1

					if exists(visitedArr, pos) {
						return true
					}

					visitedArr[strconv.Itoa(pos.X)+":"+strconv.Itoa(pos.Y)+":"+pos.Direction] = true
				}
			}
		} else if pos.Direction == "left" {
			if pos.X-1 < 0 {
				isDone = true
			} else {
				if arr[pos.Y][pos.X-1] == "#" {
					pos.Direction = "up"
				} else {
					pos.X = pos.X - 1

					if exists(visitedArr, pos) {
						return true
					}

					visitedArr[strconv.Itoa(pos.X)+":"+strconv.Itoa(pos.Y)+":"+pos.Direction] = true
				}
			}
		}
	}

	return false
}

func partB() {
	arr, pos := initMap()
	count := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == "." {
				temp := arr[i][j]
				arr[i][j] = "#"

				if isLoop(arr, pos) {
					count++
				}

				arr[i][j] = temp
			}
		}
	}

	fmt.Println(count)
}

func partA() {
	arr, pos := initMap()
	isDone := false

	visitedArr := map[string]bool{
		strconv.Itoa(pos.X) + ":" + strconv.Itoa(pos.Y): true,
	}

	for !isDone {
		if pos.Direction == "up" {
			if pos.Y-1 < 0 {
				isDone = true
			} else {
				if arr[pos.Y-1][pos.X] == "#" {
					pos.Direction = "right"
				} else {
					pos.Y = pos.Y - 1
					visitedArr[strconv.Itoa(pos.X)+":"+strconv.Itoa(pos.Y)] = true
				}
			}
		} else if pos.Direction == "down" {
			if pos.Y+1 > len(arr)-1 {
				isDone = true
			} else {
				if arr[pos.Y+1][pos.X] == "#" {
					pos.Direction = "left"
				} else {
					pos.Y = pos.Y + 1
					visitedArr[strconv.Itoa(pos.X)+":"+strconv.Itoa(pos.Y)] = true
				}
			}
		} else if pos.Direction == "right" {
			if pos.X+1 > len(arr[0])-1 {
				isDone = true
			} else {
				if arr[pos.Y][pos.X+1] == "#" {
					pos.Direction = "down"
				} else {
					pos.X = pos.X + 1
					visitedArr[strconv.Itoa(pos.X)+":"+strconv.Itoa(pos.Y)] = true
				}
			}
		} else if pos.Direction == "left" {
			if pos.X-1 < 0 {
				isDone = true
			} else {
				if arr[pos.Y][pos.X-1] == "#" {
					pos.Direction = "up"
				} else {
					pos.X = pos.X - 1
					visitedArr[strconv.Itoa(pos.X)+":"+strconv.Itoa(pos.Y)] = true
				}
			}
		}
	}

	fmt.Println(len(visitedArr))
}

func main() {
	partA()
	partB()
}
