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

	// Find all paths from StartRoom to EndRoom
	paths := findPaths(data.Tunnels, data.StartRoom, data.EndRoom)

	// Print the paths
	fmt.Println("All Paths:")
	for i, path := range paths {
		fmt.Printf("Path %d: %v\n", i+1, path)
	}

	//fmt.Println(data.NumAnts)
	// fmt.Printf("Parsed Data:\nNumAnts: %d\nStart: %s\nEnd: %s\nRooms: %+v\nConnections: %+v\n",
	// 	data.NumAnts, data.StartRoom, data.EndRoom, data.Rooms, data.Connections)

}
