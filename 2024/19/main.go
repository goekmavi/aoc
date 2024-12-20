package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Value    string
	Possible bool
}

func main() {
	file, err := os.Open("./examples/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	ruleMap := map[string]bool{}
	items := []Item{}

	indexLine := 0

	for scanner.Scan() {
		line := scanner.Text()

		if indexLine == 0 {
			splittedLine := strings.Split(line, ", ")

			for _, v := range splittedLine {
				ruleMap[v] = true
			}
		} else if indexLine >= 3 {
			items = append(items, Item{
				Value: line,
			})
		}

		indexLine++
	}

	fmt.Println(ruleMap)
	fmt.Println(items)
}
