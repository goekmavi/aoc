package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var secretNumbers = []int{}

func calcSecretNumber(secretNumber int, rounds int) int {
	secretFinal := secretNumber

	for range rounds {
		_rMult := secretNumber * 64
		mixedSecretNumber := _rMult ^ secretNumber
		prunedSecretNumber := mixedSecretNumber % 16777216

		_rDivide := prunedSecretNumber / 32
		mixedSecretNumber_2 := _rDivide ^ prunedSecretNumber
		prunedSecretNumber_2 := mixedSecretNumber_2 % 16777216

		_rMultFinal := prunedSecretNumber_2 * 2048
		mixedSecretNumber_3 := _rMultFinal ^ prunedSecretNumber_2
		prunedSecretNumber_3 := mixedSecretNumber_3 % 16777216

		secretFinal = prunedSecretNumber_3
		secretNumber = prunedSecretNumber_3
	}

	return secretFinal
}

func main() {
	file, err := os.Open("./examples/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		secretNumbers = append(secretNumbers, number)
	}

	sumA := 0
	sumB := 0

	for _, s := range secretNumbers {
		sumA += calcSecretNumber(s, 2000)
		sumB += calcSecretNumber(s, 10)
	}

	fmt.Println("a)", sumA)
	fmt.Println("b)", sumB)
}
