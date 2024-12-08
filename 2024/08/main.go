package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Antenna struct {
	X         int
	Y         int
	Frequency string
}

func getAntiNodesB(antenna string, mapArr []Antenna, antiNodes map[string][]Antenna, valueMap [][]string) map[string][]Antenna {
	for _, v := range mapArr {
		antiNodes[strconv.Itoa(v.X)+":"+strconv.Itoa(v.Y)] = append(antiNodes[strconv.Itoa(v.X)+":"+strconv.Itoa(v.Y)], Antenna{
			X:         v.X,
			Y:         v.Y,
			Frequency: v.Frequency,
		})
	}

	for i := 0; i < len(mapArr); i++ {
		for j := i + 1; j < len(mapArr); j++ {
			diffX := mapArr[j].X - mapArr[i].X
			diffY := mapArr[j].Y - mapArr[i].Y

			outOfMap := false
			prevX := mapArr[i].X - diffX
			prevY := mapArr[i].Y - diffY

			for !outOfMap {
				if prevX >= 0 && prevY >= 0 && prevX < len(valueMap[0]) && prevY < len(valueMap) {
					antiNodes[strconv.Itoa(prevX)+":"+strconv.Itoa(prevY)] = append(antiNodes[strconv.Itoa(prevX)+":"+strconv.Itoa(prevY)], Antenna{
						X:         prevX,
						Y:         prevY,
						Frequency: antenna,
					})

					prevX -= diffX
					prevY -= diffY
				} else {
					outOfMap = true
				}
			}

			outOfMap = false
			nextX := mapArr[j].X + diffX
			nextY := mapArr[j].Y + diffY

			for !outOfMap {
				if nextX >= 0 && nextY >= 0 && nextX < len(valueMap[0]) && nextY < len(valueMap) {
					antiNodes[strconv.Itoa(nextX)+":"+strconv.Itoa(nextY)] = append(antiNodes[strconv.Itoa(nextX)+":"+strconv.Itoa(nextY)], Antenna{
						X:         nextX,
						Y:         nextY,
						Frequency: antenna,
					})

					nextX += diffX
					nextY += diffY
				} else {
					outOfMap = true
				}
			}
		}
	}

	return antiNodes
}

func partB() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	indexValue := 0

	valueMap := [][]string{}
	antiNodes := map[string][]Antenna{}
	antennas := map[string][]Antenna{}

	for scanner.Scan() {
		line := scanner.Text()
		arr := []string{}

		for j, v := range line {
			_v := string(v)
			arr = append(arr, _v)

			if _v != "." {
				antennas[_v] = append(antennas[_v], Antenna{
					X:         j,
					Y:         indexValue,
					Frequency: _v,
				})
			}
		}

		valueMap = append(valueMap, arr)
		indexValue++
	}

	for i, tempArr := range antennas {
		antenna := i
		mapArr := tempArr

		antiNodes = getAntiNodesB(antenna, mapArr, antiNodes, valueMap)
	}

	fmt.Println(len(antiNodes))
}

func getAntiNodesA(antenna string, mapArr []Antenna, antiNodes map[string][]Antenna, valueMap [][]string) map[string][]Antenna {
	for i := 0; i < len(mapArr); i++ {
		for j := i + 1; j < len(mapArr); j++ {
			diffX := mapArr[j].X - mapArr[i].X
			diffY := mapArr[j].Y - mapArr[i].Y

			prevX := mapArr[i].X - diffX
			prevY := mapArr[i].Y - diffY

			if prevX >= 0 && prevY >= 0 && prevX < len(valueMap[0]) && prevY < len(valueMap) {
				antiNodes[strconv.Itoa(prevX)+":"+strconv.Itoa(prevY)] = append(antiNodes[strconv.Itoa(prevX)+":"+strconv.Itoa(prevY)], Antenna{
					X:         prevX,
					Y:         prevY,
					Frequency: antenna,
				})
			}

			nextX := mapArr[j].X + diffX
			nextY := mapArr[j].Y + diffY

			if nextX >= 0 && nextY >= 0 && nextX < len(valueMap[0]) && nextY < len(valueMap) {
				antiNodes[strconv.Itoa(nextX)+":"+strconv.Itoa(nextY)] = append(antiNodes[strconv.Itoa(nextX)+":"+strconv.Itoa(nextY)], Antenna{
					X:         nextX,
					Y:         nextY,
					Frequency: antenna,
				})
			}
		}
	}

	return antiNodes
}

func partA() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	indexValue := 0

	valueMap := [][]string{}
	antiNodes := map[string][]Antenna{}
	antennas := map[string][]Antenna{}

	for scanner.Scan() {
		line := scanner.Text()
		arr := []string{}

		for j, v := range line {
			_v := string(v)
			arr = append(arr, _v)

			if _v != "." {
				antennas[_v] = append(antennas[_v], Antenna{
					X:         j,
					Y:         indexValue,
					Frequency: _v,
				})
			}
		}

		valueMap = append(valueMap, arr)
		indexValue++
	}

	for i, tempArr := range antennas {
		antenna := i
		mapArr := tempArr

		antiNodes = getAntiNodesA(antenna, mapArr, antiNodes, valueMap)
	}

	fmt.Println(len(antiNodes))
}

func main() {
	partA()
	partB()
}
