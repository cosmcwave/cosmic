package main

import (
	"context"
	"cosmic/backend"
	"cosmic/config"
	"flag"
	"log"
)

func main() {
	flag.Parse()

	backendConf, err := config.Load(context.Background(), "./config/backend.toml")
	if err != nil {
		panic(err)
	}

	backendsvc, err := backend.New(backendConf)
	if err != nil {
		panic(err)
	}

	if err := backendsvc.Start(); err != nil {
		log.Fatal(err)
	}
}