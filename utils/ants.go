package utils

import (
	"fmt"
	"strings"
)

// Ant represents an ant in the simulation.
type Ant struct {
	ID       int
	Position *Room
}

// MoveTo updates the ant's position to the given room.
func (a *Ant) MoveTo(room *Room) {
	a.Position = room
}

func CalculateAntMovements(numAnts int, path []*Room) []string {
	var movements []string
	antPositions := make([]int, numAnts) // Track ant positions

	for step := 0; step < len(path)+numAnts-1; step++ {
		var stepMovements []string

		for ant := 0; ant < numAnts; ant++ {
			if antPositions[ant] < len(path) { // Ant still has room to move
				stepMovements = append(stepMovements,
					fmt.Sprintf("L%d-%s", ant+1, path[antPositions[ant]].Name))
				antPositions[ant]++ // Move the ant forward
			}
		}

		if len(stepMovements) > 0 {
			movements = append(movements, strings.Join(stepMovements, " "))
		}
	}

	return movements
}
