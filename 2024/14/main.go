package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	PosX int
	PosY int
	VelX int
	VelY int
}

var regT = `#{21}` // just tried out until 21 (https://www.reddit.com/r/adventofcode/comments/1hebtjt/2024_day_14_part_2_i_see_every_ones_solutions/)
var regTComp = regexp.MustCompile(regT)

func checkTree(space [][]int) bool {
	for _, v := range space {
		str := ""
		for _, _v := range v {
			if _v == 0 {
				str += "."
			} else {
				str += "#"
			}
		}

		res := regTComp.FindString(str)

		if res != "" {
			print(space)

			return true
		}
	}

	return false
}

func print(space [][]int) {
	for _, v := range space {
		str := ""
		for _, _v := range v {
			if _v == 0 {
				str += "."
			} else {
				str += "#"
			}
		}
		fmt.Println(str)
	}
}

func moveB(posMap map[string][]Robot, space [][]int) {
	found := false
	foundIndex := 0

	for !found {
		newPosMap := map[string][]Robot{}
		_space := make([][]int, len(space))

		for i := range space {
			_space[i] = append([]int(nil), space[i]...)
		}

		found = checkTree(_space)

		for key, robots := range posMap {
			keyData := strings.Split(key, ":")
			x, _ := strconv.Atoi(keyData[0])
			y, _ := strconv.Atoi(keyData[1])

			for _, robot := range robots {
				velX := robot.VelX
				velY := robot.VelY

				newX := x + velX

				if newX < 0 || newX >= len(_space[0]) {
					newX = newX % len(_space[0])

					if newX < 0 {
						newX += len(_space[0])
					}
				}

				newY := y + velY

				if newY < 0 || newY >= len(_space) {
					newY = newY % len(_space)

					if newY < 0 {
						newY += len(_space)
					}
				}

				_newRobot := Robot{
					PosX: newX,
					PosY: newY,
					VelX: velX,
					VelY: velY,
				}

				_key := strconv.Itoa(newX) + ":" + strconv.Itoa(newY)

				_, exists := newPosMap[_key]

				if exists {
					newPosMap[_key] = append(newPosMap[_key], _newRobot)
				} else {
					newPosMap[_key] = []Robot{
						_newRobot,
					}
				}

				_space[y][x]--
				_space[newY][newX]++
			}
		}

		space = _space
		posMap = newPosMap

		foundIndex++
	}

	fmt.Println("b)", foundIndex-1)
}

func moveA(posMap map[string][]Robot, space [][]int, seconds int) (map[string][]Robot, [][]int) {
	for i := 0; i < seconds; i++ {
		newPosMap := map[string][]Robot{}
		_space := make([][]int, len(space))

		for i := range space {
			_space[i] = append([]int(nil), space[i]...)
		}

		for key, robots := range posMap {
			keyData := strings.Split(key, ":")
			x, _ := strconv.Atoi(keyData[0])
			y, _ := strconv.Atoi(keyData[1])

			for _, robot := range robots {
				velX := robot.VelX
				velY := robot.VelY

				newX := x + velX

				if newX < 0 || newX >= len(_space[0]) {
					newX = newX % len(_space[0])

					if newX < 0 {
						newX += len(_space[0])
					}
				}

				newY := y + velY

				if newY < 0 || newY >= len(_space) {
					newY = newY % len(_space)

					if newY < 0 {
						newY += len(_space)
					}
				}

				_newRobot := Robot{
					PosX: newX,
					PosY: newY,
					VelX: velX,
					VelY: velY,
				}

				_key := strconv.Itoa(newX) + ":" + strconv.Itoa(newY)

				_, exists := newPosMap[_key]

				if exists {
					newPosMap[_key] = append(newPosMap[_key], _newRobot)
				} else {
					newPosMap[_key] = []Robot{
						_newRobot,
					}
				}

				_space[y][x]--
				_space[newY][newX]++
			}
		}

		space = _space
		posMap = newPosMap
	}

	return posMap, space
}

func main() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	reg := `-?\d+`
	regComp := regexp.MustCompile(reg)

	tilesWide := 101
	tilesTall := 103

	posMap := map[string][]Robot{}

	space := make([][]int, tilesTall)

	for i := range space {
		space[i] = make([]int, tilesWide)
	}

	for scanner.Scan() {
		line := scanner.Text()
		data := regComp.FindAllString(line, -1)

		posX, _ := strconv.Atoi(data[0])
		posY, _ := strconv.Atoi(data[1])
		velX, _ := strconv.Atoi(data[2])
		velY, _ := strconv.Atoi(data[3])

		robot := Robot{
			PosX: posX,
			PosY: posY,
			VelX: velX,
			VelY: velY,
		}

		_key := strconv.Itoa(posX) + ":" + strconv.Itoa(posY)

		_, exists := posMap[_key]

		if exists {
			posMap[_key] = append(posMap[_key], robot)
		} else {
			posMap[_key] = []Robot{
				robot,
			}
		}

		space[posY][posX] = len(posMap[_key])
	}

	_, _spaceA := moveA(posMap, space, 100)
	moveB(posMap, space)

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for i := 0; i < tilesTall; i++ {
		k := tilesTall - i - 1

		for j := 0; j < tilesWide; j++ {
			l := tilesWide - j - 1

			if i < tilesTall/2 {
				if j < tilesWide/2 {
					if _spaceA[i][j] > 0 {
						q1 += _spaceA[i][j]
					}
				}

				if l > tilesWide/2 {
					if _spaceA[i][l] > 0 {
						q2 += _spaceA[i][l]
					}
				}
			}

			if k > tilesTall/2 {
				if j < tilesWide/2 {
					if _spaceA[k][j] > 0 {
						q3 += _spaceA[k][j]
					}
				}

				if l > tilesWide/2 {
					if _spaceA[k][l] > 0 {
						q4 += _spaceA[k][l]
					}
				}
			}

			l--
		}

		k--
	}

	mul := q1 * q2 * q3 * q4

	fmt.Println("a)", mul)
}
