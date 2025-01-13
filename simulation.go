package main

import (
	"fmt"
	"math"
	"strings"
)

type Ant struct {
	ID         int    // Ant number
	Position   string // Current room
	ReachedEnd bool
}

func simulateAntMovement(paths [][]string, numAnts int, start, end string) []string {
	// Initialize ants
	ants := make([]Ant, numAnts)
	for i := 0; i < numAnts; i++ {
		ants[i] = Ant{ID: i + 1, Position: start}
	}

	// Assign ants to paths
	assignedPath := assignAntsToPaths(paths, numAnts)

	// Turn counter
	//turns := 0

	// Slice to store movements
	movements := []string{}

	// Simulation loop
	for {
		//turns++
		allFinished := true
		tunnelInUse := make(map[string]bool) // Reset tunnel usage for each turn

		turnMovements := []string{} // Store movements for this turn

		//fmt.Printf("Turn %d:\n", turns)
		for i := range ants {
			ant := &ants[i]

			// Skip if the ant has already reached the end
			if ant.ReachedEnd {
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

					// Record the movement
					turnMovements = append(turnMovements, fmt.Sprintf("L%d-%s", ant.ID, ant.Position))

					// Print movement
					//fmt.Printf("L%d - %s ", ant.ID, ant.Position)

					// Mark as finished if the ant reaches the end
					if nextRoom == end {
						ant.ReachedEnd = true
					}
				}
			}

			// Check if all ants are finished
			if !ant.ReachedEnd {
				allFinished = false
			}
		}

		if len(turnMovements) > 0 {
			movements = append(movements, fmt.Sprintf(strings.Join(turnMovements, " ")))
		}

		// Break if all ants are finished
		if allFinished {
			break
		}
	}

	return movements
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
	// Initialize the assignedPath map
	assignedPath := make(map[int][]string) // Ant name -> Path

	// Track the number of ants assigned to each path
	pathAntCounts := make([]int, len(paths))

	// Assign ants to paths
	for antName := 1; antName <= numAnts; antName++ {
		bestPath := -1
		minWeight := math.MaxInt // Max int value
		//minWeight := int(^uint(0) >> 1) // Max int value

		for i, path := range paths {
			weight := len(path) + pathAntCounts[i]
			if weight < minWeight {
				bestPath = i
				minWeight = weight
			}
		}

		// Assign the ant to the best path
		pathAntCounts[bestPath]++
		assignedPath[antName] = paths[bestPath]
	}

	return assignedPath
}
