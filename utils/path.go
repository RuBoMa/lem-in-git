package utils

import "container/list"

// FindShortestPath uses BFS to find the shortest path from start to end.
func FindShortestPath(start, end *Room) []*Room {
	visited := make(map[*Room]bool)
	prev := make(map[*Room]*Room)
	queue := list.New()

	queue.PushBack(start)
	visited[start] = true

	// BFS loop
	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(*Room)
		if current == end {
			return reconstructPath(prev, start, end)
		}

		for _, neighbor := range current.Connections {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				prev[neighbor] = current
			}
		}
	}

	return nil // No path found
}

// reconstructPath constructs the path from start to end using the prev map.
func reconstructPath(prev map[*Room]*Room, start, end *Room) []*Room {
	var path []*Room
	for at := end; at != nil; at = prev[at] {
		path = append([]*Room{at}, path...)
	}
	return path
}
