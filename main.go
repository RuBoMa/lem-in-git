package main

import (
	"fmt"
	"lemin/utils"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <file_name>")
		os.Exit(1)
	}

	// Construct the file path
	fileName := os.Args[1]
	filePath := filepath.Join("examples/", fileName)

	// Parse the rooms and connections
	_, startRoom, endRoom, err := utils.ParseRooms(filePath)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	// Validate start and end rooms
	if startRoom == nil || endRoom == nil {
		fmt.Println("Start or end room is missing.")
		os.Exit(1)
	}

	// Find and print the shortest path
	path := utils.FindShortestPath(startRoom, endRoom)
	if path == nil {
		fmt.Println("No path found!")
		os.Exit(1)
	}

	fmt.Println("Shortest path found:")
	for _, room := range path {
		fmt.Println(room.Name)
	}
	// Example: Initialize ants
	ants := []*utils.Ant{
		{ID: 1, Position: startRoom},
		{ID: 2, Position: startRoom},
		{ID: 3, Position: startRoom},
	}

	// 	// Move ants along the path
	fmt.Println("Ant movements:")
	for _, ant := range ants {
		for _, step := range path[1:] { // Skip the start room
			ant.MoveTo(step)
			fmt.Printf("L%d-%s\n", ant.ID, ant.Position.Name)
		}
	}
}
