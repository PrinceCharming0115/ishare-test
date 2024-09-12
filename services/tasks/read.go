package tasks

import (
	"ishare-test/models"

	"github.com/google/uuid"
)

func (service *Service) ReadTaskByID(id uuid.UUID, task *models.Task) error {
	return service.DB.
		First(task, id).
		Error
}

func (service *Service) ReadAllTask(tasks *[]models.Task) error {
	return service.DB.
		Find(tasks).
		Error
}
