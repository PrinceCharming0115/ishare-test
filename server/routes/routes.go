package routes

import (
	s "ishare-test/server"

	swagger "github.com/swaggo/fiber-swagger"
)

func ConfigureRoutes(server *s.Server) {
	server.App.Get("/docs/*", swagger.WrapHandler)
}
