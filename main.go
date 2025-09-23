package main

import (
	"time"

	"github.com/Alb3G/pokedexcli/internal"
)

func main() {
	// Initialize Client
	client := internal.NewClient(5*time.Second, 5*time.Minute)
	conf := &internal.Config{
		Client: client,
	}

	startRepl(conf)
}
