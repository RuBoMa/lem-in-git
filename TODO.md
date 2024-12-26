# TODO List for Lem-in Project

## 1. Input Handling
### [ ] Read input file
- Parse the input file given as an argument.
- Check if the file is readable.

### [ ] Validate input format
- Check for:
  - Number of ants (positive integer on the first line).
  - Rooms (`name coord_x coord_y`).
  - Tunnels (`name1-name2`).
  - Special commands like `##start` and `##end`.
- Handle malformed input (e.g., missing start/end room or duplicate room names).

---

## 2. Data Structures
### [ ] Create a graph representation
- Use an adjacency list for rooms and their connections.
- Suggested struct:
```go
    type Graph struct {
        Rooms map[string]*Room
    }

    type Room struct {
        Name      string
        CoordX    int
        CoordY    int
        Links     []string
        IsStart   bool
        IsEnd     bool
    }
```

### [ ] Store the number of ants
- Keep track of the total number of ants.

---

## 3. Pathfinding
### [ ] Implement pathfinding algorithm
- Use **Breadth-First Search (BFS)** to find the shortest path(s) from `##start` to `##end`.
- If no path exists, print an error like `ERROR: no valid path found`.

### [ ] Handle multiple paths
- Find and store all shortest paths (for later simulation).

---

## 4. Simulation
### [ ] Simulate ant movement
- Use the shortest path(s) to move ants from `##start` to `##end`.
- Rules:
  - Each room can only hold one ant (except `##start` and `##end`).
  - Tunnels can only be used once per turn.
  - Avoid congestion in the paths.

### [ ] Output each turn
- Use the format `Lx-y` (ant `x` moves to room `y`).
- Print only the ants that moved in each turn.

---

## 5. Error Handling
### [ ] Detect invalid inputs
- Missing or invalid `##start` or `##end`.
- Duplicated rooms or invalid coordinates.
- Invalid links (e.g., connecting unknown rooms).
- Circular paths or self-links.

### [ ] Output meaningful error messages
- Example: `ERROR: invalid data format, no start room found`.

---

## 6. Testing
### [ ] Create test cases
- Valid examples:
  - Simple colony with one path.
  - Complex colony with multiple paths.
- Invalid examples:
  - Missing start/end.
  - Malformed input.

### [ ] Write unit tests
- Test functions like input parsing, graph construction, pathfinding, and simulation.

---

## 7. Optimization
### [ ] Optimize path selection
- Ensure ants are distributed efficiently across multiple paths to minimize moves.

---

## 8. Finalize & Document
### [ ] Clean code
- Refactor into smaller, reusable functions.
- Use descriptive variable and function names.

### [ ] Document usage
- Write a README with:
  - How to run the program.
  - Input format description.
  - Examples of output.

### [ ] Submit/test your program
