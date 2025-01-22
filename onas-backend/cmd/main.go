package main

import (
	"log"

	"github.com/Wacky404/onasdeals/onas-backend/internal/envs"
)

func main() {
	cfg := config{
		addr: envs.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
