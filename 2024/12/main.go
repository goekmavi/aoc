package main

import (
	"bufio"
	"fmt"
	"os"
)

type Plot struct {
	Value     string
	I         int
	J         int
	HasRegion bool
}

type Region struct {
	Type  string
	Plots []Plot
}

func calcPerimeter(i int, j int, mapArr [][]Plot) int {
	perimeter := 0
	currentValue := mapArr[i][j].Value

	if i == 0 { // top
		perimeter++
	} else {
		topValue := mapArr[i-1][j].Value

		if currentValue != topValue {
			perimeter++
		}
	}

	if i == len(mapArr)-1 { // bottom
		perimeter++
	} else {
		bottomValue := mapArr[i+1][j].Value

		if currentValue != bottomValue {
			perimeter++
		}
	}

	if j == 0 { // left
		perimeter++
	} else {
		leftValue := mapArr[i][j-1].Value

		if currentValue != leftValue {
			perimeter++
		}
	}

	if j == len(mapArr[0])-1 { // right
		perimeter++
	} else {
		rightValue := mapArr[i][j+1].Value

		if currentValue != rightValue {
			perimeter++
		}
	}

	return perimeter
}

func createRegion(current Plot, mapArr [][]Plot, currentPlots []Plot, last Plot) []Plot {
	if mapArr[current.I][current.J].HasRegion {
		return currentPlots
	}

	mapArr[current.I][current.J].HasRegion = true
	currentPlots = append(currentPlots, current)

	if current.I+1 < len(mapArr) { // down
		next := mapArr[current.I+1][current.J]

		if next.Value != "start" {
			if last.Value == "start" || (last.I != next.I || last.J != next.J) {
				if next.Value == current.Value {
					_currentPlots := createRegion(mapArr[current.I+1][current.J], mapArr, currentPlots, current)

					for _, plot := range _currentPlots {
						found := false

						for _, plotTarget := range currentPlots {
							if plot.I == plotTarget.I && plot.J == plotTarget.J {
								found = true

								break
							}
						}

						if !found {
							currentPlots = append(currentPlots, plot)
						}
					}
				}
			}
		}
	}

	if current.I-1 >= 0 { // up
		next := mapArr[current.I-1][current.J]

		if next.Value != "start" {
			if last.Value == "start" || (last.I != next.I || last.J != next.J) {
				if next.Value == current.Value {
					_currentPlots := createRegion(mapArr[current.I-1][current.J], mapArr, currentPlots, current)

					for _, plot := range _currentPlots {
						found := false

						for _, plotTarget := range currentPlots {
							if plot.I == plotTarget.I && plot.J == plotTarget.J {
								found = true

								break
							}
						}

						if !found {
							currentPlots = append(currentPlots, plot)
						}
					}
				}
			}
		}
	}

	if current.J+1 < len(mapArr[0]) { // right
		next := mapArr[current.I][current.J+1]

		if next.Value != "start" {
			if last.Value == "start" || (last.I != next.I || last.J != next.J) {
				if next.Value == current.Value {
					_currentPlots := createRegion(mapArr[current.I][current.J+1], mapArr, currentPlots, current)

					for _, plot := range _currentPlots {
						found := false

						for _, plotTarget := range currentPlots {
							if plot.I == plotTarget.I && plot.J == plotTarget.J {
								found = true

								break
							}
						}

						if !found {
							currentPlots = append(currentPlots, plot)
						}
					}
				}
			}
		}
	}

	if current.J-1 >= 0 { // left
		next := mapArr[current.I][current.J-1]

		if next.Value != "start" {
			if last.Value == "start" || (last.I != next.I || last.J != next.J) {
				if next.Value == current.Value {
					_currentPlots := createRegion(mapArr[current.I][current.J-1], mapArr, currentPlots, current)

					for _, plot := range _currentPlots {
						found := false

						for _, plotTarget := range currentPlots {
							if plot.I == plotTarget.I && plot.J == plotTarget.J {
								found = true

								break
							}
						}

						if !found {
							currentPlots = append(currentPlots, plot)
						}
					}
				}
			}
		}
	}

	return currentPlots
}

// func calcSides(plots []Plot, mapArr [][]Plot) int {
// 	//fmt.Println(plots)
// 	return 0
// }

func totalPrice(mapArr [][]Plot) (int, int) {
	regions := make([]Region, 0)

	for _, vI := range mapArr {
		for j := range vI { // in Go for loops iterate with copies
			if !vI[j].HasRegion {
				_plots := createRegion(vI[j], mapArr, []Plot{}, Plot{Value: "start", I: -1, J: -1})

				for _, pV := range _plots {
					mapArr[pV.I][pV.J].HasRegion = true
				}

				_region := Region{
					Type:  vI[j].Value,
					Plots: _plots,
				}

				regions = append(regions, _region)
			}
		}
	}

	sumA := 0
	sumB := 0
	//edgeAll := [][]Plot{}

	for _, region := range regions {
		perimeter := 0
		//edgeArr := []Plot{}

		for _, plots := range region.Plots {
			_perimeter := calcPerimeter(plots.I, plots.J, mapArr)
			perimeter += _perimeter

			if _perimeter > 0 {
				//edgeArr = append(edgeArr, plots)
			}
		}

		//edgeAll = append(edgeAll, edgeArr)
		sumA += perimeter * len(region.Plots)
	}

	// for _, edges := range edgeAll {
	// 	fmt.Println(edges[0].Value, calcSides(edges, mapArr))
	// }

	return sumA, sumB
}

func main() {
	file, err := os.Open("./examples/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	mapArr := make([][]Plot, 0)
	indexI := 0

	for scanner.Scan() {
		line := scanner.Text()

		_map := make([]Plot, 0)

		for j, v := range line {
			_map = append(_map, Plot{
				Value:     string(v),
				I:         indexI,
				J:         j,
				HasRegion: false,
			})
		}

		mapArr = append(mapArr, _map)

		indexI++
	}

	totalPriceA, totalPriceB := totalPrice(mapArr)

	fmt.Println("Total price a): ", totalPriceA)
	fmt.Println("Total price b): ", totalPriceB)
}
