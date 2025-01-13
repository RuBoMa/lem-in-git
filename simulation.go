package main

import "fmt"

type Ant struct {
	ID       int    // Ant number
	Position string // Current room
	Done     bool
}

func simulateAntMovement(paths [][]string, numAnts int, start, end string) int {
	// Initialize ants
	ants := make([]Ant, numAnts)
	for i := 0; i < numAnts; i++ {
		ants[i] = Ant{ID: i + 1, Position: start}
	}

	// Assign ants to paths
	assignedPath := assignAntsToPaths(paths, numAnts)

	// Turn counter
	turns := 0

	// Simulation loop
	for {
		turns++
		allFinished := true
		tunnelInUse := make(map[string]bool) // Reset tunnel usage for each turn

		//fmt.Printf("Turn %d:\n", turns)
		for i := range ants {
			ant := &ants[i]

			// Skip if the ant has already reached the end
			if ant.Done {
				continue
			}

			// Get the path assigned to this ant
			path := assignedPath[ant.ID]
			currentIdx := indexOf(path, ant.Position)

			// Determine the next room
			if currentIdx+1 < len(path) {
				nextRoom := path[currentIdx+1]
				tunnel := fmt.Sprintf("%s->%s", ant.Position, nextRoom)

				// Move only if the tunnel is not in use
				if !tunnelInUse[tunnel] {
					// Mark tunnel as in use for this turn
					tunnelInUse[tunnel] = true

					// Move the ant
					ant.Position = nextRoom

					// Print movement
					//fmt.Printf("L%d - %s ", ant.ID, ant.Position)

					// Mark as finished if the ant reaches the end
					if nextRoom == end {
						ant.Done = true
					}
				}
			}

			// Check if all ants are finished
			if !ant.Done {
				allFinished = false
			}
		}

		// Break if all ants are finished
		if allFinished {
			break
		}
	}

	return turns
}

func indexOf(path []string, room string) int {
	for i, r := range path {
		if r == room {
			return i
		}
	}
	return -1 // Shouldn't happen in valid scenarios
}

func assignAntsToPaths(paths [][]string, numAnts int) map[int][]string {
	// Step 1: Initialize the assignedPath map
	assignedPath := make(map[int][]string) // Ant ID -> Path

	// Step 2: Track the number of ants assigned to each path
	pathAntCounts := make([]int, len(paths))

	// Step 3: Assign ants to paths
	for antID := 1; antID <= numAnts; antID++ {
		bestPath := -1
		minWeight := int(^uint(0) >> 1) // Max int value

		for i, path := range paths {
			weight := len(path) + pathAntCounts[i]
			if weight < minWeight {
				bestPath = i
				minWeight = weight
			}
		}

		// Assign the ant to the best path
		pathAntCounts[bestPath]++
		assignedPath[antID] = paths[bestPath]
	}

	return assignedPath
}
