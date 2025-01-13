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

// Function to check if two paths overlap (share rooms)
func pathsOverlap(path1, path2 []string) bool {
	rooms := make(map[string]bool)

	// Add rooms from the first path to the map
	for _, room := range path1[1 : len(path1)-1] { // Skip start and end
		rooms[room] = true
	}

	// Check rooms in the second path
	for _, room := range path2[1 : len(path2)-1] { // Skip start and end
		if rooms[room] {
			return true
		}
	}

	return false
}

// Recursive function to generate all non-crossing path combinations
func findNonCrossingCombinations(paths [][]string, currentCombination [][]string, index int, result *[][][]string) {
	// Try adding more paths to the current combination
	for i := index; i < len(paths); i++ {
		overlaps := false
		for _, existingPath := range currentCombination {
			if pathsOverlap(existingPath, paths[i]) {
				overlaps = true
				break
			}
		}

		if !overlaps {
			// Add the path and recurse
			newCombination := append(currentCombination, paths[i])
			*result = append(*result, newCombination)
			findNonCrossingCombinations(paths, newCombination, i+1, result)
		}
	}
}

func findShortestPath(graph map[string][]string, start, end string) ([]string, error) {
	// Use the existing DFS function to find all paths
	paths := findPaths(graph, start, end)

	// Check if no paths exist
	if len(paths) == 0 {
		return nil, fmt.Errorf("no path found from %s to %s", start, end)
	}

	// The shortest path is the first in the sorted list
	return paths[0], nil
}
