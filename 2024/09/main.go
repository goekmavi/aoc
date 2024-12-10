package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func compactB(block []string) []string {
	compact := make([]string, len(block))
	copy(compact, block)

	rightIndex := len(compact) - 1

	for rightIndex >= 0 {
		if compact[rightIndex] != "." {

			startRightIndex := rightIndex
			tempRightIndex := rightIndex
			amount := 0

			for tempRightIndex >= 0 {
				if compact[tempRightIndex] == compact[rightIndex] {
					amount++
					tempRightIndex--
				} else {
					break
				}
			}

			for leftIndex := 0; leftIndex < rightIndex; leftIndex++ {
				if compact[leftIndex] == "." {
					diff := 0
					startLeftIndex := leftIndex

					for tempLeftIndex := leftIndex; tempLeftIndex < rightIndex; tempLeftIndex++ {
						if compact[tempLeftIndex] != "." {
							diff = tempLeftIndex - leftIndex

							break
						}
					}

					if diff >= amount {
						for k := 0; k < amount; k++ {
							compact[startLeftIndex+k], compact[startRightIndex-k] = compact[startRightIndex-k], compact[startLeftIndex+k]
						}

						rightIndex = tempRightIndex

						break
					} else {
						leftIndex = diff + leftIndex
					}
				}
			}

			rightIndex = tempRightIndex
		} else {
			rightIndex--
		}
	}

	return compact
}

func compact(block []string) []int {
	compactInt := []int{}
	compact := make([]string, len(block))
	copy(compact, block)

	rightIndex := len(compact) - 1

	for leftIndex := 0; leftIndex < rightIndex; leftIndex++ {
		if compact[leftIndex] == "." {
			for rightIndex >= 0 {
				if compact[rightIndex] != "." {
					compact[leftIndex], compact[rightIndex] = compact[rightIndex], "."
					rightIndex--

					break
				} else {
					rightIndex--
				}
			}
		}
	}

	for _, v := range compact {
		if v != "." {
			nr, errNr := strconv.Atoi(v)

			if errNr != nil {
				panic(errNr)
			}

			compactInt = append(compactInt, nr)
		}
	}

	return compactInt
}

func createBlock(diskMap []int) []string {
	block := []string{}

	for i, v := range diskMap {
		isFile := i%2 == 0

		if isFile {
			for j := 0; j < v; j++ {
				block = append(block, strconv.Itoa(i/2))
			}
		} else {
			for j := 0; j < v; j++ {
				block = append(block, ".")
			}
		}
	}

	return block
}

func partA(diskBlock []string) {
	diskBlockCompact := compact(diskBlock)

	checksum := 0

	for i, v := range diskBlockCompact {
		checksum += i * v
	}

	fmt.Println(checksum)
}

func partB(diskBlock []string) {
	diskBlockCompact := compactB(diskBlock)

	checksum := 0

	for i, v := range diskBlockCompact {
		if v != "." {
			nr, errNr := strconv.Atoi(v)

			if errNr != nil {
				panic(errNr)
			}

			checksum += i * nr
		}
	}

	fmt.Println(checksum)
}

func main() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	diskMap := []int{}

	for scanner.Scan() {
		line := scanner.Text() // only 1 line

		for _, v := range line {
			diskMap = append(diskMap, int(v-'0'))
		}
	}

	diskBlock := createBlock(diskMap)

	partA(diskBlock)
	partB(diskBlock)
}
