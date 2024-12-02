package utilities

import (
	"bufio"
	"os"
)

func ReadFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return []string{}, err
	}

	defer file.Close()

	var inputFile []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputFile = append(inputFile, line)
	}

	return inputFile, nil

}
