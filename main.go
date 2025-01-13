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

	fileContents(os.Args[1])

	data, err := parseInput()
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}

	// Find all paths from StartRoom to EndRoom
	paths := findPaths(data.Tunnels, data.StartRoom, data.EndRoom)

	// Print the paths
	// fmt.Println("All Paths:")
	// for i, path := range paths {
	// 	fmt.Printf("Path %d: %v\n", i+1, path)
	// }

	// Result to hold all non-crossing combinations
	var result [][][]string

	// Start finding combinations
	findNonCrossingCombinations(paths, [][]string{}, 0, &result)

	if result == nil {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}

	// Printing file contents
	for _, line := range fileContent {
		fmt.Println(line)
	}
	fmt.Println()

	var combLength [][]string

	// Print all combinations
	for _, combination := range result {
		movements := simulateAntMovement(combination, data.NumAnts, data.StartRoom, data.EndRoom)
		combLength = append(combLength, movements)
	}

	// Sorthing paths from shortest to longest
	sort.Slice(combLength, func(i, j int) bool {
		return len(combLength[i]) < len(combLength[j])
	})

	for i, combinations := range combLength {
		if i == 0 {
			for _, turn := range combinations {
				fmt.Println(turn)
			}
		}

	}
}
