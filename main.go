package main

import (
	"fmt"
	"lemin/utils"
	"os"
	"path/filepath"
)

func main() {
	// Ensure a file is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <example_file>")
		os.Exit(1)
	}
	// Construct the file path
	fileName := os.Args[1]
	filePath := filepath.Join("examples/", fileName)

	// Parse the example file
	numAnts, rooms, connections, startRoom, endRoom, err := utils.ParseExample(filePath)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}

	// Print the example content (number of ants, rooms, and connections)
	fmt.Println(numAnts)
	for _, room := range rooms {
		if room == startRoom {
			fmt.Println("##start")
		}
		if room == endRoom {
			fmt.Println("##end")
		}
		fmt.Printf("%s %d %d\n", room.Name, room.X, room.Y)
	}
	for _, conn := range connections {
		fmt.Printf("%s-%s\n", conn.From, conn.To)
	}

	// Find the shortest path
	path := utils.FindShortestPath(startRoom, endRoom, connections)
	if path == nil {
		fmt.Println("No path found!")
		os.Exit(1)
	}

	// Print ant movements
	movements := utils.CalculateAntMovements(numAnts, path)
	for _, move := range movements {
		fmt.Println(move)
	}
}
