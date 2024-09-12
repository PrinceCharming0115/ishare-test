package handlers

import (
	"encoding/json"
	"ishare-test/models"
	"ishare-test/requests"
	"ishare-test/responses"
	s "ishare-test/server"
	taskservice "ishare-test/services/tasks"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HandlerTasks struct {
	Server *s.Server
}

func NewHandlerTasks(server *s.Server) *HandlerTasks {
	return &HandlerTasks{
		Server: server,
	}
}

// Refresh godoc
// @Summary Create task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param params body requests.RequestTask true "Task Request"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /tasks [post]
func (h *HandlerTasks) CreateTask(c *fiber.Ctx) error {
	request := requests.RequestTask{}

	err := json.Unmarshal(c.Body(), &request)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, "Task request data is invalid.")
	}

	now := time.Now()

	task := models.Task{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	service := taskservice.NewService(h.Server.DB)
	err = service.CreateTask(&task)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create task.")
	}

	return responses.MessageResponse(c, fiber.StatusOK, "Task is successfully created.")
}

// Refresh godoc
// @Summary List tasks
// @Tags Tasks
// @Accept json
// @Produce json
// @Success 200 {object} []responses.ResponseTask
// @Failure 500 {object} responses.Error
// @Router /tasks [get]
func (h *HandlerTasks) ListTasks(c *fiber.Ctx) error {
	tasks := []models.Task{}

	service := taskservice.NewService(h.Server.DB)
	err := service.ReadAllTask(&tasks)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to read all task.")
	}

	return responses.NewResponseTasks(c, fiber.StatusOK, tasks)
}

// Refresh godoc
// @Summary Get task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} responses.ResponseTask
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /tasks/{id} [get]
func (h *HandlerTasks) GetTask(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, "Task uuid is invalid.")
	}

	task := models.Task{}

	service := taskservice.NewService(h.Server.DB)
	err = service.ReadTaskByID(id, &task)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusNotFound, "There is no task has that id.")
	}

	return responses.NewResponseTask(c, fiber.StatusOK, task)
}

// Refresh godoc
// @Summary Update task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param params body requests.RequestTask true "Task Request"
// @Success 200 {object} []responses.ResponseTask
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /tasks/{id} [put]
func (h *HandlerTasks) UpdateTask(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, "Task uuid is invalid.")
	}

	request := requests.RequestTask{}
	err = json.Unmarshal(c.Body(), &request)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, "Task request data is invalid.")
	}

	task := models.Task{}

	service := taskservice.NewService(h.Server.DB)
	err = service.ReadTaskByID(id, &task)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusNotFound, "There is no task has that id.")
	}

	now := time.Now()

	task.UpdatedAt = &now
	task.Title = request.Title
	task.Description = request.Description
	task.Status = request.Status

	err = service.UpdateTask(&task)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update task.")
	}

	return responses.NewResponseTask(c, fiber.StatusOK, task)
}

// Refresh godoc
// @Summary Delete task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /tasks/{id} [delete]
func (h *HandlerTasks) DeleteTask(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, "Task uuid is invalid.")
	}

	task := models.Task{}

	service := taskservice.NewService(h.Server.DB)
	err = service.ReadTaskByID(id, &task)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusNotFound, "There is no task has that id.")
	}

	err = service.DeleteTask(id, &task)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete task.")
	}

	return responses.NewResponseTask(c, fiber.StatusOK, task)
}
