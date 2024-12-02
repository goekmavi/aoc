package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partA() {
	firstCol := []int{}
	secondCol := []int{}

	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)

		firstNumber, errFirstNumber := strconv.Atoi(cols[0])
		secondNumber, errSecondNumber := strconv.Atoi(cols[1])

		if errFirstNumber != nil || errSecondNumber != nil {
			panic("strconv error")
		}

		firstCol = append(firstCol, firstNumber)
		secondCol = append(secondCol, secondNumber)
	}

	err = scanner.Err()

	if err != nil {
		panic(err)
	}

	sort.Ints(firstCol)
	sort.Ints(secondCol)

	diff := 0

	for i, v := range firstCol {
		firstNumber := v
		secondNumber := secondCol[i]

		if firstNumber > secondNumber {
			diff += firstNumber - secondNumber
		} else {
			diff += secondNumber - firstNumber
		}
	}

	fmt.Println(diff)
}

func partB() {
	firstCol := []int{}
	secondCol := []int{}

	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)

		firstNumber, errFirstNumber := strconv.Atoi(cols[0])
		secondNumber, errSecondNumber := strconv.Atoi(cols[1])

		if errFirstNumber != nil || errSecondNumber != nil {
			panic("strconv error")
		}

		firstCol = append(firstCol, firstNumber)
		secondCol = append(secondCol, secondNumber)
	}

	err = scanner.Err()

	if err != nil {
		panic(err)
	}

	score := 0

	for _, v := range firstCol {
		count := 0

		for _, w := range secondCol {
			if v == w {
				count++
			}
		}

		score += (count * v)
	}

	fmt.Println(score)
}

func main() {
	partA()
	partB()
}
