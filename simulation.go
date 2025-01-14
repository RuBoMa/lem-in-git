package main

import (
	"fmt"
	"math"
	"strings"
)

type Ant struct {
	ID         int
	Name       int
	Position   string // Keeping track of the current room
	ReachedEnd bool
}

// Simulating ant movements by minimal amount of turns for the given path combination
func simulateAntMovement(paths [][]string, numAnts int, start, end string) []string {
	// Initialize ants
	ants := make([]Ant, numAnts)
	for i := 0; i < numAnts; i++ {
		ants[i] = Ant{ID: i + 1, Position: start}
	}

	// Assign ants to paths
	assignedPath := assignAntsToPaths(paths, numAnts)

	// Slice to store movements
	movements := []string{}
	nextAntName := 1

	// Simulation loop
	for {
		allFinished := true
		tunnelInUse := make(map[string]bool) // Reset tunnel usage for each turn

		turnMovements := []string{} // Store movements for this turn

		for i := range ants {
			ant := &ants[i]

			// Skip if the ant has already reached the end
			if ant.ReachedEnd {
				continue
			}

			// Get the path assigned to this ant
			path := assignedPath[ant.ID]
			currentRoom := indexOf(path, ant.Position)
			if currentRoom == -1 {
				return nil
			}

			// Determine the next room
			if currentRoom+1 < len(path) {
				nextRoom := path[currentRoom+1]
				tunnel := fmt.Sprintf("%s->%s", ant.Position, nextRoom)

				// Move only if the tunnel is not in use
				if !tunnelInUse[tunnel] {
					// Mark tunnel as in use for this turn
					tunnelInUse[tunnel] = true

					// Move the ant
					ant.Position = nextRoom

					// Giving the name for the ant in the order of moving
					if ant.Name == 0 {
						ant.Name = nextAntName
						nextAntName++
					}

					// Record the movement
					turnMovements = append(turnMovements, fmt.Sprintf("L%d-%s", ant.Name, ant.Position))

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

		// Add current turn movements to all movements struct
		if len(turnMovements) > 0 {
			movements = append(movements, strings.Join(turnMovements, " "))
		}

		// Break if all ants are finished
		if allFinished {
			break
		}
	}

	return movements
}

// Finding the index for the room on given path
func indexOf(path []string, room string) int {
	for i, r := range path {
		if r == room {
			return i
		}
	}
	return -1 // Shouldn't happen in valid scenarios
}

// Assigning paths to all ants based on the queue length
func assignAntsToPaths(paths [][]string, numAnts int) map[int][]string {
	// Initialize the assignedPath map
	assignedPath := make(map[int][]string) // Ant ID -> Path

	// Track the number of ants assigned to each path
	pathAntCounts := make([]int, len(paths))

	// Assign ants to paths
	for antID := 1; antID <= numAnts; antID++ {
		bestPath := -1           // Giving value outside of possible path index (assuming invalid scenario)
		minLength := math.MaxInt // Max int value

		for i, path := range paths {
			length := len(path) + pathAntCounts[i]
			if length < minLength {
				bestPath = i
				minLength = length
			}
		}

		// Assign the ant to the best path
		pathAntCounts[bestPath]++
		assignedPath[antID] = paths[bestPath]
	}

	return assignedPath
}
