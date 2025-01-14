package main

import (
	"bufio"
	"fmt"
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
	line := bufio.NewScanner(file)
	for line.Scan() {
		fileContent = append(fileContent, line.Text())
	}
	if err := line.Err(); err != nil {
		return nil, err
	}

	if len(fileContent) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	return fileContent, nil
}
