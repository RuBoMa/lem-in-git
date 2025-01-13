package main

import (
	"fmt"
	"os"
	
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
		// Find the shortest path from StartRoom to EndRoom
		shortestPath, err := findShortestPath(data.Tunnels, data.StartRoom, data.EndRoom)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

	// Find all paths from StartRoom to EndRoom
	paths := findPaths(data.Tunnels, data.StartRoom, data.EndRoom)
		// Print the shortest path
	fmt.Println("Shortest Path:", shortestPath)

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
	// for _, line := range fileContent {
	// 	fmt.Println(line)
	// }

	// Print all combinations
	for i, combination := range result {
		fmt.Printf("Combination %d:\n", i+1)
		for _, path := range combination {
			fmt.Println(path)
		}
		fmt.Println()
	}
}
