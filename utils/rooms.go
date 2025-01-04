package utils

import (
	"strconv"
	"strings"
)

// Room represents a room in the graph.
type Room struct {
	Name        string
	X, Y        int
	Connections []*Room
	IsStart     bool
	IsEnd       bool
}

type Connection struct {
	From string
	To   string
}

// NewRoomFromLine parses a line to create a room.
func NewRoomFromLine(line string) *Room {
	parts := strings.Fields(line)
	x, _ := strconv.Atoi(parts[1])
	y, _ := strconv.Atoi(parts[2])
	return &Room{
		Name: parts[0],
		X:    x,
		Y:    y,
	}
}

// NewConnection creates a new connection between two rooms.
func NewConnection(from, to string) *Connection {
	return &Connection{
		From: from,
		To:   to,
	}
}
