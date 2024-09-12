package server

import (
	"ishare-test/config"
	"ishare-test/db"
	"ishare-test/models"

	_ "ishare-test/docs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	App    *fiber.App
	Config *config.Config
	DB     *gorm.DB
}

func NewServer(config *config.Config) (*Server, error) {
	app := fiber.New()

	db, err := db.Init(config)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Task{})

	return &Server{
		App:    app,
		Config: config,
		DB:     db,
	}, nil
}

func (server *Server) Listen() {
	server.App.Listen(":" + server.Config.ServerPort)
}
