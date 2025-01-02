package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Graph struct {
	Rooms map[string]*Room
}

type Room struct {
	Name    string
	CoordX  int
	CoordY  int
	Links   []string
	IsStart bool
	IsEnd   bool
}

func ParseInput(data string) (interface{}, int, error) {
	dataSlice := strings.Split(data, "\n")
	var ants int

	var err error

	for i := 0; i < len(dataSlice)-1; i++ {
		if i == 0 {
			ants, err = strconv.Atoi(dataSlice[i])
			if err != nil {
				return nil, 0, err
			} else if ants < 1 {
				return nil, 0, fmt.Errorf("number of ants (%d) is not valid", ants)
			}
		} else {
			parts := strings.Fields(dataSlice[i])
			if len(parts) == 3 {
				name := parts[0]
				coordX, errX := strconv.Atoi(parts[1])
				coordY, errY := strconv.Atoi(parts[2])

				if errX != nil || errY != nil {
					return nil, 0, fmt.Errorf("couldn't convert coordinates from %v", parts)
				}

				start := false
				end := false

				if dataSlice[i-1] == "##start" {
					start = true
				}
				if dataSlice[i-1] == "##end" {
					end = true
				}

				roomDetails := &Room{
					Name:    name,
					CoordX:  coordX,
					CoordY:  coordY,
					IsStart: start,
					IsEnd:   end,
				}

				fmt.Println(roomDetails)
			}

		}
	}

	return "test", ants, nil

}
