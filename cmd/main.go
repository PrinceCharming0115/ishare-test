package main

import (
	"ishare-test/config"
	"ishare-test/server"
	"ishare-test/server/routes"
	"log"
)

// @Title IShare TEST
// @Version 1.0

func main() {
	config := config.NewConfig()
	err := config.LoadEnvironment()
	if err != nil {
		log.Fatalln(err)
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatalln(err)
	}

	routes.ConfigureRoutes(server)
	server.Listen()
}
