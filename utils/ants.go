package utils

// Ant represents an ant in the simulation.
type Ant struct {
	ID       int   // Unique identifier for the ant
	Position *Room // Current position of the ant
}

// MoveTo updates the ant's position to the given room.
func (a *Ant) MoveTo(room *Room) {
	a.Position = room
}
