package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func partA() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := ""

	for scanner.Scan() {
		data += scanner.Text() + "\n"
	}

	regStr := `mul\(\d{1,3},\d{1,3}\)`
	regNum := `(\d{1,3})`
	regStrEx := regexp.MustCompile(regStr)
	regNumEx := regexp.MustCompile(regNum)

	matches := regStrEx.FindAllString(data, -1)
	sum := 0

	for _, v := range matches {
		numbers := regNumEx.FindAllString(v, -1)
		firstNr, errFirstNr := strconv.Atoi(numbers[0])
		secondNr, errSecondNr := strconv.Atoi(numbers[1])

		if errFirstNr != nil || errSecondNr != nil {
			panic("strconv error")
		}

		sum += (firstNr * secondNr)
	}

	fmt.Println(sum)
}

func partB() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := ""

	for scanner.Scan() {
		data += scanner.Text() + "\n"
	}

	regStr := `don't\(\)|do\(\)|mul\(\d{1,3},\d{1,3}\)`
	regNum := `(\d{1,3})`
	regStrEx := regexp.MustCompile(regStr)
	regNumEx := regexp.MustCompile(regNum)

	matches := regStrEx.FindAllString(data, -1)
	sum := 0

	do := true

	for _, v := range matches {
		if v == "do()" {
			do = true
		} else if v == "don't()" {
			do = false
		} else {
			if do {
				numbers := regNumEx.FindAllString(v, -1)
				firstNr, errFirstNr := strconv.Atoi(numbers[0])
				secondNr, errSecondNr := strconv.Atoi(numbers[1])

				if errFirstNr != nil || errSecondNr != nil {
					panic("strconv error")
				}

				sum += (firstNr * secondNr)
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	partA()
	partB()
}
