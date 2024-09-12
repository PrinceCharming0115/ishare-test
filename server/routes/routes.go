package routes

import (
	s "ishare-test/server"
	"ishare-test/server/handlers"
	"ishare-test/server/middlewares"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"
)

func ConfigureRoutes(server *s.Server) {
	server.App.Get("/docs/*", swagger.WrapHandler)

	groupTasks := server.App.Group("/tasks")
	groupTasks.Use(middlewares.AuthMiddleware)
	GroupTasks(server, groupTasks)

	groupAuth := server.App.Group("/auth")
	GroupAuth(server, groupAuth)
}

func GroupTasks(server *s.Server, group fiber.Router) {
	handler := handlers.NewHandlerTasks(server)

	group.Post("", handler.CreateTask)
	group.Get("", handler.ListTasks)
	group.Get("/:id", handler.GetTask)
	group.Put("/:id", handler.UpdateTask)
	group.Delete("/:id", handler.DeleteTask)
}

func GroupAuth(server *s.Server, group fiber.Router) {
	handler := handlers.NewHandlerAuth(server)

	group.Get("/login", handler.Login)
	group.Get("/callback", handler.Callback)
}
