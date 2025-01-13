package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: 'go run . [filename]'")
		os.Exit(1)
	}

	// Adding file contents to a global variable
	content, err := fileContents(os.Args[1])
	if err != nil {
		fmt.Println("Error reading the file contents: ", err)
		os.Exit(1)
	}

	// Parsing file contents into ParsedData struct
	data, err := parseInput(content)
	if err != nil {
		fmt.Println("ERROR: invalid data format:", err)
		os.Exit(1)
	}

	// Find all paths from StartRoom to EndRoom
	paths := findPaths(data.Tunnels, data.StartRoom, data.EndRoom)

	// Variable to hold all non-crossing combinations
	var allCombinations [][][]string

	// Recursively finding all non-crossing combinations
	findNonCrossingCombinations(paths, [][]string{}, 0, &allCombinations)

	if allCombinations == nil {
		fmt.Println("ERROR: invalid data format, no valid combinations")
		os.Exit(1)
	}

	// Printing file contents
	for _, line := range content {
		fmt.Println(line)
	}
	fmt.Println()

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

	// Print the turns on the shortest solution
	for i, solution := range allSolutions {
		if i == 0 {
			for _, turn := range solution {
				fmt.Println(turn)
			}
		}
	}
}
