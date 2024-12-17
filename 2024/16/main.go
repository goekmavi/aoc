package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var mapArr = [][]Node{}
var priorityQueue = []Node{}
var queueDone = false
var startI int
var startJ int
var endI int
var endJ int

var distances = map[string]int{}

type Node struct {
	Distance    int
	HasDistance bool
	Current     string
	Previous    string
	CurrentI    int
	CurrentJ    int
	OriginI     int
	OriginJ     int
	Facing      string
	Value       string
}

func move(node Node) []Node {
	queue := []Node{}

	prevI := node.CurrentI - 1
	prevJ := node.CurrentJ - 1
	nextI := node.CurrentI + 1
	nextJ := node.CurrentJ + 1

	update := func(i, j int, facing string, cost int) {
		if i < 0 || i >= len(mapArr) || j < 0 || j >= len(mapArr[0]) {
			return
		}

		if mapArr[i][j].Value == "#" {
			return
		}

		newDistance := node.Distance + cost

		key := strconv.Itoa(i) + ":" + strconv.Itoa(j) + ":" + facing

		if newDistance < distances[key] {
			distances[key] = newDistance
			queue = append(queue, Node{
				Distance:    newDistance,
				HasDistance: true,
				Current:     strconv.Itoa(i) + ":" + strconv.Itoa(j),
				Previous:    node.Current,
				CurrentI:    i,
				CurrentJ:    j,
				OriginI:     node.CurrentI,
				OriginJ:     node.CurrentJ,
				Facing:      facing,
				Value:       mapArr[i][j].Value,
			})
		}
	}

	if node.Facing == "toRight" {
		if prevI >= 0 {
			update(prevI, node.CurrentJ, "toTop", 1001)
		}

		if nextI < len(mapArr) {
			update(nextI, node.CurrentJ, "toBottom", 1001)
		}

		if nextJ < len(mapArr[0]) {
			update(node.CurrentI, nextJ, "toRight", 1)
		}
	} else if node.Facing == "toLeft" {
		if prevI >= 0 {
			update(prevI, node.CurrentJ, "toTop", 1001)
		}

		if nextI < len(mapArr) {
			update(nextI, node.CurrentJ, "toBottom", 1001)
		}

		if prevJ >= 0 {
			update(node.CurrentI, prevJ, "toLeft", 1)
		}
	} else if node.Facing == "toTop" {
		if prevI >= 0 {
			update(prevI, node.CurrentJ, "toTop", 1)
		}

		if prevJ >= 0 {
			update(node.CurrentI, prevJ, "toLeft", 1001)
		}

		if nextJ < len(mapArr[0]) {
			update(node.CurrentI, nextJ, "toRight", 1001)
		}
	} else if node.Facing == "toBottom" {
		if nextI < len(mapArr) {
			update(nextI, node.CurrentJ, "toBottom", 1)
		}

		if prevJ >= 0 {
			update(node.CurrentI, prevJ, "toLeft", 1001)
		}

		if nextJ < len(mapArr[0]) {
			update(node.CurrentI, nextJ, "toRight", 1001)
		}
	}

	return queue
}

func calcPoints() {
	for i := 0; i < len(mapArr); i++ {
		for j := 0; j < len(mapArr[0]); j++ {
			for _, f := range []string{"toRight", "toBottom", "toLeft", "toTop"} {
				distances[strconv.Itoa(i)+":"+strconv.Itoa(j)+":"+f] = math.MaxInt
			}
		}
	}

	startKey := strconv.Itoa(startI) + ":" + strconv.Itoa(startJ) + ":toRight"
	distances[startKey] = 0

	priorityQueue = append(priorityQueue, Node{
		Distance: 0,
		Current:  strconv.Itoa(startI) + ":" + strconv.Itoa(startJ),
		Previous: strconv.Itoa(startI) + ":" + strconv.Itoa(startJ),
		CurrentI: startI,
		CurrentJ: startJ,
		OriginI:  startI,
		OriginJ:  startJ,
		Facing:   "toRight",
		Value:    mapArr[startI][startJ].Value,
	})

	for !queueDone {
		if len(priorityQueue) == 0 {
			break
		}

		sort.Slice(priorityQueue, func(i, j int) bool {
			return priorityQueue[i].Distance < priorityQueue[j].Distance
		})

		current := priorityQueue[0]
		priorityQueue = priorityQueue[1:]

		if current.CurrentI == endI && current.CurrentJ == endJ {
			queueDone = true
			fmt.Println(current.Distance)

			return
		}

		_tmpQueue := move(current)

		if len(_tmpQueue) != 0 {
			priorityQueue = append(priorityQueue, _tmpQueue...)
		}
	}
}

func main() {
	file, err := os.Open("./inputs/input-1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := []Node{}

		for j, l := range line {
			_currentStr := string(l)
			lineArr = append(lineArr, Node{
				Value: _currentStr,
			})

			if _currentStr == "S" {
				startI = lineIndex
				startJ = j
			} else if _currentStr == "E" {
				endI = lineIndex
				endJ = j
			}
		}

		mapArr = append(mapArr, lineArr)
		lineIndex++
	}

	calcPoints()
}
