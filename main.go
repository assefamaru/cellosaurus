package main

import (
	"log"
	"os"

	"github.com/assefamaru/cellosaurus-api/src/cellosaurus"
)

func main() {
	var ctx cellosaurus.Context

	ctx.Mode = "release"
	ctx.Port = os.Getenv("PORT")

	if ctx.Port == "" {
		log.Fatal("PORT must be set")
	}

	// Init server
	cellosaurus.Init(&ctx)
}
