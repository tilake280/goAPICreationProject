package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}
	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Println("server has failed to start, err: %s", err)
		os.Exit(1)
	}
}
