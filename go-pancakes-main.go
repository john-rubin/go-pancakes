package main

import (
	"fmt"
	"strconv"
	"strings"
)

var debugMode = true

var input = `8
-
++
--+--+
---++++-+-+
+-+-+-+-+-+-+-
++--+------+++++
++-+---+---+-+-+---
+++++++++++++++++++++`

const (
	HappySide rune = '+'
	SadSide   rune = '-'
)

type Pancake struct {
	UpSide rune
}

func flip(pancake *Pancake) {
	if pancake.UpSide == HappySide {
		pancake.UpSide = SadSide
	} else {
		pancake.UpSide = HappySide
	}
}

func main() {
	processCount, inputLines := tryParseRawInput(input)

	//if there are actually items to process, continue
	if processCount > 0 {
		fmt.Println(fmt.Sprintf("Processing %v cases...", processCount))
		for i, v := range inputLines {
			caseNumber := i + 1
			pancakeCase := getPancakesFromInput(v)
			caseIterationCount := getIterationCountForCase(pancakeCase)
			fmt.Println(fmt.Sprintf("%v: %v", caseNumber, caseIterationCount))
		}
	} else {
		fmt.Println("Nothing to do.")
	}
}

func tryParseRawInput(input string) (int, []string) {
	splits := strings.Split(input, "\n")

	//validate the first item is numeric
	count, err := strconv.Atoi(splits[0])
	if err != nil {
		fmt.Println(fmt.Sprintf("The first line (%s) must specify a single number of lines to process.", splits[0]))
	} else {
		if len(splits) < count {
			count = len(splits)
			fmt.Println(fmt.Sprintf("There are fewer lines than requested to process.  The count has been set to %v", count))
		}
		itemsToProcess := splits[1:]
		validationPassed := true
		for _, v := range itemsToProcess {
			if !validateInputLine(v) {
				fmt.Println(fmt.Sprintf("Only characters '+-' are allowed.\nThe input %s contains invalid characters.", v))
				validationPassed = false
			}
		}

		if validationPassed {
			return count, itemsToProcess
		}
	}

	return 0, nil
}

func validateInputLine(input string) bool {
	split := getRunes(input)
	for _, v := range split {
		if v != '+' && v != '-' {
			return false
		}
	}

	return true
}

func getPancakesFromInput(input string) []Pancake {
	split := getRunes(input)
	pancakes := make([]Pancake, len(split))
	z := 0
	for i := len(split) - 1; i >= 0; i-- {
		pancakes[z] = Pancake{UpSide: rune(split[i])}
		z++
	}

	return pancakes
}

func getIterationCountForCase(pancakeCase []Pancake) int {
	iterationCount := 0

	if debugMode {
		printCase(pancakeCase)
	}

	for !isCaseComplete(pancakeCase) {
		performNextFlip(pancakeCase)
		iterationCount++

		if debugMode {
			printCase(pancakeCase)
		}
	}
	return iterationCount
}

func performNextFlip(pancakeCase []Pancake) {
	for i := len(pancakeCase) - 1; i >= 0; i-- {
		//get the current pancake
		currentPancake := pancakeCase[i]

		//if there's no pancake one index lower than this pancake, immediately perform a flip
		if i-1 < 0 {
			flipPancakeGrouping(pancakeCase, i)
			break
		} else {
			previousPancake := pancakeCase[i-1]

			//if these two values are inequal, this is the grouping that needs to be flipped, starting at this index
			if currentPancake.UpSide != previousPancake.UpSide {
				flipPancakeGrouping(pancakeCase, i)
				break
			}
		}
	}
}

func flipPancakeGrouping(pancakeCase []Pancake, startIndex int) {
	flipIndex := len(pancakeCase) - 1
	for i := startIndex; i <= flipIndex; i++ {
		swapTemp := pancakeCase[flipIndex]
		pancakeCase[flipIndex] = pancakeCase[i]
		pancakeCase[i] = swapTemp
		flip(&pancakeCase[flipIndex])

		//if the indices do not match, they both need to be flipped.
		if i != flipIndex {
			flip(&pancakeCase[i])
		}
		flipIndex--
	}
}

func isCaseComplete(pancakeCase []Pancake) bool {
	for _, val := range pancakeCase {
		if val.UpSide != HappySide {
			return false
		}
	}

	return true
}

func printCase(pancakeCase []Pancake) {
	output := ""
	for i := len(pancakeCase) - 1; i >= 0; i-- {
		p := pancakeCase[i]
		output = fmt.Sprintf("%s%s", output, string(p.UpSide))
	}
	fmt.Println(output)
}

func getRunes(input string) []rune {
	runes := []rune(input)

	return runes
}
