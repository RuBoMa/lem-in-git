package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//validation of the arguments

	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	// parse input
	graph, ants, err := ParseInput(string(data))
	if err != nil || ants == 0 {
		log.Fatalln("Error parsing input:", err)
		return
	}

	fmt.Println(graph, ants)
	// // find shortest path
	// paths := findpaths(graph)

	// // simulate ants movement
	// simulateants(ants, paths)

	// // output result
	// outputResults()
}
