package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	if len(os.Args) != 2 {
		Exit(fmt.Sprintf("Usage: 'go run . [filename]'"))
	}

	// Adding file contents to a global variable
	content, err := fileContents(os.Args[1])
	if err != nil {
		Exit(fmt.Sprint("Error reading the file contents: ", err))
	}

	// Parsing file contents into ParsedData struct
	data, err := parseInput(content)
	if err != nil {
		Exit(fmt.Sprint("ERROR: invalid data format: ", err))
	}

	// Find all paths from StartRoom to EndRoom
	paths := findPaths(data.Tunnels, data.StartRoom, data.EndRoom)

	// Variable to hold all non-crossing combinations
	var allCombinations [][][]string

	// Recursively finding all non-crossing combinations
	findNonCrossingCombinations(paths, [][]string{}, 0, &allCombinations)

	if allCombinations == nil {
		Exit(fmt.Sprint("ERROR: invalid data format: no valid combinations"))
	}

	var allSolutions [][]string

	// Print all combinations
	for _, combination := range allCombinations {
		movements := simulateAntMovement(combination, data.NumAnts, data.StartRoom, data.EndRoom)
		allSolutions = append(allSolutions, movements)
	}

	// Sorthing solutions from shortest to longest
	sort.Slice(allSolutions, func(i, j int) bool {
		return len(allSolutions[i]) < len(allSolutions[j])
	})

	PrintResult(content, allSolutions[0])
}

// Print exit message and exit program
func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Print file contents and the turns on the shortest path combination
func PrintResult(content, solution []string) {
	// Printing file contents
	for _, line := range content {
		fmt.Println(line)
	}
	fmt.Println()

	// Print the turns
	for i, turn := range solution {

		fmt.Printf("Turn %d: %v\n", i+1, turn)
	}
}
