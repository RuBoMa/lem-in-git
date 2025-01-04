package utils

func FindShortestPath(start, end *Room, connections []*Connection) []*Room {
	// Create a graph adjacency list
	graph := make(map[string][]string)
	for _, conn := range connections {
		graph[conn.From] = append(graph[conn.From], conn.To)
		graph[conn.To] = append(graph[conn.To], conn.From)
	}

	// BFS to find the shortest path
	queue := [][]string{{start.Name}}
	visited := make(map[string]bool)
	visited[start.Name] = true

	for len(queue) > 0 {
		// Dequeue the first path
		path := queue[0]
		queue = queue[1:]

		// Get the last room in the current path
		last := path[len(path)-1]

		// If we've reached the end room, reconstruct the path
		if last == end.Name {
			var result []*Room
			for _, roomName := range path {
				result = append(result, &Room{Name: roomName})
			}
			return result
		}

		// Add neighbors to the queue
		for _, neighbor := range graph[last] {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	// If no path is found, return nil
	return nil
}

// reconstructPath constructs the path from start to end using the prev map.
func reconstructPath(prev map[*Room]*Room, start, end *Room) []*Room {
	var path []*Room
	for at := end; at != nil; at = prev[at] {
		path = append([]*Room{at}, path...)
	}
	return path
}
