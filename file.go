package main

import (
	"bufio"
	"os"
)

// Reading the file contents and stores each line to a []string
func fileContents(fileName string) ([]string, error) {

	file, err := os.Open("examples/" + fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Reading the file line by line
	var fileContent []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return fileContent, nil
}
