package main

import (
	"bufio"
	"fmt"
	"os"
)

func print(arr [][]string) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func update(mapSl [][]string, currentI int, currentJ int, direction string, isPartB bool) ([][]string, int, int) {
	_copySl := make([][]string, len(mapSl))
	_copyI := currentI
	_copyJ := currentJ

	for i := range mapSl {
		_copySl[i] = append([]string(nil), mapSl[i]...)
	}

	if !isPartB { // part a)
		if direction == "<" {
			indexJ := _copyJ
			//lastVal := ""

			for indexJ > 0 {
				current := mapSl[currentI][indexJ]

				if current == "#" || current == "." {
					break
				} else if current == "@" {
					_copySl[currentI][indexJ] = "."
					_copySl[currentI][indexJ-1] = "@"

					_copyI = currentI
					_copyJ = indexJ - 1
				} else if current == "O" {
					//_copySl[currentI][indexJ] = mapSl[currentI][indexJ]
					_copySl[currentI][indexJ-1] = "O"
				}

				indexJ--
			}
		} else if direction == ">" {
			indexJ := _copyJ

			for indexJ < len(mapSl[0]) {
				current := mapSl[currentI][indexJ]

				if current == "#" || current == "." {
					break
				} else if current == "@" {
					_copySl[currentI][indexJ] = "."
					_copySl[currentI][indexJ+1] = "@"

					_copyI = currentI
					_copyJ = indexJ + 1
				} else if current == "O" {
					_copySl[currentI][indexJ+1] = "O"
				}

				indexJ++
			}
		} else if direction == "^" {
			indexI := _copyI

			for indexI > 0 {
				current := mapSl[indexI][currentJ]

				if current == "#" || current == "." {
					break
				} else if current == "@" {
					_copySl[indexI][currentJ] = "."
					_copySl[indexI-1][currentJ] = "@"

					_copyI = indexI - 1
					_copyJ = currentJ
				} else if current == "O" {
					_copySl[indexI-1][currentJ] = "O"
				}

				indexI--
			}
		} else if direction == "v" {
			indexI := _copyI

			for indexI < len(mapSl) {
				current := mapSl[indexI][currentJ]

				if current == "#" || current == "." {
					break
				} else if current == "@" {
					_copySl[indexI][currentJ] = "."
					_copySl[indexI+1][currentJ] = "@"

					_copyI = indexI + 1
					_copyJ = currentJ
				} else if current == "O" {
					_copySl[indexI+1][currentJ] = "O"
				}

				indexI++
			}
		}
	} else { // part b)
		if direction == "<" {
			indexJ := _copyJ

			for indexJ > 0 {
				current := mapSl[currentI][indexJ]

				if current == "#" || current == "." {
					break
				} else if current == "@" {
					_copySl[currentI][indexJ] = "."
					_copySl[currentI][indexJ-1] = "@"

					_copyI = currentI
					_copyJ = indexJ - 1
				} else if current == "[" {
					_copySl[currentI][indexJ-1] = "["
				} else if current == "]" {
					_copySl[currentI][indexJ-1] = "]"
				}

				indexJ--
			}
		} else if direction == ">" {
			indexJ := _copyJ

			for indexJ < len(mapSl[0]) {
				current := mapSl[currentI][indexJ]

				if current == "#" || current == "." {
					break
				} else if current == "@" {
					_copySl[currentI][indexJ] = "."
					_copySl[currentI][indexJ+1] = "@"

					_copyI = currentI
					_copyJ = indexJ + 1
				} else if current == "[" {
					_copySl[currentI][indexJ+1] = "["
				} else if current == "]" {
					_copySl[currentI][indexJ+1] = "]"
				}

				indexJ++
			}
		} else if direction == "^" {
			indexI := _copyI

			for indexI > 0 {
				current := mapSl[indexI][currentJ]

				if current == "#" || current == "." {
					break
				} else if current == "@" {
					_copySl[indexI][currentJ] = "."
					_copySl[indexI-1][currentJ] = "@"

					_copyI = indexI - 1
					_copyJ = currentJ
				} else if current == "O" {
					_copySl[indexI-1][currentJ] = "O"
				}

				indexI--
			}
		} else if direction == "v" {
			indexI := _copyI

			for indexI < len(mapSl) {
				current := mapSl[indexI][currentJ]

				if current == "#" || current == "." {
					break
				} else if current == "@" {
					_copySl[indexI][currentJ] = "."
					_copySl[indexI+1][currentJ] = "@"

					_copyI = indexI + 1
					_copyJ = currentJ
				} else if current == "O" {
					_copySl[indexI+1][currentJ] = "O"
				}

				indexI++
			}
		}
	}

	return _copySl, _copyI, _copyJ
}

func canMove(mapSl [][]string, currentI int, currentJ int, direction string, isPartB bool) bool {
	freeSlot := 0

	if !isPartB { // part a)
		if direction == "<" {
			for j := currentJ; j >= 0; j-- {
				prev := mapSl[currentI][j]

				if prev == "#" {
					break
				} else if prev == "." {
					freeSlot++
				}
			}
		} else if direction == ">" {
			for j := currentJ; j < len(mapSl[0]); j++ {
				next := mapSl[currentI][j]

				if next == "#" {
					break
				} else if next == "." {
					freeSlot++
				}
			}
		} else if direction == "^" {
			for i := currentI; i >= 0; i-- {
				prev := mapSl[i][currentJ]

				if prev == "#" {
					break
				} else if prev == "." {
					freeSlot++
				}
			}
		} else if direction == "v" {
			for i := currentI; i < len(mapSl); i++ {
				next := mapSl[i][currentJ]

				if next == "#" {
					break
				} else if next == "." {
					freeSlot++
				}
			}
		}
	} else { // part b)
		if direction == "<" {
			for j := currentJ; j >= 0; j-- {
				prev := mapSl[currentI][j]

				if prev == "#" {
					break
				} else if prev == "." {
					freeSlot++
				}
			}
		} else if direction == ">" {
			for j := currentJ; j < len(mapSl[0]); j++ {
				next := mapSl[currentI][j]

				if next == "#" {
					break
				} else if next == "." {
					freeSlot++
				}
			}
		} else if direction == "^" {
			for i := currentI; i >= 0; i-- {
				prev := mapSl[i][currentJ]

				if prev == "#" {
					break
				} else if prev == "." {
					freeSlot++
				}
			}
		} else if direction == "v" {
			for i := currentI; i < len(mapSl); i++ {
				next := mapSl[i][currentJ]

				if next == "#" {
					break
				} else if next == "." {
					freeSlot++
				}
			}
		}
	}

	if freeSlot > 0 {
		return true
	} else {
		return false
	}
}

func move(mapSl [][]string, direction string, currentI int, currentJ int, isPartB bool) ([][]string, int, int) {
	_copySl := make([][]string, len(mapSl))
	_copyI := currentI
	_copyJ := currentJ

	for i := range mapSl {
		_copySl[i] = append([]string(nil), mapSl[i]...)
	}

	if direction == "<" {
		if canMove(_copySl, _copyI, _copyJ, "<", isPartB) {
			_copySl, _copyI, _copyJ = update(_copySl, _copyI, _copyJ, "<", isPartB)
		}
	} else if direction == ">" {
		if canMove(_copySl, _copyI, _copyJ, ">", isPartB) {
			_copySl, _copyI, _copyJ = update(_copySl, _copyI, _copyJ, ">", isPartB)
		}
	} else if direction == "^" {
		if canMove(_copySl, _copyI, _copyJ, "^", isPartB) {
			_copySl, _copyI, _copyJ = update(_copySl, _copyI, _copyJ, "^", isPartB)
		}
	} else if direction == "v" {
		if canMove(_copySl, _copyI, _copyJ, "v", isPartB) {
			_copySl, _copyI, _copyJ = update(_copySl, _copyI, _copyJ, "v", isPartB)
		}
	}

	return _copySl, _copyI, _copyJ
}

func main() {
	file, err := os.Open("./examples/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	mapSl := [][]string{}
	mapB := [][]string{}
	instructionMap := []string{}

	indexVal := 0
	isMoveInstruction := false
	currentI := 0
	currentJ := 0
	currentBI := 0
	currentBJ := 0

	for scanner.Scan() {
		line := scanner.Text()

		if !isMoveInstruction {
			if line != "" {
				lineSl := []string{}
				lineB := []string{}

				for j, v := range line {
					_str := string(v)

					lineSl = append(lineSl, _str)

					if _str == "@" {
						currentI = indexVal
						currentJ = j
					}

					if _str == "#" {
						lineB = append(lineB, "#")
						lineB = append(lineB, "#")
					} else if _str == "O" {
						lineB = append(lineB, "[")
						lineB = append(lineB, "]")
					} else if _str == "." {
						lineB = append(lineB, ".")
						lineB = append(lineB, ".")
					} else if _str == "@" {
						lineB = append(lineB, "@")
						lineB = append(lineB, ".")
					}
				}

				mapSl = append(mapSl, lineSl)
				mapB = append(mapB, lineB)

				indexVal++
			} else {
				isMoveInstruction = true
			}
		} else {
			for _, v := range line {
				instructionMap = append(instructionMap, string(v))
			}
		}
	}

	for i, v := range mapB {
		for j, w := range v {
			if w == "@" {
				currentBI = i
				currentBJ = j
				break
			}
		}
	}

	print(mapB)
	fmt.Println(currentBI, currentBJ)

	for _, instruction := range instructionMap {
		if instruction == "<" { // left
			mapSl, currentI, currentJ = move(mapSl, "<", currentI, currentJ, false)
			mapB, currentBI, currentBJ = move(mapB, "<", currentBI, currentBJ, true)
		} else if instruction == ">" { // right
			mapSl, currentI, currentJ = move(mapSl, ">", currentI, currentJ, false)
			mapB, currentBI, currentBJ = move(mapB, ">", currentBI, currentBJ, true)
		} else if instruction == "^" { // up
			mapSl, currentI, currentJ = move(mapSl, "^", currentI, currentJ, false)
			mapB, currentBI, currentBJ = move(mapB, "^", currentBI, currentBJ, true)
		} else if instruction == "v" { // down
			mapSl, currentI, currentJ = move(mapSl, "v", currentI, currentJ, false)
			mapB, currentBI, currentBJ = move(mapB, "v", currentBI, currentBJ, true)
		}
	}

	print(mapB)
	fmt.Println(currentBI, currentBJ)

	sum := 0

	for i, v := range mapSl {
		for j, w := range v {
			if w == "O" {
				sum += ((100 * i) + j)
			}
		}
	}

	fmt.Println("a)", sum)
}
