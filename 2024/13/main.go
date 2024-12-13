package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var regComp = regexp.MustCompile(`\d+`)

func calcTotalTokens(resultX string, partX1 string, partY1 string, resultY string, partX2 string, partY2 string, isPartB bool) float64 {
	resultXInt, _ := strconv.Atoi(regComp.FindString(resultX))
	resultYInt, _ := strconv.Atoi(regComp.FindString(resultY))

	if isPartB { // part B
		resultXInt += 10000000000000
		resultYInt += 10000000000000
	}

	partX1Int, _ := strconv.Atoi(regComp.FindString(partX1))
	partY1Int, _ := strconv.Atoi(regComp.FindString(partY1))

	partX2Int, _ := strconv.Atoi(regComp.FindString(partX2))
	partY2Int, _ := strconv.Atoi(regComp.FindString(partY2))

	_a := (float64(partY2Int) / float64(partY1Int)) * -1
	_partX1Float := float64(partX1Int) * _a
	_resultXInt := float64(resultXInt) * _a
	_partX := _partX1Float + float64(partX2Int)
	_result := _resultXInt + float64(resultYInt)

	_x := math.Round(_result / _partX)

	_b := float64(resultYInt) - (float64(partX2Int) * float64(_x))

	_y := math.Round(_b / float64(partY2Int))

	_t1 := (float64(partX1Int) * _x) + (float64(partY1Int) * _y)
	_t2 := (float64(partX2Int) * _x) + (float64(partY2Int) * _y)

	if _t1 == float64(resultXInt) && _t2 == float64(resultYInt) {
		return (float64(3)*_x + _y)
	} else {
		return 0
	}
}

func main() {
	file, err := os.Open("./inputs/input-1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	regX := `X\+\d+|X\-\d+`
	regY := `Y\+\d+|Y\-\d+`
	regResultX := `X=\d+`
	regResultY := `Y=\d+`

	regXComp := regexp.MustCompile(regX)
	regYComp := regexp.MustCompile(regY)
	regResultXComp := regexp.MustCompile(regResultX)
	regResultYComp := regexp.MustCompile(regResultY)

	var partX1 string
	var partY1 string
	var partX2 string
	var partY2 string
	var resultX string
	var resultY string

	lineIndex := 0
	totalTokensA := 0.0
	totalTokensB := 0.0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			if lineIndex == 0 {
				partX1 = regXComp.FindString(line)
				partX2 = regYComp.FindString(line)
			} else if lineIndex == 1 {
				partY1 = regXComp.FindString(line)
				partY2 = regYComp.FindString(line)
			} else {
				resultX = regResultXComp.FindString(line)
				resultY = regResultYComp.FindString(line)

				totalTokensA += calcTotalTokens(resultX, partX1, partY1, resultY, partX2, partY2, false)
				totalTokensB += calcTotalTokens(resultX, partX1, partY1, resultY, partX2, partY2, true)
			}

			lineIndex++
		} else {
			lineIndex = 0
			partX1 = ""
			partY1 = ""
			partX2 = ""
			partY2 = ""
			resultX = ""
			resultY = ""
		}
	}

	fmt.Println(totalTokensA, totalTokensB)
}
