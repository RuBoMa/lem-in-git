package main

import (
	"log"
	"os"
)

func main() {
	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	// parse input
	graph, ants, err := ParseInput(string(data))
	if err != nil {
		log.Fatalln("Error parsing input:", err)
		return
	}
	// // find shortest path
	// paths := findpaths(graph)

	// // simulate ants movement
	// simulateants(ants, paths)

	// // output result
	// outputResults()
}
