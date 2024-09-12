package main

import (
	"ishare-test/config"
	"ishare-test/server"
	"ishare-test/server/routes"
	"log"
)

// @Title Secure REST API
// @Version 1.0
// @Description This is a Secure REST API written in GO.
// @SecurityDefinitions.apikey ApiKeyAuth
// @Type apiKey
// @In header
// @Name Authorization

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
