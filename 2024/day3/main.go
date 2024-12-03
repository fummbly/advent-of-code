package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fummbly/advent-of-code/utilities"
)

func main() {

	regMulTotal := 0
	conditionalMulTotal := 0

	regMuls := [][]string{}
	conditionalMuls := [][]string{}

	input, err := utilities.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range input {
		regMuls = append(regMuls, getMuls(line, []string{}))
		conditionalMuls = append(conditionalMuls, getMulsConditional(line, []string{}))
	}

	for _, line := range regMuls {
		for _, command := range line {
			product := calculateMuls(command)
			regMulTotal += product
		}
	}

	for _, line := range conditionalMuls {
		for _, command := range line {
			product := calculateMuls(command)
			if product == 0 {
				fmt.Printf("Invalid Command: %s\n", command)
			}

			conditionalMulTotal += product
		}
	}

	fmt.Printf("MulTotal (Part 1): %d\n", regMulTotal)
	fmt.Printf("MulTotal with Conditions (Part 2): %d\n", conditionalMulTotal)

}

func calculateMuls(mulCommand string) int {
	if string(mulCommand[3]) != "(" {
		return 0
	}

	endParenth := strings.LastIndex(mulCommand, ")")

	if utilities.IntAbs(3, endParenth) > 8 {
		return 0
	}

	numberString := mulCommand[4 : len(mulCommand)-1]

	splitNums := strings.Split(numberString, ",")

	if len(splitNums) != 2 {
		return 0
	}

	a, err := strconv.Atoi(splitNums[0])
	if err != nil {
		return 0
	}

	b, err := strconv.Atoi(splitNums[1])
	if err != nil {
		return 0
	}

	//fmt.Printf("Valid Mul: %s\n", mulCommand)
	//fmt.Println(splitNums)

	return a * b

}

func getMuls(line string, mulList []string) []string {

	if len(line) < 9 {
		return mulList
	}

	mulStart := strings.Index(line, "mul")
	if mulStart == -1 {
		return mulList
	}

	if string(line[mulStart+3]) != "(" {
		return getMuls(line[mulStart+3:], mulList)
	}

	endParen := strings.Index(line[mulStart:], ")")
	if endParen == -1 {
		return mulList
	}

	endParen += mulStart

	insideMul := strings.Index(line[mulStart+1:endParen+1], "mul")
	if insideMul != -1 {
		mulStart += insideMul + 1

	}

	mulList = append(mulList, line[mulStart:endParen+1])

	return getMuls(line[endParen+1:], mulList)

}

func getMulsConditional(line string, mulList []string) []string {

	if len(line) < 9 {
		return mulList
	}

	//dont := false

	mulStart := strings.Index(line, "mul")
	if mulStart == -1 {
		return mulList
	}

	if string(line[mulStart+3]) != "(" {
		return getMuls(line[mulStart+3:], mulList)
	}

	endParen := strings.Index(line[mulStart:], ")")
	if endParen == -1 {
		return mulList
	}

	endParen += mulStart

	insideMul := strings.Index(line[mulStart+1:endParen+1], "mul")
	if insideMul != -1 {
		mulStart += insideMul + 1

	}

	/*

		if dont {
			fmt.Printf("Dont Found: %s\n", line[dontIndex:endParen+1])
			return getMuls(line[endParen+1:], mulList)
		}
	*/

	dontIndex := strings.Index(line, "don't()")
	if dontIndex != -1 && dontIndex < endParen {
		fmt.Printf("Dont Found: %s\n", line[dontIndex:endParen+1])
	}

	mulList = append(mulList, line[mulStart:endParen+1])

	return getMuls(line[endParen+1:], mulList)

}
