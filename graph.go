package main

import "sort"

// Finds paths between rooms by using Distributed File System algorithm
func findPaths(graph map[string][]string, start, end string) [][]string {
	var paths [][]string
	var currentPath []string
	visited := make(map[string]bool)

	// Helper function for DFS
	var dfs func(room string)
	dfs = func(room string) {
		// Adding the current room to the path and mark it as visited
		currentPath = append(currentPath, room)
		visited[room] = true

		// If end is reached, saving the current path
		if room == end {
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			paths = append(paths, pathCopy)
		} else {
			// Exploring all unvisited neighbors
			for _, neighbor := range graph[room] {
				if !visited[neighbor] {
					dfs(neighbor)
				}
			}
		}

		// Backtracking: removing the current room and marking it as unvisited
		currentPath = currentPath[:len(currentPath)-1]
		visited[room] = false
	}

	// Staring DFS from the start room
	dfs(start)

	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	return paths
}
