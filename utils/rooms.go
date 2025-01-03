package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Room represents a room in the graph.
type Room struct {
	Name        string
	Connections []*Room
	IsStart     bool
	IsEnd       bool
}

// AddConnection adds a connection to another room.
func (r *Room) AddConnection(other *Room) {
	r.Connections = append(r.Connections, other)
}

// ParseRooms parses rooms and connections from the example file.
func ParseRooms(filePath string) (map[string]*Room, *Room, *Room, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	rooms := make(map[string]*Room)
	var startRoom, endRoom *Room

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##start") {
			scanner.Scan()
			startRoom = createRoom(scanner.Text(), true, false, rooms)
		} else if strings.HasPrefix(line, "##end") {
			scanner.Scan()
			endRoom = createRoom(scanner.Text(), false, true, rooms)
		} else if strings.Contains(line, "-") {
			connectRooms(line, rooms)
		} else if len(line) > 0 && line[0] != '#' {
			createRoom(line, false, false, rooms)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return rooms, startRoom, endRoom, nil
}

// createRoom creates a room and adds it to the map.
func createRoom(line string, isStart, isEnd bool, rooms map[string]*Room) *Room {
	parts := strings.Fields(line)
	if len(parts) < 1 {
		return nil // Invalid line
	}
	name := parts[0]
	room := &Room{
		Name:    name,
		IsStart: isStart,
		IsEnd:   isEnd,
	}
	rooms[name] = room
	return room
}

// connectRooms creates a connection between two rooms.
func connectRooms(line string, rooms map[string]*Room) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return
	}
	room1, room2 := rooms[parts[0]], rooms[parts[1]]
	if room1 != nil && room2 != nil {
		room1.AddConnection(room2)
		room2.AddConnection(room1)
	}
}
