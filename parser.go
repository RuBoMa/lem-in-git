package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X    int
	Y    int
}

type ParsedData struct {
	NumAnts   int
	StartRoom string
	EndRoom   string
	Rooms     map[string]Room
	Tunnels   map[string][]string
}

// Parses global variable fileContent as per structs and returns the data
func parseInput(fileContent []string) (*ParsedData, error) {
	// creating dynamic data
	parsedData := &ParsedData{
		Rooms:   make(map[string]Room),
		Tunnels: make(map[string][]string),
	}

	// Getting the number of ants
	numAnts, err := strconv.Atoi(fileContent[0])
	if err != nil || numAnts == 0 {
		return nil, fmt.Errorf("invalid number of ants '%v'", fileContent[0])
	}
	parsedData.NumAnts = numAnts

	// Parsing rooms and tunnels
	var isStart, isEnd bool
	for _, line := range fileContent[1:] {
		line = strings.TrimSpace(line)
		if line == "##start" { // flagging that the next rooms is a starting room
			isStart = true
			continue
		}
		if line == "##end" { // flagging that the next rooms is a finishing room
			isEnd = true
			continue
		}

		// Checking if the line defines a room (length of the slice is 3)
		parts := strings.Fields(line)
		if len(parts) == 3 {
			name := parts[0]
			if name[0] == 'L' || name[0] == '#' { // Name cannot start with a L or #
				return nil, fmt.Errorf("invalid room name")
			}
			x, err1 := strconv.Atoi(parts[1])
			y, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("invalid room coordinates")
			}
			room := Room{Name: name, X: x, Y: y} // creating the room
			parsedData.Rooms[name] = room        // adding the room to the map

			if isStart {
				if parsedData.StartRoom != "" {
					return nil, fmt.Errorf("several start rooms defined")
				}
				parsedData.StartRoom = name
				isStart = false
			} else if isEnd {
				if parsedData.EndRoom != "" {
					return nil, fmt.Errorf("several end rooms defined")
				}
				parsedData.EndRoom = name
				isEnd = false
			}
			continue
		}

		// Checking if the line defines a connection (includes "-")
		if strings.Contains(line, "-") {
			connParts := strings.Split(line, "-")
			if len(connParts) != 2 {
				return nil, fmt.Errorf("invalid connection: %v", line)
			}
			room1, room2 := connParts[0], connParts[1]
			parsedData.Tunnels[room1] = append(parsedData.Tunnels[room1], room2) // adding the connection on both tunnel maps
			parsedData.Tunnels[room2] = append(parsedData.Tunnels[room2], room1)
		}
	}

	// Checking if the star or end room is missing
	if parsedData.StartRoom == "" {
		return nil, fmt.Errorf("star room not defined")
	}
	if parsedData.EndRoom == "" {
		return nil, fmt.Errorf("end room not defined")
	}

	return parsedData, nil
}
