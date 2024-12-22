package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Node struct {
	Value     string
	Distance  int
	I         int
	J         int
	PrevI     int
	PrevJ     int
	IsWall    bool
	IsVisited bool
}

var raceTrack = [][]Node{}
var wallTrack = []Node{}
var startNode = Node{}
var endNode = Node{}
var cheated = false

var heap = []Node{}

func checkNeighbours(node Node, allowCheating bool) {
	if !allowCheating {
		cheated = true
	}

	if !node.IsVisited {
		raceTrack[node.I][node.J].IsVisited = true

		topIndex := node.I - 1
		bottomIndex := node.I + 1
		leftIndex := node.J - 1
		rightIndex := node.J + 1

		if topIndex >= 0 {
			topNode := &raceTrack[topIndex][node.J]

			if topNode.Value == "#" {
				if !cheated {
					cheated = true

					if !topNode.IsVisited {
						topDistance := topNode.Distance
						newDistance := node.Distance + 1

						if newDistance < topDistance {
							topNode.Distance = newDistance
							topNode.PrevI = node.I
							topNode.PrevJ = node.J

							heap = append(heap, *topNode)
						}
					}
				}
			} else {
				if !topNode.IsVisited {
					topDistance := topNode.Distance
					newDistance := node.Distance + 1

					if newDistance < topDistance {
						topNode.Distance = newDistance
						topNode.PrevI = node.I
						topNode.PrevJ = node.J

						heap = append(heap, *topNode)
					}
				}
			}
		}

		if bottomIndex < len(raceTrack) {
			bottomNode := &raceTrack[bottomIndex][node.J]

			if bottomNode.Value == "#" {
				if !cheated {
					cheated = true

					if !bottomNode.IsVisited {
						bottomDistance := bottomNode.Distance
						newDistance := node.Distance + 1

						if newDistance < bottomDistance {
							bottomNode.Distance = newDistance
							bottomNode.PrevI = node.I
							bottomNode.PrevJ = node.J

							heap = append(heap, *bottomNode)
						}
					}
				}
			} else {
				if !bottomNode.IsVisited {
					bottomDistance := bottomNode.Distance
					newDistance := node.Distance + 1

					if newDistance < bottomDistance {
						bottomNode.Distance = newDistance
						bottomNode.PrevI = node.I
						bottomNode.PrevJ = node.J

						heap = append(heap, *bottomNode)
					}
				}
			}
		}

		if leftIndex >= 0 {
			leftNode := &raceTrack[node.I][leftIndex]

			if leftNode.Value == "#" {
				if !cheated {
					cheated = true

					if !leftNode.IsVisited {
						bottomDistance := leftNode.Distance
						newDistance := node.Distance + 1

						if newDistance < bottomDistance {
							leftNode.Distance = newDistance
							leftNode.PrevI = node.I
							leftNode.PrevJ = node.J

							heap = append(heap, *leftNode)
						}
					}
				}
			} else {
				if !leftNode.IsVisited {
					bottomDistance := leftNode.Distance
					newDistance := node.Distance + 1

					if newDistance < bottomDistance {
						leftNode.Distance = newDistance
						leftNode.PrevI = node.I
						leftNode.PrevJ = node.J

						heap = append(heap, *leftNode)
					}
				}
			}
		}

		if rightIndex < len(raceTrack[0]) {
			rightNode := &raceTrack[node.I][rightIndex]

			if rightNode.Value == "#" {
				if !cheated {
					cheated = true

					if !rightNode.IsVisited {
						bottomDistance := rightNode.Distance
						newDistance := node.Distance + 1

						if newDistance < bottomDistance {
							rightNode.Distance = newDistance
							rightNode.PrevI = node.I
							rightNode.PrevJ = node.J

							heap = append(heap, *rightNode)
						}
					}
				}
			} else {
				if !rightNode.IsVisited {
					bottomDistance := rightNode.Distance
					newDistance := node.Distance + 1

					if newDistance < bottomDistance {
						rightNode.Distance = newDistance
						rightNode.PrevI = node.I
						rightNode.PrevJ = node.J

						heap = append(heap, *rightNode)
					}
				}
			}
		}
	}
}

func shortestPath(allowCheating bool) int {
	for len(heap) != 0 {
		sort.Slice(heap, func(i, j int) bool {
			return heap[i].Distance < heap[j].Distance
		})

		currentNode := heap[0]

		if currentNode.I == endNode.I && currentNode.J == endNode.J {
			return currentNode.Distance
		}

		heap = heap[1:]

		checkNeighbours(currentNode, allowCheating)
	}

	return 0
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

		for j, v := range line {
			if string(v) != "#" {
				lineArr = append(lineArr, Node{
					Value:    string(v),
					I:        lineIndex,
					J:        j,
					Distance: math.MaxInt,
				})
			} else {
				lineArr = append(lineArr, Node{
					Value:    string(v),
					I:        lineIndex,
					J:        j,
					Distance: math.MaxInt,
					IsWall:   true,
				})

				wallTrack = append(wallTrack, Node{
					Value:  string(v),
					I:      lineIndex,
					J:      j,
					IsWall: true,
				})
			}

			if string(v) == "S" {
				startNode = Node{
					Value:    "S",
					I:        lineIndex,
					J:        j,
					Distance: 0,
				}
			} else if string(v) == "E" {
				endNode = Node{
					Value:    "E",
					I:        lineIndex,
					J:        j,
					Distance: math.MaxInt,
				}
			}
		}

		raceTrack = append(raceTrack, lineArr)
		lineIndex++
	}

	heap = append(heap, startNode)
	ps := shortestPath(false)
	psSum := 0

	for _, w := range wallTrack {
		mustChangedI := w.I
		mustChangedJ := w.J

		oldVal := raceTrack[mustChangedI][mustChangedJ].Value
		oldWall := raceTrack[mustChangedI][mustChangedJ].IsWall

		raceTrack[mustChangedI][mustChangedJ].Value = "."
		raceTrack[mustChangedI][mustChangedJ].IsWall = false

		heap = append(heap, startNode)
		_ps := shortestPath(true)

		if ps-_ps >= 100 {
			fmt.Println(ps, _ps, ps-_ps)
			psSum++
		}

		raceTrack[mustChangedI][mustChangedJ].Value = oldVal
		raceTrack[mustChangedI][mustChangedJ].IsWall = oldWall

		for i, v := range raceTrack {
			for j := range v {
				raceTrack[i][j].Distance = math.MaxInt
				raceTrack[i][j].PrevI = 0
				raceTrack[i][j].PrevJ = 0
				raceTrack[i][j].IsWall = false
				raceTrack[i][j].IsVisited = false
			}
		}

		cheated = false
		heap = []Node{}
	}

	fmt.Println(psSum)
}
