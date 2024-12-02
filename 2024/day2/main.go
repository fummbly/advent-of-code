package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fummbly/advent-of-code/utilities"
)

func main() {

	PartOneSafeCount := 0
	PartTwoSafeCount := 0

	input, err := utilities.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	intList, err := convertToInts(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range intList {
		PartOneSafe := partOneSolve(line)
		if PartOneSafe {
			PartOneSafeCount += 1
		}
		PartTwoSafe := partTwoSolve(line)
		if PartTwoSafe {
			PartTwoSafeCount += 1
		}
		if !PartTwoSafe {
			fmt.Printf("Line: %v\n", line)
			fmt.Printf("\tSafe (Part 1): %v\n", PartOneSafe)
			fmt.Printf("\tSafe (Part 2): %v\n", PartTwoSafe)
		}

	}

	fmt.Printf("Safe Count (Part 1): %d\n", PartOneSafeCount)
	fmt.Printf("Safe Count (Part 2): %d\n", PartTwoSafeCount)

}

func partOneSolve(ints []int) bool {

	var increase bool

	for i := 1; i < len(ints); i++ {
		if ints[i-1] == ints[i] {
			return false
		}
		if i == 1 {
			increase = ints[i-1] < ints[i]
		}

		if increase && ints[i-1] > ints[i] {
			return false
		} else if !increase && ints[i-1] < ints[i] {
			return false
		}

		distance := utilities.IntAbs(ints[i-1], ints[i])
		if distance > 3 {
			return false
		}

	}

	return true

}

func partTwoSolve(ints []int) bool {

	var increase bool
	var unsafeCount int

	for i := 1; i < len(ints); i++ {
		if ints[i-1] == ints[i] {
			unsafeCount += 1
		}
		if i == 1 {
			increase = ints[i-1] < ints[i]
		}

		if increase && ints[i-1] > ints[i] {
			unsafeCount += 1
		} else if !increase && ints[i-1] < ints[i] {
			unsafeCount += 1
		}

		distance := utilities.IntAbs(ints[i-1], ints[i])
		if distance > 3 {
			unsafeCount += 1
		}

	}

	return !(unsafeCount > 1)

}

func convertToInts(input []string) ([][]int, error) {
	intList := [][]int{}

	for _, line := range input {
		splitLine := strings.Split(line, " ")
		lineList := []int{}
		for _, stringNum := range splitLine {
			intNum, err := strconv.Atoi(stringNum)
			if err != nil {
				return [][]int{}, err
			}
			lineList = append(lineList, intNum)
		}
		fmt.Println(lineList)
		intList = append(intList, lineList)
	}

	return intList, nil

}
