package main

import (
	"log"
	"os"

	cellosaurus "github.com/assefamaru/cellosaurus-api/src"
)

func main() {
	var c cellosaurus.Context

	c.Mode = "release"
	c.Port = os.Getenv("PORT")

	if c.Port == "" {
		log.Fatal("PORT must be set")
	}

	// Init server
	cellosaurus.Init(c)
}
