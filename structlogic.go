package main

import "fmt"

// this is just to try and understand the logic in how to call struct
type room struct {
	name string
}

func (r *room) printRoom() {
	fmt.Println(r.name)
}

func main() {
	rm1 := room{name: "room1"}
	rm1.printRoom()
}
