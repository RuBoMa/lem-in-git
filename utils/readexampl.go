package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseExample(filename string) (int, []*Room, []*Connection, *Room, *Room, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Variables to hold data
	var numAnts int
	var rooms []*Room
	var connections []*Connection
	var startRoom, endRoom *Room
	roomMap := make(map[string]*Room)

	scanner := bufio.NewScanner(file)

	// Read the file line by line
	lineNumber := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNumber++

		// Skip empty lines or comments
		if line == "" || strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		// Parse number of ants
		if numAnts == 0 {
			ants, err := strconv.Atoi(line)
			if err != nil {
				return 0, nil, nil, nil, nil, fmt.Errorf("invalid number of ants on line %d: %v", lineNumber, err)
			}
			numAnts = ants
			continue
		}

		// Parse start and end directives
		if line == "##start" || line == "##end" {
			if !scanner.Scan() {
				return 0, nil, nil, nil, nil, fmt.Errorf("missing room definition after %s on line %d", line, lineNumber)
			}
			lineNumber++
			roomLine := strings.TrimSpace(scanner.Text())
			room, err := parseRoom(roomLine)
			if err != nil {
				return 0, nil, nil, nil, nil, fmt.Errorf("invalid room definition on line %d: %v", lineNumber, err)
			}
			roomMap[room.Name] = room
			rooms = append(rooms, room)
			if line == "##start" {
				startRoom = room
			} else {
				endRoom = room
			}
			continue
		}

		// Parse rooms
		if strings.Contains(line, " ") {
			room, err := parseRoom(line)
			if err != nil {
				return 0, nil, nil, nil, nil, fmt.Errorf("invalid room definition on line %d: %v", lineNumber, err)
			}
			roomMap[room.Name] = room
			rooms = append(rooms, room)
			continue
		}

		// Parse connections
		if strings.Contains(line, "-") {
			conn, err := parseConnection(line)
			if err != nil {
				return 0, nil, nil, nil, nil, fmt.Errorf("invalid connection definition on line %d: %v", lineNumber, err)
			}
			// Ensure rooms in the connection exist
			if _, ok := roomMap[conn.From]; !ok {
				return 0, nil, nil, nil, nil, fmt.Errorf("connection references unknown room '%s' on line %d", conn.From, lineNumber)
			}
			if _, ok := roomMap[conn.To]; !ok {
				return 0, nil, nil, nil, nil, fmt.Errorf("connection references unknown room '%s' on line %d", conn.To, lineNumber)
			}
			connections = append(connections, conn)
			continue
		}
	}

	// Reorder rooms: start first, end last
	var orderedRooms []*Room
	if startRoom != nil {
		orderedRooms = append(orderedRooms, startRoom)
	}
	for _, room := range rooms {
		if room != startRoom && room != endRoom {
			orderedRooms = append(orderedRooms, room)
		}
	}
	if endRoom != nil {
		orderedRooms = append(orderedRooms, endRoom)
	}

	return numAnts, orderedRooms, connections, startRoom, endRoom, nil
}

// Helper function to parse a room definition
func parseRoom(line string) (*Room, error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid room format: %s", line)
	}
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid X coordinate: %v", err)
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invalid Y coordinate: %v", err)
	}
	return &Room{Name: parts[0], X: x, Y: y}, nil
}

// Helper function to parse a connection definition
func parseConnection(line string) (*Connection, error) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid connection format: %s", line)
	}
	return &Connection{From: parts[0], To: parts[1]}, nil
}
