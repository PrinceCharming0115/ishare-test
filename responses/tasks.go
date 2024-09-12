package responses

import (
	"ishare-test/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ResponseTask struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func NewResponseTasks(c *fiber.Ctx, statusCode int, tasks []models.Task) error {
	responses := []ResponseTask{}

	for _, task := range tasks {
		responses = append(responses, ResponseTask{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return Response(c, statusCode, responses)
}

func NewResponseTask(c *fiber.Ctx, statusCode int, task models.Task) error {
	return Response(c, statusCode, ResponseTask{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	})
}
