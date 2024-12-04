package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func partA() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		charSlice := []string{}

		for _, v := range line {
			char := string(v)
			charSlice = append(charSlice, char)
		}

		matrix = append(matrix, charSlice)
	}

	word := "XMAS"
	wordReverse := "SAMX"
	wordReg := regexp.MustCompile(word)
	wordReverseReg := regexp.MustCompile(wordReverse)
	count := 0

	// horizontal
	for i := 0; i < len(matrix); i++ {
		fullString := ""

		for j := 0; j < len(matrix[i]); j++ {
			fullString += matrix[i][j]
		}

		matches := wordReg.FindAllString(fullString, -1)
		matchesReverse := wordReverseReg.FindAllString(fullString, -1)
		// fmt.Println("horizontal")
		// fmt.Println(matches, matchesReverse)
		count += len(matches) + len(matchesReverse)
	}

	// vertical
	for i := 0; i < len(matrix[0]); i++ {
		fullString := ""

		for j := 0; j < len(matrix); j++ {
			fullString += matrix[j][i]
		}

		matches := wordReg.FindAllString(fullString, -1)
		matchesReverse := wordReverseReg.FindAllString(fullString, -1)
		// fmt.Println("vertical")
		// fmt.Println(matches, matchesReverse)
		count += len(matches) + len(matchesReverse)
	}

	// diagonal
	for i := 0; i < len(matrix[0]); i++ {
		j := i
		k := 0

		fullString := ""

		for j < len(matrix[0]) && k < len(matrix) {
			fullString += matrix[k][j]

			j++
			k++
		}

		matches := wordReg.FindAllString(fullString, -1)
		matchesReverse := wordReverseReg.FindAllString(fullString, -1)
		// fmt.Println("diagonal 1.1")
		// fmt.Println(matches, matchesReverse)
		count += len(matches) + len(matchesReverse)
	}

	for i := 1; i < len(matrix); i++ {
		j := i
		k := 0

		fullString := ""

		for j < len(matrix) && k < len(matrix[0]) {
			fullString += matrix[j][k]

			j++
			k++
		}

		matches := wordReg.FindAllString(fullString, -1)
		matchesReverse := wordReverseReg.FindAllString(fullString, -1)
		// fmt.Println("diagonal 1.2")
		// fmt.Println(matches, matchesReverse)
		count += len(matches) + len(matchesReverse)
	}

	for i := len(matrix[0]) - 1; i >= 0; i-- {
		j := i
		k := 0

		fullString := ""

		for j >= 0 && k < len(matrix) {
			fullString += matrix[k][j]

			j--
			k++
		}

		matches := wordReg.FindAllString(fullString, -1)
		matchesReverse := wordReverseReg.FindAllString(fullString, -1)
		// fmt.Println("diagonal 2.1")
		// fmt.Println(matches, matchesReverse)
		count += len(matches) + len(matchesReverse)
	}

	for i := 1; i < len(matrix); i++ {
		j := i
		k := len(matrix[0]) - 1

		fullString := ""

		for j < len(matrix) && k >= 0 {
			fullString += matrix[j][k]

			j++
			k--
		}

		matches := wordReg.FindAllString(fullString, -1)
		matchesReverse := wordReverseReg.FindAllString(fullString, -1)
		// fmt.Println("diagonal 2.2")
		// fmt.Println(matches, matchesReverse)
		count += len(matches) + len(matchesReverse)
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

	matrix := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		charSlice := []string{}

		for _, v := range line {
			char := string(v)
			charSlice = append(charSlice, char)
		}

		matrix = append(matrix, charSlice)
	}

	count := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			currentChar := matrix[i][j]

			if currentChar == "A" {
				if i-1 < 0 || i+1 >= len(matrix) || j-1 < 0 || j+1 >= len(matrix[i]) {
					continue
				}

				lt := matrix[i-1][j-1]
				lb := matrix[i+1][j-1]
				rt := matrix[i-1][j+1]
				rb := matrix[i+1][j+1]

				if ((lt == "M" && rb == "S") || (lt == "S" && rb == "M")) && ((lb == "M" && rt == "S") || (lb == "S" && rt == "M")) {
					// fmt.Println(i, j)
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func main() {
	partA()
	partB()
}
