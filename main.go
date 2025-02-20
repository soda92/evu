package main

import (
	"log"
	"os"

	"github.com/gohugoio/hugo/commands"
)

func main() {
	log.SetFlags(0)
	err := commands.Execute(os.Args[1:])
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
