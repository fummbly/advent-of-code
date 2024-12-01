package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func intAbs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a

}

func main() {

	var leftList []int
	var rightList []int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// get the 2 lists numbers
		leftInt, rightInt, err := getLineNumbers(line)
		if err != nil {
			log.Fatal(err)
		}
		// append the left list int
		leftList = append(leftList, leftInt)
		// append the right list int
		rightList = append(rightList, rightInt)
		if err != nil {
			log.Fatal(err)
		}
	}

	totalDistance := partOneSolve(leftList, rightList)
	simularityScore := partTwoSolve(leftList, rightList)

	fmt.Printf("Total Distance (Part 1): %d\n", totalDistance)
	fmt.Printf("Simularity Score (Part 2): %d\n", simularityScore)
}

func partTwoSolve(leftList, rightList []int) int {
	// setup the simularity score
	simularityScore := 0
	// make a map for holding the numbers in the left list and the occurance of it
	idMap := make(map[int]int)
	for _, leftnum := range leftList {

		// initualize with 0
		idMap[leftnum] = 0

		for _, rightnum := range rightList {
			// if the number is in the right list
			if leftnum == rightnum {
				// increment the number in the map
				idMap[leftnum]++
			}

		}

	}

	// loop through the map
	for k, v := range idMap {
		// get the simularity number by multiplying the number and the occurance
		simulatrity := k * v
		// print the numbers for debugging
		fmt.Printf("Number: %d, Occurance: %d, Simularity %d\n", k, v, simulatrity)
		fmt.Printf("Simularity Score: %d\n", simularityScore)
		// add the simularity number to the simularityScore
		simularityScore += simulatrity

	}

	return simularityScore
}

func partOneSolve(leftList, rightList []int) int {
	// initualize the totalDistance value
	totalDistance := 0

	// sort each list
	slices.Sort(leftList)
	slices.Sort(rightList)
	// loop through the list
	for i := range leftList {
		fmt.Printf("Left List Number: %d, Right List Number: %d\n", leftList[i], rightList[i])
		// get the distance
		distance := intAbs(leftList[i], rightList[i])
		fmt.Printf("Distance: %d\n", distance)
		// add the distance to the totalDistance
		totalDistance += distance
	}

	return totalDistance

}

func getLineNumbers(line string) (int, int, error) {

	// Split the list by the 3 spaces seperating them
	splitLists := strings.Split(line, "   ")
	// convert the left number to an int
	leftInt, err := strconv.Atoi(splitLists[0])
	if err != nil {
		return 0, 0, err
	}
	// convert the right number to an int
	rightInt, err := strconv.Atoi(splitLists[1])
	if err != nil {
		return 0, 0, err
	}

	return leftInt, rightInt, nil

}
