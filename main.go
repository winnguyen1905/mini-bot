package main

import (
	"blitzbot/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
