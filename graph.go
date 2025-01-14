package main

// Finds paths between rooms by using Depth First Search algorithm
func findPaths(links map[string][]string, start, end string) [][]string {
	var paths [][]string
	var currentPath []string
	visited := make(map[string]bool)

	// Recursive function to find all possible paths
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
			// Exploring all unvisited linkedRooms
			for _, linkedRoom := range links[room] {
				if !visited[linkedRoom] {
					dfs(linkedRoom)
				}
			}
		}

		// Backtracking: removing the current room and marking it as unvisited
		currentPath = currentPath[:len(currentPath)-1]
		visited[room] = false
	}

	// Staring DFS from the start room
	dfs(start)

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
			return true // If room already exists in the map, the paths are overlapping
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
