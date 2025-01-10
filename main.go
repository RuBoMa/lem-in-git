package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Room represents a room with its name and connections to other rooms
type Room struct {
	Name  string
	Links []*Room
}

// Ant represents an ant moving through rooms
type Ant struct {
	ID   int
	Path []*Room
	Pos  int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <input_file>")
		return
	}

	// Parse the input file
	ants, start, end, rooms, err := ParseInput(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Print the input file content
	printFileContent(os.Args[1])

	// Print the rooms and their connections
	fmt.Println("\nParsed Rooms and Connections:")
	for name, room := range rooms {
		links := []string{}
		for _, link := range room.Links {
			links = append(links, link.Name)
		}
		fmt.Printf("Room: %s, Links: %v\n", name, links)
	}

	// Find valid paths
	paths := FindPaths(start, end)
	if len(paths) == 0 {
		fmt.Println("ERROR: No valid path found")
		return
	}

	// Simulate ants moving along the paths
	SimulateAnts(ants, paths)
}

// ParseInput processes the input file to extract room details and the number of ants
func ParseInput(filename string) (int, *Room, *Room, map[string]*Room, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ants int
	rooms := make(map[string]*Room)
	var start, end *Room
	parsingRooms := true

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Handle the start and end markers
		if line == "##start" {
			parsingRooms = true
			continue
		}
		if line == "##end" {
			parsingRooms = true
			continue
		}

		// Parse rooms
		if parsingRooms && strings.Contains(line, " ") {
			fields := strings.Fields(line)
			roomName := fields[0]
			room := &Room{Name: roomName}
			rooms[roomName] = room
			if line == "##start" {
				start = room
			}
			if line == "##end" {
				end = room
			}
		} else if strings.Contains(line, "-") {
			// Parse tunnels between rooms
			parts := strings.Split(line, "-")
			room1, room2 := rooms[parts[0]], rooms[parts[1]]
			room1.Links = append(room1.Links, room2)
			room2.Links = append(room2.Links, room1)
			parsingRooms = false
		} else {
			// Parse the number of ants
			var err error
			ants, err = strconv.Atoi(line)
			if err != nil {
				return 0, nil, nil, nil, fmt.Errorf("invalid number of ants: %s", line)
			}
		}
	}

	if start == nil || end == nil {
		return 0, nil, nil, nil, fmt.Errorf("missing start or end room")
	}

	return ants, start, end, rooms, nil
}

// FindPaths finds all valid paths from the start room to the end room
func FindPaths(start, end *Room) [][]*Room {
	var paths [][]*Room
	queue := [][]*Room{{start}}

	// Perform a BFS to find paths
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		current := path[len(path)-1]
		if current == end {
			paths = append(paths, path)
			continue
		}

		for _, neighbor := range current.Links {
			if !contains(path, neighbor) {
				newPath := append([]*Room{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return paths
}

// contains checks if a room is already in the path
func contains(path []*Room, room *Room) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

// SimulateAnts moves ants along the valid paths and prints their moves
func SimulateAnts(ants int, paths [][]*Room) {
	var antList []Ant
	// Assign each ant a path
	for i := 1; i <= ants; i++ {
		antList = append(antList, Ant{ID: i, Path: paths[i%len(paths)], Pos: 0})
	}

	// Simulate the movement of ants until they reach the end
	for len(antList) > 0 {
		var moved []string
		// Move each ant one step along its path
		for i := 0; i < len(antList); i++ {
			ant := &antList[i]
			if ant.Pos < len(ant.Path)-1 {
				ant.Pos++
				moved = append(moved, fmt.Sprintf("L%d-%s", ant.ID, ant.Path[ant.Pos].Name))
			}
		}
		// Print the moves of all ants
		fmt.Println(strings.Join(moved, " "))
		// Remove ants that have reached the end
		antList = filter(antList, func(a Ant) bool { return a.Pos < len(a.Path)-1 })
	}
}

// filter removes ants that have completed their journey
func filter(ants []Ant, predicate func(Ant) bool) []Ant {
	var result []Ant
	for _, ant := range ants {
		if predicate(ant) {
			result = append(result, ant)
		}
	}
	return result
}

// printFileContent prints the content of the input file
func printFileContent(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR: Could not open input file.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Print the file's content
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
