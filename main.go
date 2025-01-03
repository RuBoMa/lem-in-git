package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]

	data, err := os.ReadFile("examples/" + filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	fmt.Println(string(data))
}
