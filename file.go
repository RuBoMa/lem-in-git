package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileContent []string

// Reading the file contents into a global variable fileContent
func fileContents(fileName string) {

	file, err := os.Open("examples/" + fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Reading the file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Saving the contents into a global variable
	fileContent = lines

	// Printing file contents
	for _, line := range fileContent {
		fmt.Println(line)
	}

}
