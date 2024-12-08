package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equationStruct struct {
	testValue int
	numbers   []int
}

func isEquationB(finalValue int, currentSum int, currentIndex int, values []int) bool {
	currentPlusSum := currentSum
	currentMulSum := currentSum
	currentConcat := ""

	if currentSum != 0 {
		currentConcat = strconv.Itoa(currentSum)
	}

	currentPlusSum += values[currentIndex]
	currentMulSum *= values[currentIndex]
	currentConcat += strconv.Itoa(values[currentIndex])

	tempCurrentConcat, errTempCurrentConcat := strconv.Atoi(currentConcat)

	if errTempCurrentConcat != nil {
		panic(errTempCurrentConcat)
	}

	if currentSum > finalValue && currentMulSum > finalValue && tempCurrentConcat > finalValue {
		//fmt.Println(">")
		return false
	}

	if currentIndex == len(values)-1 {
		if currentPlusSum == finalValue || currentMulSum == finalValue || tempCurrentConcat == finalValue {
			// if currentPlusSum == finalValue {
			// 	fmt.Println("+")
			// } else {
			// 	fmt.Println("*")
			// }
			return true
		} else {
			//fmt.Println("<")
			return false
		}
	}

	currentIndex++

	return isEquationB(finalValue, currentPlusSum, currentIndex, values) || isEquationB(finalValue, currentMulSum, currentIndex, values) || isEquationB(finalValue, tempCurrentConcat, currentIndex, values)
}

func totalCalibrationB(equation equationStruct) int {
	if isEquationB(equation.testValue, 0, 0, equation.numbers) {
		return equation.testValue
	} else {
		return 0
	}
}

func isEquationA(finalValue int, currentSum int, currentIndex int, values []int) bool {
	currentPlusSum := currentSum
	currentMulSum := currentSum

	currentPlusSum += values[currentIndex]
	currentMulSum *= values[currentIndex]

	if currentSum > finalValue && currentMulSum > finalValue {
		//fmt.Println(">")
		return false
	}

	if currentIndex == len(values)-1 {
		if currentPlusSum == finalValue || currentMulSum == finalValue {
			// if currentPlusSum == finalValue {
			// 	fmt.Println("+")
			// } else {
			// 	fmt.Println("*")
			// }
			return true
		} else {
			//fmt.Println("<")
			return false
		}
	}

	currentIndex++

	return isEquationA(finalValue, currentPlusSum, currentIndex, values) || isEquationA(finalValue, currentMulSum, currentIndex, values)
}

func totalCalibrationA(equation equationStruct) int {
	if isEquationA(equation.testValue, 0, 0, equation.numbers) {
		return equation.testValue
	} else {
		return 0
	}
}

func partA() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := []equationStruct{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		equation := equationStruct{}

		for i, field := range fields {
			if i == 0 {
				field = strings.TrimSuffix(field, ":")
				number, err := strconv.Atoi(field)

				if err != nil {
					panic(err)
				}

				equation.testValue = number
			} else {
				number, err := strconv.Atoi(field)

				if err != nil {
					panic(err)
				}

				equation.numbers = append(equation.numbers, number)
			}
		}

		arr = append(arr, equation)
	}

	sum := 0

	for _, v := range arr {
		sum += totalCalibrationA(v)
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
	arr := []equationStruct{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		equation := equationStruct{}

		for i, field := range fields {
			if i == 0 {
				field = strings.TrimSuffix(field, ":")
				number, err := strconv.Atoi(field)

				if err != nil {
					panic(err)
				}

				equation.testValue = number
			} else {
				number, err := strconv.Atoi(field)

				if err != nil {
					panic(err)
				}

				equation.numbers = append(equation.numbers, number)
			}
		}

		arr = append(arr, equation)
	}

	sum := 0

	for _, v := range arr {
		sum += totalCalibrationB(v)
	}

	fmt.Println(sum)
}

func main() {
	partA()
	partB()
}
