package main

import (
	"log"
	"xr-central/pkg/app/server"
	"xr-central/pkg/config"
)

const confPath = "./config.json"

func main() {

	cfg, err := config.New(confPath)
	if err != nil {
		log.Fatal(err)
	}

	svr := server.New(cfg)
	svr.Serve()

}
