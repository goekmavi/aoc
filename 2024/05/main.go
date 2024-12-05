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
	regStr := `\d+`
	regComp := regexp.MustCompile(regStr)
	rules := [][]int{}
	updates := [][]int{}
	isManuals := false

	for scanner.Scan() {
		line := scanner.Text()

		if !isManuals && line == "" {
			isManuals = true
			continue
		}

		matches := regComp.FindAllString(line, -1)
		matchesInt := make([]int, len(matches))

		for i, v := range matches {
			tempV, tempVErr := strconv.Atoi(v)

			if tempVErr != nil {
				panic(tempVErr)
			}

			matchesInt[i] = tempV
		}

		if !isManuals {
			firstInt := matchesInt[0]
			secondInt := matchesInt[1]

			rules = append(rules, []int{firstInt, secondInt})
		} else {
			updates = append(updates, matchesInt)
		}
	}

	finalSlice := [][]int{}

	for _, update := range updates {
		isUpdate := true

		for _, rule := range rules {
			if !isUpdate {
				break
			}

			first := rule[0]
			second := rule[1]

			found := false

			for _, current := range update {
				if second == current {
					found = true
					continue
				}

				if first == current && found {
					isUpdate = false
					break
				}
			}
		}

		if isUpdate {
			finalSlice = append(finalSlice, update)
		}
	}

	sum := 0

	for _, list := range finalSlice {
		sum += list[(len(list) / 2)]
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
	regStr := `\d+`
	regComp := regexp.MustCompile(regStr)
	rules := [][]int{}
	updates := [][]int{}
	isManuals := false

	for scanner.Scan() {
		line := scanner.Text()

		if !isManuals && line == "" {
			isManuals = true
			continue
		}

		matches := regComp.FindAllString(line, -1)
		matchesInt := make([]int, len(matches))

		for i, v := range matches {
			tempV, tempVErr := strconv.Atoi(v)

			if tempVErr != nil {
				panic(tempVErr)
			}

			matchesInt[i] = tempV
		}

		if !isManuals {
			firstInt := matchesInt[0]
			secondInt := matchesInt[1]

			rules = append(rules, []int{firstInt, secondInt})
		} else {
			updates = append(updates, matchesInt)
		}
	}

	incorrectSlice := [][]int{}

	for _, update := range updates {
		isUpdate := true

		for _, rule := range rules {
			if !isUpdate {
				incorrectSlice = append(incorrectSlice, update)
				break
			}

			first := rule[0]
			second := rule[1]

			found := false

			for _, current := range update {
				if second == current {
					found = true
					continue
				}

				if first == current && found {
					isUpdate = false
					break
				}
			}
		}
	}

	finalSlice := [][]int{}

	for _, update := range incorrectSlice {
		finalSlice = append(finalSlice, repair(update, rules))
	}

	sum := 0

	for _, list := range finalSlice {
		sum += list[(len(list) / 2)]
	}

	fmt.Println(sum)
}

func repair(tempItems []int, rules [][]int) []int {
	isDone := false

	for !isDone {
		swapped := false

		for _, rule := range rules {
			first := rule[0]
			second := rule[1]

			found := false
			swapped = false
			var lastIndex int

			for i := 0; i < len(tempItems); i++ {
				current := tempItems[i]

				if second == current {
					found = true
					lastIndex = i

					continue
				}

				if first == current && found {
					swapped = true
					tempItems[i], tempItems[lastIndex] = tempItems[lastIndex], tempItems[i]

					break
				}
			}

			if swapped {
				break
			}
		}

		if !swapped {
			isDone = true
		}
	}

	return tempItems
}

func main() {
	partA()
	partB()
}
