package utils

// Ant represents an ant in the simulation.
type Ant struct {
	ID       int
	Position *Room
}

// MoveTo updates the ant's position to the given room.
func (a *Ant) MoveTo(room *Room) {
	a.Position = room
}
