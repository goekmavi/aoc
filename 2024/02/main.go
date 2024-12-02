package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []string) bool {
	isIncreasing := false
	unSafe := false

	for i, v := range levels {
		if i != len(levels)-1 {
			current, errCurrent := strconv.Atoi(v)
			next, errNext := strconv.Atoi(levels[i+1])

			if errCurrent != nil || errNext != nil {
				panic("strconv error")
			}

			if i == 0 {
				if current < next {
					isIncreasing = true
				}
			}

			if current == next {
				unSafe = true
				break
			} else if current < next {
				if !isIncreasing {
					unSafe = true
					break
				}

				diff := next - current

				if diff < 1 || diff > 3 {
					unSafe = true
					break
				}
			} else {
				if isIncreasing {
					unSafe = true
					break
				}

				diff := current - next

				if diff < 1 || diff > 3 {
					unSafe = true
					break
				}
			}
		}
	}

	return unSafe
}

func partA() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Fields(report)

		if !isSafe(levels) {
			count++
		}
	}

	fmt.Println(count)
}

func partB() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Fields(report)

		for i := 0; i < len(levels); i++ {
			tempSlice := make([]string, len(levels))
			copy(tempSlice, levels) // problems without copy
			tempSlice = append(tempSlice[:i], tempSlice[i+1:]...)

			if !isSafe(tempSlice) {
				count++
				break
			}
		}
	}

	fmt.Println(count)
}

func main() {
	partA()
	partB()
}
